package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	USER_TABLE_NAME = "userTable"
)

type DynamoDBClient struct {
	dbStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		dbStore: db,
	}
}

func (dbClient DynamoDBClient) DoesUserExist(userbase string) (bool, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(userbase),
			},
		},
	}

	result, err := dbClient.dbStore.GetItem(getItemInput)

	if err != nil {
		return false, err
	}

	return result.Item != nil, nil
}

func (dbClient DynamoDBClient) CreateUser(user types.RegisterUser) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}

	_, err := dbClient.dbStore.PutItem(item)

	return err
}
