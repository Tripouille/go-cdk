package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	databaseStore database.DatabaseStore
}

func NewApiHandler(databaseStore database.DatabaseStore) ApiHandler {
	return ApiHandler{
		databaseStore: databaseStore,
	}
}

func (apiHandler ApiHandler) RegisterUserHandler(registerUser types.RegisterUser) error {
	if registerUser.Username == "" || registerUser.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	if userExists, err := apiHandler.databaseStore.DoesUserExist(registerUser.Username); err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user already exists")
	}

	return apiHandler.databaseStore.InsertUser(registerUser)
}
