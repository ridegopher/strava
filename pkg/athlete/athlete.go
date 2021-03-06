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
	MeasurementPref   string              `json:"measurement_preference"`
	DateFormat        format.DateFormat   `json:"date_format"`
	Activities        map[string]Activity `json:"activities,omitempty"`
	Commutes          map[string]Commute  `json:"commutes,omitempty"`
	Locations         map[string]string   `json:"locations"`
}

type Activity struct {
	strava.ActivityType `json:"activity_type"`
	IsTrainer           bool                `json:"trainer"`
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
	table *string
}

func New() (*Service, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(cfg)
	db := dynamodbiface.DynamoDBAPI(svc)

	return &Service{DynamoDBAPI: db, table: aws.String("Athletes")}, nil

}

func (s *Service) Get(id int) (*Athlete, error) {
	aId := strconv.Itoa(id)

	input := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(aId),
			},
		},
		TableName: s.table,
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

func (s *Service) UpdateLocations(athlete *Athlete, locationKey, location string) error {
	input := &dynamodb.UpdateItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(athlete.Id)),
			},
		},
		UpdateExpression: aws.String("set locations.#key = :location"),
		ExpressionAttributeNames: map[string]string{
			"#key": locationKey,
		},
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":location": {
				S: aws.String(location),
			},
		},
		ConditionExpression: aws.String("attribute_not_exists(locations.#key)"),

		ReturnValues: "UPDATED_NEW",
		TableName:    s.table,
	}

	req := s.DynamoDBAPI.UpdateItemRequest(input)
	if _, err := req.Send(); err == nil {
		if err != nil {
			return err
		}
	}
	return nil
}
