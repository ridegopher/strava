package location

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/format"
	"os"
	"testing"
)

func TestFormat_Location(t *testing.T) {

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		t.Skip("skipping test; $GOOGLE_API_KEY not set")
	}

	locSvc, err := format.NewLocationService()
	if err != nil {
		t.Error(err)
		return
	}

	inputs := map[string][]float64{
		"River North Chicago IL":          {41.89, -87.64},
		"Wrightwood Neighbors Chicago IL": {41.93, -87.66},
	}

	for place, coords := range inputs {
		rv, err := locSvc.GetLocation(coords[0], coords[1])
		if err != nil {
			t.Error(err)
			return
		}
		if place != rv {
			t.Error(fmt.Sprintf("Result %s doesn't match test %s", rv, place))
		}
	}

}
