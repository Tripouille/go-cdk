package api

import (
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	databaseStore database.DatabaseStore
}

func NewApiHandler(databaseStore database.DatabaseStore) ApiHandler {
	return ApiHandler{
		databaseStore: databaseStore,
	}
}

func (apiHandler ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var plainUser types.PlainUser

	err := json.Unmarshal([]byte(request.Body), &plainUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error unmarshalling request",
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("Error unmarshalling request: %w", err)
	}

	if plainUser.Username == "" || plainUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Username and password are required",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	userExists, err := apiHandler.databaseStore.DoesUserExist(plainUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error: %w", err)
	}

	if userExists {
		return events.APIGatewayProxyResponse{
			Body:       "User already exists",
			StatusCode: http.StatusConflict,
		}, nil
	}

	user, err := plainUser.ConvertToUser()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error: %w", err)
	}

	err = apiHandler.databaseStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "User registered successfully",
		StatusCode: http.StatusCreated,
	}, nil
}

func (apiHandler ApiHandler) LoginUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var plainUser types.PlainUser

	err := json.Unmarshal([]byte(request.Body), &plainUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error unmarshalling request",
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("Error unmarshalling request: %w", err)
	}

	user, err := apiHandler.databaseStore.GetUser(plainUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error: %w", err)
	}

	if user == nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid username or password",
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	passwordValidator := types.NewPasswordValidator(plainUser.Password, user.PasswordHash)
	err = passwordValidator.Validate()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid username or password",
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Login successful",
		StatusCode: http.StatusOK,
	}, nil
}
