package repository

import "apitest/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
}
