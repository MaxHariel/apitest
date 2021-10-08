package persistence

import (
	"apitest/domain/entity"
	"apitest/domain/repository"
	"errors"

	"gorm.io/gorm"
)

//UserRepo implements UserRepository
var _ repository.UserRepository = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

//NewUserRepository return instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, errors.New("error to create user")
	}

	return user, nil
}
