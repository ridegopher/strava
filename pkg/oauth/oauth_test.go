package oauth_test

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/oauth"
	"github.com/strava/go.strava"
	"testing"
)

func TestOauth_Authorize(t *testing.T) {
	t.Skipf("Run manually")
	fmt.Println(strava.ClientId)
	fmt.Println(strava.ClientSecret)
	auth, _ := oauth.New()
	token, _ := auth.Authorize("xxx")
	fmt.Println("token", token)
}
