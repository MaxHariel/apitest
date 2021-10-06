package user

import (
	"apitest/application"
	"apitest/domain/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Users struct {
	us application.UserAppInterface
}

func NewUsers(us application.UserAppInterface) *Users {
	return &Users{
		us: us,
	}
}

func (s *Users) SaveUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newUser)
}
