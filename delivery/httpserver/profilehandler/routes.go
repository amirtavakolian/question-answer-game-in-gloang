package profilehandler

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SetProfileRoutes(c *echo.Echo) {

	c.GET("/player/profile", h.PlayerProfile, mw.JWT(h.authSignKey))
}
