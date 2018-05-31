package format_test

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/format"
	"testing"
	"time"
)

func TestFormat_StartDate(t *testing.T) {

	test, err := time.Parse(time.RFC3339, "2018-05-23T07:57:12Z")
	if err != nil {
		t.Error(err)
	}

	inputs := map[string]format.DateFormat{
		"23-05-2018": format.DateFormats.DMYDash,
		"23/05/2018": format.DateFormats.DMYSlash,
		"23.05.2018": format.DateFormats.DMYDot,
	}

	for expected, dateFormat := range inputs {
		rv := format.StartDate(test, dateFormat)
		if expected != rv {
			t.Error(fmt.Sprintf("Result %s doesn't match test %s", rv, test))
		}
	}
}
