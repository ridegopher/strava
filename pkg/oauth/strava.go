package oauth

import (
	"errors"
	"fmt"
	"github.com/strava/go.strava"
	"os"
	"strconv"
)

func init() {
	var err error

	strava.ClientId, err = strconv.Atoi(os.Getenv("STRAVA_CLIENT_ID"))
	if err != nil {
		fmt.Println(err)
	}

	strava.ClientSecret = os.Getenv("STRAVA_CLIENT_SECRET")
	if strava.ClientSecret == "" {
		fmt.Println(errors.New("problem with env STRAVA_CLIENT_SECRET"))
	}
}

func StravaAuthenticate(svc strava.OAuthAuthenticator, code string) (*strava.AuthorizationResponse, error) {

	resp, err := svc.Authorize(code, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
