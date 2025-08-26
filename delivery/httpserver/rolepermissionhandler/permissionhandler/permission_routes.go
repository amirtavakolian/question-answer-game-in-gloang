package permissionhandler

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h PermissionHandler) SetPermissionRoutes(c *echo.Echo) {

	permissionPrefix := c.Group("/permission", mw.JWT(h.signKey))

	permissionPrefix.POST("/store", h.Store)
	permissionPrefix.POST("/assign/role", h.AssignPermToRole)

}
