package user_test

import (
	"apitest/infrastructure/database"
	"apitest/interfaces"
	"apitest/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/assert.v1"
)

func TestUser(t *testing.T) {
	utils.LoadEnv()
	var (
		id   string
		user = map[string]interface{}{
			"name": utils.RandomString(10),
		}
	)

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

	t.Run("SaveUser", func(t *testing.T) {
		m, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(m))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		e.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)

		id = w.Body.String()

		fmt.Println("ID ---> ", id)

		assert.Equal(t, 201, w.Code)
	})

}
