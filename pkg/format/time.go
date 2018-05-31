package format

import (
	"math"
	"strconv"
	"strings"
)

func Time(seconds int) string {
	hrs := math.Floor(float64(seconds) / 60 / 60)
	secs := seconds % (60 * 60)
	mins := math.Floor(float64(secs) / 60)
	secs = seconds % 60

	var rv []string
	if hrs > 0 {
		rv = []string{
			strconv.Itoa(int(hrs)),
			doubleZeroString(int(mins)),
			doubleZeroString(int(secs)),
		}
	} else {
		rv = []string{
			strconv.Itoa(int(mins)),
			doubleZeroString(int(secs)),
		}
	}
	return strings.Join(rv, ":")
}

func doubleZeroString(input int) string {
	if input < 10 {
		return "0" + strconv.Itoa(input)
	}
	return strconv.Itoa(input)
}
