package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type CustomEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *CustomEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hello %s!", event.Name)

	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
