package httpdelivery

import (
	"QA-Game/services/auth"
	"github.com/labstack/echo/v4"
	)

func register(c echo.Context) error {

	authService := auth.NewAuthService()

	result := authService.Register(c)

	return c.JSON(result.GetStatus(), result)
}

func login(c echo.Context) error {

	authService := auth.NewAuthService()

	result := authService.Login(c)

	return c.JSON(result.GetStatus(), result)
}
