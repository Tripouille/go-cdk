package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (apiHandler ApiHandler) RegisterUserHandler(registerUser types.RegisterUser) error {
	if registerUser.Username == "" || registerUser.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	if userExists, err := apiHandler.dbStore.DoesUserExist(registerUser.Username); err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user already exists")
	}

	return apiHandler.dbStore.CreateUser(registerUser)
}
