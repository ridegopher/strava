package location

import (
	"context"
	"errors"
	"googlemaps.github.io/maps"
	"os"
	"strings"
)

type Service struct {
	*maps.Client
}

func New() (*Service, error) {
	if os.Getenv("GOOGLE_API_KEY") == "" {
		return nil, errors.New("skipping getLocation; $GOOGLE_API_KEY not set")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		return nil, errors.New("problem with Google Maps Client")
	}
	return &Service{Client: c}, nil

}

// Get
func (s *Service) GetLocation(lat, lng float64) (string, error) {

	req := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},

		ResultType:   []string{"neighborhood|locality|administrative_area_level_1"},
		LocationType: []maps.GeocodeAccuracy{"APPROXIMATE"},
	}
	result, err := s.Client.Geocode(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(result) > 0 {
		address := result[0]

		var neighborhood, locality, area string
		for _, comp := range address.AddressComponents {
			for _, t := range comp.Types {
				switch t {
				case "neighborhood":
					neighborhood = comp.ShortName
				case "locality":
					locality = comp.ShortName
				case "administrative_area_level_1":
					area = comp.ShortName
				}
			}
		}
		var rv []string

		if neighborhood != "" {
			rv = append(rv, neighborhood)
		}
		if locality != "" {
			rv = append(rv, locality)
		}
		if area != "" {
			rv = append(rv, area)
		}
		return strings.Join(rv, " "), nil
	}

	return "", errors.New("no address found")
}
