package categoryhandler

import "github.com/labstack/echo/v4"

func (h *Handler) SetCategoryRoutes(c *echo.Echo) {

	playerGroup := c.Group("/category")

	playerGroup.POST("/store", h.Store)
}
