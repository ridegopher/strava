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

type Service struct {
	event Event
}

func New(input map[string]string) (*Service, error) {

	mode := input["hub.mode"]
	if mode == "" {
		return nil, errors.New("missing hub.mode")
	}

	verifyToken := input["hub.verify_token"]
	if verifyToken == "" {
		return nil, errors.New("missing hub.verify_token")
	}

	challenge := input["hub.challenge"]
	if challenge == "" {
		return nil, errors.New("missing hub.challenge")
	}

	service := &Service{
		Event{
			Mode:        mode,
			VerifyToken: verifyToken,
			Challenge:   challenge,
		},
	}

	return service, nil

}

// VerifyToken verifies a token provided to Strava during the subscription
func (s *Service) VerifyToken() (*Response, error) {

	token := os.Getenv("STRAVA_VERIFY_TOKEN")
	if token == "" {
		return nil, errors.New("missing env var STRAVA_VERIFY_TOKEN")
	}

	if s.event.Mode != "subscribe" || s.event.VerifyToken != token {
		return &Response{}, errors.New("you have been dropped")
	}

	return &Response{Challenge: s.event.Challenge}, nil

}
