package format_test

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/format"
	"testing"
)

func TestFormat_Distance(t *testing.T) {

	type tester struct {
		d float64
		f format.Conversion
	}

	inputs := map[string]tester{
		"8k":      {d: 8000, f: format.Conversions.K},
		"4.97mi":  {d: 8000, f: format.Conversions.M},
		"18.5k":   {d: 18500, f: format.Conversions.K},
		"11.18mi": {d: 18000, f: format.Conversions.M},
	}

	for expected, test := range inputs {
		rv := format.Distance(test.d, test.f)
		if expected != rv {
			t.Error(fmt.Sprintf("Result %s doesn't match test %s", rv, expected))
		}
	}
}
