package playerhandler

import "github.com/labstack/echo/v4"

func (h *Handler) SetPlayerRoutes(c *echo.Echo) {

	playerGroup := c.Group("/auth")

	playerGroup.POST("/register", h.PlayerRegister)

	playerGroup.POST("/login", h.PlayerLogin)
}
