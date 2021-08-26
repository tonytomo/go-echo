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
		return c.String(http.StatusOK, "This is ECHO!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middlewares.IsAutheticated)
	e.POST("/pegawai", controllers.StorePegawai, middlewares.IsAutheticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middlewares.IsAutheticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middlewares.IsAutheticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
