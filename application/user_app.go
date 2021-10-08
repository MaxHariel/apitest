package application

import (
	"apitest/domain/entity"
	"apitest/domain/repository"
)

//UserApp implements UserAppInterface
var _ UserAppInterface = &userApp{}

type userApp struct {
	us repository.UserRepository
}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, error)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, error) {
	return u.us.SaveUser(user)
}
