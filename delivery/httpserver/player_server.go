package httpserver

import (
	"QA-Game/services/auth"
	"github.com/labstack/echo/v4"
)

func (server *HttpServer) playerRegister(c echo.Context) error {

	authService := auth.NewAuthService()

	result := authService.Register(c)

	return c.JSON(result.GetStatus(), result)
}

func (server *HttpServer) playerLogin(c echo.Context) error {

	authService := auth.NewAuthService()

	result := authService.Login(c)

	return c.JSON(result.GetStatus(), result)
}
