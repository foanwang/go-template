package repository

import (
	"fmt"
	"testing/src/model"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(interface{}) (interface{}, error)
	FindById(uint) error
	DeleteById(uint) error
}

// var db = config.GetDB()

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Save(user *model.User) (data *model.User, err error) {
	if err := u.DB.Create(user).Error; err != nil {
		fmt.Println("error != nil user:", user, " err:", err)
		return nil, err
	}
	fmt.Println("error == nil user:", user, " err:", err)
	return user, nil
}

func (u *UserRepository) FindById(id uint) (err error) {
	if err := u.DB.Where(&model.User{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteById(id uint) (err error) {
	if err := u.DB.Delete(&model.User{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
