package oauth_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/cognitoidentityiface"
	"github.com/ridegopher/strava/pkg/oauth"
	"testing"
)

var (
	mockIdentityId = "us-east-1:identityId"
	mockToken      = "mock-token-here"
	mockUserId     = "123456"
)

type mockCognitoIdentityClient struct {
	cognitoidentityiface.CognitoIdentityAPI
}

func (m *mockCognitoIdentityClient) GetOpenIdTokenForDeveloperIdentityRequest(input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) cognitoidentity.GetOpenIdTokenForDeveloperIdentityRequest {

	return cognitoidentity.GetOpenIdTokenForDeveloperIdentityRequest{
		Request: &aws.Request{
			Data: &cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput{
				IdentityId: &mockIdentityId,
				Token:      &mockToken,
			},
		},
		Input: input,
	}

}

func TestOauth_CognitoGetIdentity(t *testing.T) {

	mockClient := &mockCognitoIdentityClient{}
	out, err := oauth.CognitoGetIdentity(mockClient, mockUserId)
	if err != nil {
		t.Error(err)
	}

	if out.IdentityId != mockIdentityId {
		t.Error(fmt.Sprintf("expected IdentityId %s got %s", mockIdentityId, out.IdentityId))
	}

	if out.Token != mockToken {
		t.Error(fmt.Sprintf("expected Token %s got %s", mockToken, out.Token))
	}

}
