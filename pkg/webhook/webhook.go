package webhook

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/snsiface"
	"github.com/ridegopher/strava/pkg/activity"
	"os"
	"strconv"
)

type AspectType string

const (
	AspectTypeCreate AspectType = "create"
	AspectTypeUpdate AspectType = "update"
	AspectTypeDelete AspectType = "delete"
)

type ObjectType string

const (
	ObjectTypeAthlete  ObjectType = "athlete"
	ObjectTypeActivity ObjectType = "activity"
)

// Event holds the data that is submitted to our webhook
type Event struct {
	OwnerId        int   `json:"owner_id"`
	EventTime      int64 `json:"event_time"`
	ObjectId       int64 `json:"object_id"`
	SubscriptionId int64 `json:"subscription_id"`
	AspectType     `json:"aspect_type"`
	ObjectType     `json:"object_type"`
	Updates        `json:"updates,omitempty"`
}

// Updates key/value pairs with extra info
type Updates struct {
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Private    string `json:"private,omitempty"`
	Authorized string `json:"authorized,omitempty"`
}

type Service struct {
	SNS   snsiface.SNSAPI
	Event *Event
}

func New(event *Event) (*Service, error) {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}

	service := &Service{
		SNS:   sns.New(cfg),
		Event: event,
	}

	return service, nil

}

// EventToSNS sends the uploaded event to SNS to be processed by a different lambda
func (s *Service) ToSNS() (string, error) {

	message, err := json.Marshal(s.Event)
	if err != nil {
		return "", err
	}

	params := &sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(os.Getenv("SNS_TOPIC")),
	}

	req := s.SNS.PublishRequest(params)
	output, err := req.Send()
	if err != nil {
		return "", err
	}

	return *output.MessageId, nil
}

// CheckSubscriptionId validate the request subscription id
func (s *Service) CheckSubscriptionId() error {

	subId := os.Getenv("STRAVA_SUBSCRIPTION_ID")
	if subId == "" {
		return errors.New("missing env var $STRAVA_SUBSCRIPTION_ID")
	}

	subscriptionId, err := strconv.ParseInt(subId, 10, 64)
	if err != nil {
		return err
	}

	if s.Event.SubscriptionId != subscriptionId {
		return errors.New("subscription id is not valid")
	}

	return nil
}

type ProcessEventOutput string

func (s *Service) ProcessEvent() (ProcessEventOutput, error) {

	switch s.Event.ObjectType {

	case ObjectTypeActivity:

		switch s.Event.AspectType {

		case AspectTypeCreate:

			activitySvc, err := activity.New(s.Event.OwnerId, s.Event.ObjectId)
			if err != nil {
				return "", err
			}

			msg, err := activitySvc.ProcessActivityCreate()

			return ProcessEventOutput(msg), err

		case AspectTypeUpdate:
			return "activity type is update. nothing to do for update", nil

		case AspectTypeDelete:
			return "activity type is delete. nothing to do for delete", nil

		}
	case ObjectTypeAthlete:
		return "event object type is athlete. nothing to do for athletes", nil
	}

	return "success", nil

}
