package rolepermissionhandler

import (
	"QA-Game/param/roleparam"
	"QA-Game/services/rolepermission"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	RoleService rolepermission.Service
	signKey     []byte
}

func New(signKey string) Handler {
	return Handler{
		RoleService: rolepermission.New(),
		signKey:     []byte(signKey),
	}
}

func (h Handler) Store(c echo.Context) error {

	roleParam := roleparam.StoreRoleParam{}

	c.Bind(&roleParam)

	result := h.RoleService.Store(roleParam)

	return c.JSON(result.GetStatus(), result)
}
