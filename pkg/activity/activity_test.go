package activity_test

import (
	"testing"
)

func TestAthlete_GetNoMock(t *testing.T) {
	// I can't find a way to mock this since the http.Client is private.
	// The stub response client doesn't work.

	//result := `{"chris": "one"}`
	//client := strava.NewStubResponseClient(result, 200)
	//
	//stravaSvc := strava.NewAthletesService(client)
	//client := newCassetteClient(testToken, "activity_get")
	//activity, err := NewActivitiesService(client).Get(9460264).Do()
	//
	//fmt.Printf("%+v\n", activity)
	//fmt.Println(activity, err)

	//client := strava.NewStubResponseClient(`[{"id": 1000,"name": "Team Strava Cycling"}`, http.StatusOK)
	//clubs, _ := strava.NewClubsService(client).Get(1000).Do()
	//
	//fmt.Printf("%+v\n", clubs)

	//client := strava.NewStubResponseClient(`[{"id": 1,"name": "Team Strava Cycling"}`, http.StatusOK)
	//clubs, _ := strava.NewClubsService(client).Get(1000).Do()

}
