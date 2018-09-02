package oauth

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/cognitoidentityiface"
	"github.com/strava/go.strava"
	"strconv"
)

type Service struct {
	StravaAuthenticator strava.OAuthAuthenticator
	CognitoIdentity     cognitoidentityiface.CognitoIdentityAPI
}

type AuthorizeOutput struct {
	Token      string `json:"token"`
	IdentityId string `json:"identity_id"`
}

func New() (*Service, error) {

	stravaAuth := strava.OAuthAuthenticator{}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}
	cognito := cognitoidentity.New(cfg)
	return &Service{StravaAuthenticator: stravaAuth, CognitoIdentity: cognito}, nil

}

func (s *Service) Authorize(code string) (*AuthorizeOutput, error) {

	resp, err := StravaAuthenticate(s.StravaAuthenticator, code)
	if err != nil {
		return nil, err
	}

	userId := strconv.FormatInt(resp.Athlete.Id, 10)
	identityOutput, err := CognitoGetIdentity(s.CognitoIdentity, userId)
	if err != nil {
		return nil, err
	}

	authOut := &AuthorizeOutput{
		Token:      identityOutput.Token,
		IdentityId: identityOutput.IdentityId,
	}

	return authOut, nil
}
