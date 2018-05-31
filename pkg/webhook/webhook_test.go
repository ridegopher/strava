package webhook_test

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/snsiface"
	"github.com/ridegopher/strava/pkg/webhook"
	"os"
	"testing"
)

var testEvent = webhook.Event{
	AspectType:     webhook.AspectTypeCreate,
	ObjectType:     webhook.ObjectTypeActivity,
	EventTime:      1516126040,
	ObjectId:       1360128428,
	OwnerId:        134815,
	SubscriptionId: 12345,
	Updates: webhook.Updates{
		Title: "Some Ride Yo",
	},
}

func TestWebhook_CheckSubscriptionIdOk(t *testing.T) {

	os.Setenv("STRAVA_SUBSCRIPTION_ID", "12345")
	//defer os.Unsetenv("STRAVA_SUBSCRIPTION_ID")

	svc, err := webhook.New(&testEvent)
	svc.SNS = &mockSNSClient{}

	if err != nil {
		t.Error(err.Error())
		return
	}

	err = svc.CheckSubscriptionId()
	if err != nil {
		t.Error(err.Error())
	}

}

func TestWebhook_CheckSubscriptionIdFail(t *testing.T) {

	os.Setenv("STRAVA_SUBSCRIPTION_ID", "1234500")
	//defer os.Unsetenv("STRAVA_SUBSCRIPTION_ID")

	svc, err := webhook.New(&testEvent)
	svc.SNS = &mockSNSClient{}
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = svc.CheckSubscriptionId()
	if err == nil {
		t.Error("Subscription is valid when it shouldn't be")
	}

}

// Mock SNS
type mockSNSClient struct {
	snsiface.SNSAPI
	Request sns.PublishRequest
}

func (m *mockSNSClient) PublishRequest(input *sns.PublishInput) sns.PublishRequest {
	return m.Request
}

func TestWebhook_ToSNS(t *testing.T) {

	cases := []struct {
		Request          sns.PublishRequest
		ExpectedResponse string
	}{
		{
			Request: sns.PublishRequest{
				Input: &sns.PublishInput{
					Message:  aws.String(`{"from":"test_1","to":"test_1","msg":"Hi test_1 :)"}`),
					TopicArn: aws.String("aws:sns:*:*:strava-events"),
					Subject:  aws.String("strava-webhook"),
				},
				Request: &aws.Request{
					Data: &sns.PublishOutput{
						MessageId: aws.String("uuid-1"),
					},
				},
			},
			ExpectedResponse: "uuid-1",
		},
	}

	for _, c := range cases {
		svc, err := webhook.New(&testEvent)
		svc.SNS = &mockSNSClient{Request: c.Request}
		if err != nil {
			t.Error(err.Error())
			return
		}

		response, err := svc.ToSNS()
		if err != nil {
			t.Error(err.Error())
			return
		}

		if response != c.ExpectedResponse {
			t.Error("Response doesn't match expected output")
		}

	}
}

func TestWebhook_CheckObjectType(t *testing.T) {

	if testEvent.ObjectType != webhook.ObjectTypeActivity {
		t.Error("Response doesn't match expected output")
	}
}
