package main

import (
	"apitest/infrastructure/database"
	"apitest/interfaces"
	"apitest/logger"

	_ "apitest/docs/swagger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}
}

// @title Swagger Example API
// @version 1.0
// @description Api

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	services, err := database.InitConfiguration()
	if err != nil {
		logger.Fatal(err)
	}

	if err = services.AutoMigrate(); err != nil {
		logger.Fatal(err)
	}

	e := echo.New()

	//Config Routes
	interfaces.Routes(e, services)

	e.Use(logger.LoggerWithConfig())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
