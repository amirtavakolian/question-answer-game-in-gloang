package rolepermissionhandler

import (
	"QA-Game/param/roleparam"
	"QA-Game/services/rolepermission"
	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	RoleService rolepermission.RoleService
	signKey     []byte
}

func NewRoleHandler(signKey string) RoleHandler {
	return RoleHandler{
		RoleService: rolepermission.New(),
		signKey:     []byte(signKey),
	}
}

func (h RoleHandler) Store(c echo.Context) error {

	roleParam := roleparam.StoreRoleParam{}

	c.Bind(&roleParam)

	result := h.RoleService.Store(roleParam)

	return c.JSON(result.GetStatus(), result)
}
