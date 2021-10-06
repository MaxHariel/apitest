package persistence

import (
	"apitest/domain/entity"
	"apitest/domain/repository"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Create(&user).Error
	if err != nil {
		dbErr["error"] = "Error to create user"
	}

	return user, nil
}
