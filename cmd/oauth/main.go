package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ridegopher/strava/pkg/oauth"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	code := request.QueryStringParameters["code"]
	if code == "" {
		return events.APIGatewayProxyResponse{Body: "Missing code from query string", StatusCode: 400}, nil
	}

	a, err := oauth.New()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	response, err := a.Authorize(code)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
	}

	resp := events.APIGatewayProxyResponse{
		Body: string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Origin":  "*",
		},
		StatusCode: 200,
	}

	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}
