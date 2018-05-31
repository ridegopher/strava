package format_test

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/format"
	"testing"
)

func TestFormat_ElapsedTime(t *testing.T) {

	inputs := map[string]int{
		"244:23:54": 879834,
		"5:45":      345,
		"39:05":     2345,
	}

	for test, seconds := range inputs {
		rv := format.Time(seconds)
		if test != rv {
			t.Error(fmt.Sprintf("Result %s doesn't match test %s", rv, test))
		}
	}
}

func TestFormat_MovingTime(t *testing.T) {

	inputs := map[string]int{
		"24:26:24": 87984,
		"57:34":    3454,
		"22:25":    1345,
	}

	for test, seconds := range inputs {
		rv := format.Time(seconds)
		if test != rv {
			t.Error(fmt.Sprintf("Result %s doesn't match test %s", rv, test))
		}
	}
}
