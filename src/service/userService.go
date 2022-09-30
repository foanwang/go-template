package service

import (
	"testing/src/model"
	"testing/src/repository"
)

type IUserService interface {
	CreateUser(mu *model.User) (data *model.User, err error)
}

type UserService struct {
	repository.UserRepository
}

func (service *UserService) CreateUser(mu *model.User) (data *model.User, err error) {
	return nil, nil
}
