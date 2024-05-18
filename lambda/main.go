package main

import (
	"fmt"
	"lambda-func/app"
	"lambda-func/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ProtectectedHanlder(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Protected",
		StatusCode: http.StatusOK,
	}, nil
}

func apiGateway(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	myApp := app.NewApp()

	switch request.Path {
	case "/register":
		return myApp.ApiHandler.RegisterUserHandler(request)
	case "/login":
		return myApp.ApiHandler.LoginUserHandler(request)
	case "/protected":
		return middleware.ValidateJWTMiddleware(ProtectectedHanlder)(request)
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
