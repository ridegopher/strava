package format

type NameFormat string

var NameFormats = struct {
	ExistingName NameFormat
	Location     NameFormat
	ElapsedTime  NameFormat
	MovingTime   NameFormat
	StartDate    NameFormat
	Distance     NameFormat
	DistanceKm   NameFormat
	Dash         NameFormat
	Slash        NameFormat
}{
	"existing",
	"location",
	"elapsed_time",
	"moving_time",
	"start_date",
	"distance",
	"distance_km",
	"dash",
	"slash",
}

var DefaultNameFormat = []NameFormat{
	NameFormats.ExistingName,
	NameFormats.Dash,
	NameFormats.MovingTime,
	NameFormats.StartDate,
}

//func Name(activity *strava.ActivityDetailed, formats []NameFormat) (string, error) {
//
//	if len(formats) == 0 {
//		return "", errors.New("missing name formats")
//	}
//
//	var rv []string
//
//	for _, f := range formats {
//
//		switch f {
//		case NameFormats.ExistingName:
//			rv = append(rv, activity.Name)
//
//		case NameFormats.Location:
//			locSvc, err := NewLocationService()
//			if err != nil {
//				return "", err
//			}
//
//			if len(activity.ActivitySummary.StartLocation) < 2 {
//				return "", errors.New("no start location coords to use")
//			}
//
//			lat := activity.ActivitySummary.StartLocation[0]
//			lng := activity.ActivitySummary.StartLocation[1]
//
//			location, err := locSvc.GetLocation(lat, lng)
//			if err != nil {
//				return "", err
//			}
//			if location != "" {
//				rv = append(rv, location)
//			}
//
//		case NameFormats.ElapsedTime:
//			elapsedTime := Time(activity.ElapsedTime)
//			if elapsedTime != "" {
//				rv = append(rv, elapsedTime)
//			}
//
//		case NameFormats.MovingTime:
//			movingTime := Time(activity.MovingTime)
//			if movingTime != "" {
//				rv = append(rv, movingTime)
//			}
//
//		case NameFormats.StartDate:
//			//startDate := StartDate(activity)
//			//if startDate != "" {
//			//	rv = append(rv, startDate)
//			//}
//
//		case NameFormats.Dash:
//			rv = append(rv, "-")
//
//		case NameFormats.Slash:
//			rv = append(rv, "/")
//
//		}
//	}
//
//	return strings.Join(rv, " ")
//}
