package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	dynamoDB := database.NewDynamoDB()
	apiHandler := api.NewApiHandler(dynamoDB)

	return App{
		ApiHandler: apiHandler,
	}
}
