package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (server *HttpServer) Serve() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth/playerRegister", server.playerRegister)
	e.POST("/auth/playerLogin", server.playerLogin)

	e.GET("/player/profile", server.playerProfile)

	e.Start("127.0.0.1:8000")

}
