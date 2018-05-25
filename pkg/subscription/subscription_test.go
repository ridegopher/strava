package subscription_test

import (
	"github.com/ridegopher/strava/pkg/subscription"
	"os"
	"reflect"
	"testing"
)

func TestSubscription_SubscribeOK(t *testing.T) {

	os.Setenv("STRAVA_VERIFY_TOKEN", "TEST_VERIFY_TOKEN_VALUE")
	defer os.Unsetenv("STRAVA_VERIFY_TOKEN")

	challenge := "some-challenge-input"

	input := map[string]string{
		"hub.mode":         "subscribe",
		"hub.verify_token": "TEST_VERIFY_TOKEN_VALUE",
		"hub.challenge":    challenge,
	}

	response, err := subscription.VerifyToken(input)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !reflect.DeepEqual(response, &subscription.Response{Challenge: challenge}) {
		t.Error("Expected Response was not equal")
	}

}

func TestSubscription_SubscribeFail(t *testing.T) {

	os.Setenv("STRAVA_VERIFY_TOKEN", "TEST_VERIFY_TOKEN_VALUE_BAD")
	defer os.Unsetenv("STRAVA_VERIFY_TOKEN")

	challenge := "some-challenge-input"

	input := map[string]string{
		"hub.mode":         "subscribe",
		"hub.verify_token": "TEST_VERIFY_TOKEN_VALUE",
		"hub.challenge":    challenge,
	}

	response, _ := subscription.VerifyToken(input)

	if reflect.DeepEqual(response, &subscription.Response{Challenge: challenge}) {
		t.Error("Expected Response was equal, test failed")
	}

}
