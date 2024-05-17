package database

import (
	"lambda-func/types"
)

const (
	USER_TABLE_NAME = "userTable"
)

type UserStore interface {
	DoesUserExist(userbase string) (bool, error)
	InsertUser(user types.RegisterUser) error
}
