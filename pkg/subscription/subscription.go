package subscription

import (
	"errors"
	"os"
)

// See: https://developers.strava.com/docs/webhooks/

// Response is returned to Strava if the verify token is valid
type Response struct {
	Challenge string `json:"hub.challenge"`
}

// Event the Strava request to verify a token
type Event struct {
	Mode        string `json:"hub.mode"`
	VerifyToken string `json:"hub.verify_token"`
	Challenge   string `json:"hub.challenge"`
}

// VerifyToken verifies a token provided to Strava during the subscription
func VerifyToken(input map[string]string) (*Response, error) {

	event := &Event{
		Mode:        input["hub.mode"],
		VerifyToken: input["hub.verify_token"],
		Challenge:   input["hub.challenge"],
	}

	if event.Mode != "subscribe" || event.VerifyToken != os.Getenv("STRAVA_VERIFY_TOKEN") {
		return &Response{}, errors.New("you have been dropped")
	}

	return &Response{Challenge: event.Challenge}, nil

}
