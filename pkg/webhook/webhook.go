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

// Event
type Event struct {
	AspectType     string `json:"aspect_type"`
	EventTime      int64  `json:"event_time"`
	ObjectId       int64  `json:"object_id"`
	ObjectType     string `json:"object_type"`
	OwnerId        int    `json:"owner_id"`
	SubscriptionId int64  `json:"subscription_id"`
	Updates        `json:"updates,omitempty"`
}

// Updates
type Updates struct {
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Private    string `json:"private,omitempty"`
	Authorized string `json:"authorized,omitempty"`
}

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

func CheckSubscriptionId(event Event) error {
	if strconv.FormatInt(event.SubscriptionId, 10) != os.Getenv("STRAVA_SUBSCRIPTION_ID") {
		return errors.New("subscription id is not valid")
	}
	return nil
}
