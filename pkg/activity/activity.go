package activity

import (
	"errors"
	"fmt"
	"github.com/ridegopher/strava/pkg/athlete"
	"github.com/ridegopher/strava/pkg/format"
	"github.com/ridegopher/strava/pkg/location"
	"github.com/strava/go.strava"
	"strings"
)

type Service struct {
	updater    *strava.ActivitiesPutCall
	activity   *strava.ActivityDetailed
	stravaSvc  *strava.ActivitiesService
	athlete    *athlete.Athlete
	athleteSvc *athlete.Service
}

func New(athleteId int, activityId int64) (*Service, error) {

	athleteSvc, err := athlete.New()
	if err != nil {
		return nil, err
	}

	athlete, err := athleteSvc.Get(athleteId)
	if err != nil {
		return nil, err
	}

	client := strava.NewClient(athlete.AccessToken)
	stravaSvc := strava.NewActivitiesService(client)

	activity, err := stravaSvc.Get(activityId).Do()
	if err != nil {
		return nil, err
	}

	updater := stravaSvc.Update(activityId)

	newService := &Service{
		updater:    updater,
		activity:   activity,
		stravaSvc:  stravaSvc,
		athlete:    athlete,
		athleteSvc: athleteSvc,
	}

	return newService, nil
}

func (s *Service) ProcessActivityCreate() (string, error) {

	activityKey := string(s.activity.Type)

	if s.activity.Trainer == true {
		activityKey += "-trainer"
		if activity, ok := s.athlete.Activities[activityKey]; ok {
			return s.applyRules(false, activity)
		}
	}

	commuteKeys, err := s.commuteKeys()
	if err == nil {
		for _, key := range commuteKeys {

			if commute, ok := s.athlete.Commutes[key]; ok {

				if s.activity.Distance <= commute.Distance {

					activityKey = string(s.activity.Type)
					if activity, ok := commute.Activities[activityKey]; ok {
						return s.applyRules(true, activity)
					}

				}

			}
		}

	}

	if activity, ok := s.athlete.Activities[activityKey]; ok {
		return s.applyRules(false, activity)
	}

	return "Nothing processed", nil
}

func (s *Service) applyRules(isCommute bool, activity athlete.Activity) (string, error) {

	if isCommute != s.activity.Commute {
		s.updater.Commute(isCommute)
	}

	if activity.Private != s.activity.Private {
		s.updater.Private(activity.Private)
	}

	newName := s.formatName(activity.NameFormats)
	if newName != "" {
		s.updater.Name(newName)
	}

	description := activity.Description
	if description != "" || s.athlete.UpdateDescription {
		if description == "" {
			description = "updated by ridegopher.com"
		} else {
			description += " -- updated by ridegopher.com"
		}
		s.updater.Description(description)
	}

	if activity.GearId != "" {
		s.updater.Gear(activity.GearId)
	}

	_, err := s.updater.Do()
	return "", err

}

func (s *Service) formatLocation() string {
	locationKey, err := latLngKey(s.activity.StartLocation)
	if err != nil {
		return ""
	}

	if location, ok := s.athlete.Locations[locationKey]; ok {
		return location
	}

	lat, lng, err := latLng(s.activity.StartLocation)
	if err != nil {
		return ""
	}

	locSvc, err := location.New()
	if err != nil {
		return ""
	}

	location, err := locSvc.GetLocation(lat, lng)
	if err != nil {
		return ""
	}

	s.athleteSvc.UpdateLocations(s.athlete, locationKey, location)

	return location

}

func (s *Service) formatName(formats []format.NameFormat) string {
	if len(formats) == 0 {
		return ""
	}

	var rv []string

	for _, f := range formats {

		switch f {
		case format.NameFormats.ExistingName:
			rv = append(rv, s.activity.Name)

		case format.NameFormats.Location:
			location := s.formatLocation()
			if location != "" {
				rv = append(rv, location)
			}

		case format.NameFormats.ElapsedTime:
			elapsedTime := format.Time(s.activity.ElapsedTime)
			if elapsedTime != "" {
				rv = append(rv, elapsedTime)
			}

		case format.NameFormats.MovingTime:
			movingTime := format.Time(s.activity.MovingTime)
			if movingTime != "" {
				rv = append(rv, movingTime)
			}

		case format.NameFormats.StartDate:
			startDate := format.StartDate(s.activity.StartDate, s.athlete.DateFormat)
			if startDate != "" {
				rv = append(rv, startDate)
			}

		case format.NameFormats.Dash:
			rv = append(rv, "-")

		case format.NameFormats.Slash:
			rv = append(rv, "/")

		}
	}

	return strings.Join(rv, " ")

}

func (s *Service) commuteKeys() ([2]string, error) {
	start, err := latLngKey(s.activity.ActivitySummary.StartLocation)
	if err != nil {
		return [2]string{}, err
	}
	end, err := latLngKey(s.activity.ActivitySummary.EndLocation)
	if err != nil {
		return [2]string{}, err
	}

	key1 := fmt.Sprintf("%s,%s", start, end)
	key2 := fmt.Sprintf("%s,%s", end, start)

	return [2]string{key1, key2}, nil

}

func latLngKey(location strava.Location) (string, error) {
	lat, lng, err := latLng(location)
	if err != nil {
		return "", errors.New("location data not correct")
	}
	return fmt.Sprintf("%.2f,%.2f", lat, lng), nil
}

func latLng(location strava.Location) (float64, float64, error) {
	if len(location) == 2 {
		lat := location[0]
		lng := location[1]
		return lat, lng, nil
	}
	return 0, 0, errors.New("no location data found")
}
