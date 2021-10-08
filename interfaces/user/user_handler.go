package user

import (
	"apitest/application"
	"apitest/domain/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Users struct {
	ua application.UserAppInterface
}

func NewUsers(ua application.UserAppInterface) *Users {
	return &Users{
		ua: ua,
	}
}

// adicionar godoc
// @Summary Add new user
// @Description Add new user
// @Tags /users
// @Accept json
// @Produce json
// @Param tag body entity.UserDTO true "User Data"
// @Success 201 {object} entity.User "User Data"
// @Router /users [post]
func (u *Users) SaveUser(c echo.Context) error {
	var ud entity.UserDTO
	if err := c.Bind(&ud); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid_json")
	}
	user := entity.User{
		Name: ud.Name,
	}
	newUser, err := u.ua.SaveUser(&user)
	if err != nil {
		fmt.Println("------------------->", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}
