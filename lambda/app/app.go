package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	dynamoDBStore := database.NewDynamoDBStore()
	apiHandler := api.NewApiHandler(dynamoDBStore)

	return App{
		ApiHandler: apiHandler,
	}
}
