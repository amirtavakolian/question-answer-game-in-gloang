package profilehandler

import "github.com/labstack/echo/v4"

func (h *Handler) SetProfileRoutes(c *echo.Echo) {

	c.GET("/player/profile", h.PlayerProfile)
}
