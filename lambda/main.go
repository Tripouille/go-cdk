package main

import (
	"fmt"
	"lambda-func/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func apiGateway(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	myApp := app.NewApp()

	switch request.Path {
	case "/register":
		return myApp.ApiHandler.RegisterUserHandler(request)
	case "/login":
		return myApp.ApiHandler.LoginUserHandler(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Path %s not found", request.Path),
			StatusCode: http.StatusNotFound,
		}, nil
	}
}

func main() {
	lambda.Start(apiGateway)
}
