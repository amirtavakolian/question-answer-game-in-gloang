package profilehandler

import (
	"QA-Game/services/profile"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

// New returns a pointer to a new Handler.
func New() *Handler {
	return &Handler{}
}

func (h *Handler) PlayerProfile(c echo.Context) error {

	profileService := profile.NewProfileService()

	result := profileService.GetPlayerProfile(c)

	return c.JSON(result.GetStatus(), result)
}
