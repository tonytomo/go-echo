package middlewares

import (
	"github.com/labstack/echo/middleware"
)

var IsAutheticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
