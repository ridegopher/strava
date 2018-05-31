package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ridegopher/strava/pkg/webhook"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	stravaEvent := webhook.Event{}
	err := json.Unmarshal([]byte(request.Body), &stravaEvent)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}

	svc, err := webhook.New(&stravaEvent)

	err = svc.CheckSubscriptionId()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	resp, err := svc.ToSNS()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	fmt.Println("SNS response:", resp)

	return events.APIGatewayProxyResponse{Body: "Don't buy upgrades, ride up grades --Eddy Merckx", StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
