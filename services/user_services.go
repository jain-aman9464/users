package services

import (
	"github.com/federicoleon/users/domain/users"
	"github.com/federicoleon/users/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	return nil, nil
}
