package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBStore struct {
	store *dynamodb.DynamoDB
}

func NewDynamoDBStore() DynamoDBStore {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBStore{
		store: db,
	}
}

func (databaseStore DynamoDBStore) DoesUserExist(username string) (bool, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := databaseStore.store.GetItem(getItemInput)

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
			"passwordHash": {
				S: aws.String(string(user.PasswordHash)),
			},
		},
	}

	_, err := databaseStore.store.PutItem(item)

	return err
}

func (databaseStore DynamoDBStore) GetUser(username string) (*types.User, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(USER_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := databaseStore.store.GetItem(getItemInput)

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return &types.User{
		Username:     username,
		PasswordHash: types.PasswordHash(*result.Item["passwordHash"].S),
	}, nil
}
