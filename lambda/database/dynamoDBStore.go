package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBStore struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBStore() DynamoDBStore {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBStore{
		databaseStore: db,
	}
}

func (databaseStore DynamoDBStore) DoesUserExist(userbase string) (bool, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(userbase),
			},
		},
	}

	result, err := databaseStore.databaseStore.GetItem(getItemInput)

	if err != nil {
		return false, err
	}

	return result.Item != nil, nil
}

func (databaseStore DynamoDBStore) InsertUser(user types.User) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.PasswordHash.PasswordHash),
			},
		},
	}

	_, err := databaseStore.databaseStore.PutItem(item)

	return err
}
