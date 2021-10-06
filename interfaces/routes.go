package interfaces

import (
	"apitest/infrastructure/database"
	"apitest/interfaces/user"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, services *database.Repositories) {
	users := user.NewUsers(services.User)
	e.POST("/users", users.SaveUser)
}
