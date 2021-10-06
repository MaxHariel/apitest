package main

import (
	"apitest/infrastructure/database"
	"apitest/interfaces"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		echo.New().StdLogger.Fatal("no .env")
	}
}

func main() {

	services, err := database.InitConfiguration()
	if err != nil {
		echo.New().StdLogger.Fatal(err)
	}

	if err = services.AutoMigrate(); err != nil {
		echo.New().StdLogger.Fatal(err)
	}

	e := echo.New()

	//Config Routes
	interfaces.Routes(e, services)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}
