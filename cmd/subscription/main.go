package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ridegopher/strava/pkg/subscription"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	svc, err := subscription.New(request.QueryStringParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	response, err := svc.VerifyToken()
	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
