package oauth

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/cognitoidentityiface"
	"os"
)

var identityPoolId, linkedLogin string

func init() {

	identityPoolId = os.Getenv("COGNITO_IDENTITY_POOL_ID")
	if identityPoolId == "" {
		fmt.Println(errors.New("problem with env COGNITO_IDENTITY_POOL_ID"))
	}

	linkedLogin = os.Getenv("COGNITO_IDENTITY_LINKED_LOGIN")
	if linkedLogin == "" {
		fmt.Println(errors.New("problem with env COGNITO_IDENTITY_LINKED_LOGIN"))
	}
}

type CognitoGetIdentityOutput struct {
	IdentityId string `json:"identity_id"`
	Token      string `json:"token"`
}

func CognitoGetIdentity(svc cognitoidentityiface.CognitoIdentityAPI, UserId string) (*CognitoGetIdentityOutput, error) {

	logins := map[string]string{
		linkedLogin: UserId,
	}

	tokenDuration := int64(60 * 15)
	identityInput := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		IdentityPoolId: aws.String(identityPoolId),
		Logins:         logins,
		TokenDuration:  &tokenDuration,
	}

	cReq := svc.GetOpenIdTokenForDeveloperIdentityRequest(identityInput)

	cResp, err := cReq.Send()
	if err != nil {
		return nil, err
	}

	output := &CognitoGetIdentityOutput{
		IdentityId: *cResp.IdentityId,
		Token:      *cResp.Token,
	}
	return output, nil

}
