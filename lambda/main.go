package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string `json:"username"`
}

func HandleRequest(event Event) (*string, error) {
	if event.Username == "" {
		return nil, fmt.Errorf("username is required")
	}

	response := fmt.Sprintf("Hello, %s!", event.Username)

	return &response, nil
}

func main() {
	_ = app.NewApp()
	lambda.Start(HandleRequest)
}
