package application

import (
	"apitest/domain/entity"
	"apitest/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
}

var _ UserAppInterface = &userApp{}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}
