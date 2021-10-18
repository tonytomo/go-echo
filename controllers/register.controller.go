package controllers

import (
	"fmt"
	"go-echo/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	var response models.Response

	res, err := models.Register(username, password)
	if !res {
		response.Status = false
		response.ErrorCode = fmt.Sprint(http.StatusInternalServerError)
		response.ErrorDescription = err.Error()
		response.Data = nil

		return c.JSON(http.StatusInternalServerError, response)
	}

	claims := &jwtCustomClaims{
		username,
		"application",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		response.Status = false
		response.ErrorCode = fmt.Sprint(http.StatusInternalServerError)
		response.ErrorDescription = err.Error()
		response.Data = nil

		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = true
	response.ErrorCode = "-"
	response.ErrorDescription = "-"
	response.Data = echo.Map{
		"username": username,
		"password": "",
		"token":    t,
	}

	return c.JSON(http.StatusOK, response)
}
