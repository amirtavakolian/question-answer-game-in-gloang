package httpserver

import (
	"QA-Game/delivery/httpserver/playerhandler"
	"QA-Game/delivery/httpserver/profilehandler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (server *HttpServer) Serve() {

	playerHld := playerhandler.New()
	profileHld := profilehandler.New()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	playerHld.SetPlayerRoutes(e)
	profileHld.SetProfileRoutes(e)


	e.Start("127.0.0.1:8000")

}
