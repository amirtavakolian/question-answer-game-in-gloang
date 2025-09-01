package assignhandler

import (
	"QA-Game/param/rolepermissionparam"
	"QA-Game/services/rolepermission"
	"github.com/labstack/echo/v4"
)

type AssignHandler struct {
	AssignService rolepermission.AssignService
}

