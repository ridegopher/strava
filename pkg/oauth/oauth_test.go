package oauth_test

import (
	"fmt"
	"github.com/ridegopher/strava/pkg/oauth"
	"github.com/strava/go.strava"
	"testing"
)

//func TestOauth_MockAuthorize(t *testing.T) {

//t.Skipf("Used for manual testing")
//authorizer, err := oauth.New()
//if err != nil {
//	t.Fail()
//}
//out, err := authorizer.Authorize("xxx")
//if err != nil {
//	t.Fail()
//}
//
//fmt.Printf("%+v\n", out)
//}

func TestOauth_Authorize(t *testing.T) {

	fmt.Println(strava.ClientId)
	fmt.Println(strava.ClientSecret)
	auth, _ := oauth.New()
	token, _ := auth.Authorize("7b5954244dff1250bd1ca6b02e8841bb235eeaf5")
	fmt.Println("token", token)
}
