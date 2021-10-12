package routes

import (
	"go-echo/controllers"
	"go-echo/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there!")
	})

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.POST("/register", controllers.Register)

	e.GET("/data", controllers.FetchDataPbb, middlewares.IsAutheticated)

	return e
}
