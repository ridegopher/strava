package webhook

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/snsiface"
	"os"
	"strconv"
)

// Event holds the data that is submitted to our webhook
type Event struct {
	AspectType     string `json:"aspect_type"`
	EventTime      int64  `json:"event_time"`
	ObjectId       int64  `json:"object_id"`
	ObjectType     string `json:"object_type"`
	OwnerId        int    `json:"owner_id"`
	SubscriptionId int64  `json:"subscription_id"`
	Updates        `json:"updates,omitempty"`
}

// Updates key/value pairs with extra info
type Updates struct {
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Private    string `json:"private,omitempty"`
	Authorized string `json:"authorized,omitempty"`
}

// EventToSNS sends the uploaded event to SNS to be processed by a different lambda
func EventToSNS(snsSvc snsiface.SNSAPI, event Event) (string, error) {

	message, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	params := &sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(os.Getenv("SNS_TOPIC")),
	}

	req := snsSvc.PublishRequest(params)
	output, err := req.Send()
	if err != nil {
		return "", err
	}

	return *output.MessageId, nil
}

// CheckSubscriptionId validate the request subscription id
func CheckSubscriptionId(event Event) error {
	if strconv.FormatInt(event.SubscriptionId, 10) != os.Getenv("STRAVA_SUBSCRIPTION_ID") {
		return errors.New("subscription id is not valid")
	}
	return nil
}
