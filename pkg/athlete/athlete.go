package athlete

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/ridegopher/strava/pkg/format"
	"github.com/strava/go.strava"
	"strconv"
)

type Athlete struct {
	Id                int                 `json:"id"`
	AccessToken       string              `json:"access_token"`
	UpdateDescription bool                `json:"update_description"`
	DateFormat        format.DateFormat   `json:"date_format"`
	Activities        map[string]Activity `json:"activities,omitempty"`
	Commutes          map[string]Commute  `json:"commutes,omitempty"`
	Locations         map[string]string   `json:"locations"`
}

type Activity struct {
	IsTrainer           bool `json:"trainer"`
	strava.ActivityType `json:"activity_type"`
	GearId              string              `json:"gear_id"`
	Private             bool                `json:"private"`
	Description         string              `json:"description"`
	NameFormats         []format.NameFormat `json:"name_formats"`
}

type Commute struct {
	Coordinates1 []float64           `json:"loc1"`
	Coordinates2 []float64           `json:"loc2"`
	Distance     float64             `json:"distance"`
	Activities   map[string]Activity `json:"activities"`
}

type Service struct {
	dynamodbiface.DynamoDBAPI
}

func New() (*Service, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(cfg)
	db := dynamodbiface.DynamoDBAPI(svc)

	return &Service{DynamoDBAPI: db}, nil

}

func (s *Service) Get(id int) (*Athlete, error) {
	aId := strconv.Itoa(id)

	input := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(aId),
			},
		},
		TableName: aws.String("Athletes"),
	}

	athlete := &Athlete{}

	req := s.DynamoDBAPI.GetItemRequest(input)
	if output, err := req.Send(); err == nil {
		err = dynamodbattribute.UnmarshalMap(output.Item, athlete)
		if err != nil {
			return nil, err
		}
	}

	return athlete, nil

}
