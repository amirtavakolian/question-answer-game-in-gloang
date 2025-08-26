package httpserver

import (
	"QA-Game/delivery/httpserver/playerhandler"
	"QA-Game/delivery/httpserver/profilehandler"
	"QA-Game/delivery/httpserver/rolepermissionhandler"
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
	rolepermissionHld := rolepermissionhandler.New(signKey)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	playerHld.SetPlayerRoutes(e)
	profileHld.SetProfileRoutes(e)
	rolepermissionHld.SetRoleRoutes(e)

	e.Start("127.0.0.1:8000")

}
