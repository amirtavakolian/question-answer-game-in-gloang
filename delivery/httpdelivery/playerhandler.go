package httpdelivery
import (

"QA-Game/services/profile"
"github.com/labstack/echo/v4"
)

func getPlayerProfile(c echo.Context) error {

	profileService := profile.NewProfileService()

	result := profileService.GetPlayerProfile(c)

	return c.JSON(result.GetStatus(), result)
}