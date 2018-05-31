package format

import (
	"time"
)

type DateFormat string

var DateFormats = struct {
	YMDDash  DateFormat
	YMDSlash DateFormat
	YMDDot   DateFormat
	DMYDash  DateFormat
	DMYSlash DateFormat
	DMYDot   DateFormat
}{
	"2006-01-02",
	"2006/01/02",
	"2006.01.02",
	"02-01-2006",
	"02/01/2006",
	"02.01.2006",
}

type Service struct {
	DateFormat
}

func StartDate(date time.Time, dateFormat DateFormat) string {
	return date.Format(string(dateFormat))
}
