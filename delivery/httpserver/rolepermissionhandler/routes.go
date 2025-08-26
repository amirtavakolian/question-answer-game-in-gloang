package rolepermissionhandler

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoleRoutes(c *echo.Echo) {

	prefix := c.Group("/role")

	prefix.POST("/store", h.Store, mw.JWT(h.signKey))
}
