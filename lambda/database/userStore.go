package database

import (
	"lambda-func/types"
)

const (
	USER_TABLE_NAME = "userTable"
)

type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user types.User) error
	GetUser(username string) (*types.User, error)
}
