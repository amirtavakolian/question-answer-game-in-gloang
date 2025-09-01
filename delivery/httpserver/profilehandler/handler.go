package profilehandler

import (
	"QA-Game/services/profile"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	authSignKey []byte
}

func New(signKey string) *Handler {
	return &Handler{
		authSignKey: []byte(signKey),
	}
}

func (h *Handler) PlayerProfile(c echo.Context) error {

	profileService := profile.NewProfileService()

	result := profileService.GetPlayerProfile(c)

	return c.JSON(result.GetStatus(), result)
}
