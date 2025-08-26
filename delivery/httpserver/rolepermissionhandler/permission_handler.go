
package rolepermissionhandler

import (
	"QA-Game/param/permissionparam"
"QA-Game/services/rolepermission"
	"github.com/labstack/echo/v4"
)

type PermissionHandler struct {
	PermissionService rolepermission.PermissionService
	signKey     []byte
}

func NewPermissionHandler(signKey string) PermissionHandler {
	return PermissionHandler{
		PermissionService: rolepermission.NewPermissionService(),
		signKey:     []byte(signKey),
	}
}

func (h PermissionHandler) Store(c echo.Context) error {

	permissionParam := permissionparam.StorePermissionParam{}

	c.Bind(&permissionParam)

	result := h.PermissionService.Store(permissionParam)

	return c.JSON(result.GetStatus(), result)
}

