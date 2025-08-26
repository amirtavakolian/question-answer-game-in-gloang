package httpserver

import (
	"QA-Game/delivery/httpserver/playerhandler"
	"QA-Game/delivery/httpserver/profilehandler"
		"QA-Game/delivery/httpserver/rolepermissionhandler/permissionhandler"
"QA-Game/delivery/httpserver/rolepermissionhandler/rolehandler"
"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

const signKey = "@@##AAAtt##$@#@%23432424asdsad345345SFD"

func (server *HttpServer) Serve() {

	playerHld := playerhandler.New()
	profileHld := profilehandler.New(signKey)
	roleHld := rolehandler.NewRoleHandler(signKey)
	permissionHld := permissionhandler.NewPermissionHandler(signKey)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	playerHld.SetPlayerRoutes(e)
	profileHld.SetProfileRoutes(e)
	roleHld.SetRoutes(e)
	permissionHld.SetPermissionRoutes(e)

	e.Start("127.0.0.1:8000")

}
