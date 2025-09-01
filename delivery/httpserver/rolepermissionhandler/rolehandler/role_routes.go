package rolehandler

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h RoleHandler) SetRoutes(c *echo.Echo) {

	rolePrefix := c.Group("/role")

	rolePrefix.POST("/store", h.Store, mw.JWT(h.signKey))
}
