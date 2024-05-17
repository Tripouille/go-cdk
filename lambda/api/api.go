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

func (apiHandler ApiHandler) RegisterUserHandler(plainUser types.PlainUser) error {
	if plainUser.Username == "" || plainUser.Password.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	if userExists, err := apiHandler.databaseStore.DoesUserExist(plainUser.Username); err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user already exists")
	}

	user, err := plainUser.ConvertToUser()

	if err != nil {
		return err
	}

	return apiHandler.databaseStore.InsertUser(*user)
}
