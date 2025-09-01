package permissionhandler

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/services/rolepermission"
	"github.com/labstack/echo/v4"
)

type PermissionHandler struct {
	PermissionService rolepermission.PermissionService
	signKey           []byte
}

func NewPermissionHandler(signKey string) PermissionHandler {
	return PermissionHandler{
		PermissionService: rolepermission.NewPermissionService(),
		signKey:           []byte(signKey),
	}
}

func (h PermissionHandler) Store(c echo.Context) error {

	permissionParam := rolepermissionparam.StorePermissionParam{}

	c.Bind(&permissionParam)

	result := h.PermissionService.Store(permissionParam)

	return c.JSON(result.GetStatus(), result)
}

func (h PermissionHandler) AssignPermToRole(c echo.Context) error {

	assignPermToRoleParams := rolepermissionparam.AssignPermissionToRoleParam{}

	c.Bind(&assignPermToRoleParams)

	result := h.PermissionService.PermToRole(assignPermToRoleParams)

	return c.JSON(result.GetStatus(), result)
}
