package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ridegopher/strava/pkg/webhook"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		webhookEvent := &webhook.Event{}
		err := json.Unmarshal([]byte(snsRecord.Message), webhookEvent)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
			return
		}

		fmt.Printf("%+v\n", webhookEvent)

		eventSvc, err := webhook.New(webhookEvent)
		if err != nil {
			fmt.Println("error processing webhook", err)
			return
		}

		message, err := eventSvc.ProcessEvent()
		if err != nil {
			fmt.Println("Error processing event", message, err)
			return
		}

		fmt.Println(message)

	}
}

func main() {
	lambda.Start(handler)
}
