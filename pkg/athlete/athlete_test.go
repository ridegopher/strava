package athlete_test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/ridegopher/strava/pkg/athlete"
	"reflect"
	"testing"
)

type mockDBService struct {
	dynamodbiface.DynamoDBAPI
	payload map[string]dynamodb.AttributeValue // Store expected return values
	err     error
}

// Mock GetItem such that the output returned carries values identical to input.
func (fd *mockDBService) GetItemRequest(input *dynamodb.GetItemInput) dynamodb.GetItemRequest {

	output := &dynamodb.GetItemOutput{
		Item: map[string]dynamodb.AttributeValue{},
	}
	for key, value := range fd.payload {
		output.Item[key] = value
	}
	req := dynamodb.GetItemRequest{
		Request: &aws.Request{
			Data:  output,
			Error: fd.err,
		},
	}

	return req
}

func TestAthlete_Get(t *testing.T) {

	payload := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String("1234567"),
		},
		"firstname": {
			S: aws.String("Chris"),
		},
		"lastname": {
			S: aws.String("Dornsife"),
		},
		"access_token": {
			S: aws.String("myTokenHere"),
		},
	}

	svc := &athlete.Service{
		DynamoDBAPI: &mockDBService{
			payload: payload,
		},
	}

	returnedAthlete, err := svc.Get(1234567)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	expectedAthlete := &athlete.Athlete{}
	err = dynamodbattribute.UnmarshalMap(payload, expectedAthlete)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(returnedAthlete, expectedAthlete) {
		t.Error("Expected Response was not equal")
	}

}

func TestAthlete_UpdateLocation(t *testing.T) {
	t.Skipf("Used for manual testing")

	svc, err := athlete.New()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	returnedAthlete, err := svc.Get(1234567)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("updating")
	svc.UpdateLocations(returnedAthlete, "123,456", "SomePlace Chicago IL")
}

func TestAthlete_GetNoMock(t *testing.T) {
	t.Skipf("Used for manual testing")
	svc, err := athlete.New()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	returnedAthlete, err := svc.Get(1234567)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	athlete, err := json.MarshalIndent(returnedAthlete, "", "\t")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(string(athlete))

}
