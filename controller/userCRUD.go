package controller

import (
	"errors"
	"gorm.io/gorm"
	"zugSystem/model"
)

type user interface {
	getUser(id uint) (user model.User, err error)
	getUsers() []model.User
	createUser(user model.User)
	deleteUser(user model.User)
	updateUser(user model.User)
}

type userImpl struct {
	Db *gorm.DB
}

func (u userImpl) getUser(id uint) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, id)
	if result != nil {
		return user, nil
	}
	return user, errors.New("Cannot find user")
}

func (u userImpl) getUsers() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	return result
}
