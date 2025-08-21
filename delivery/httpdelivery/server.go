package httpdelivery

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

	e.POST("/auth/register", register)
	e.POST("/auth/login", login)

	e.GET("/player/profile", getPlayerProfile)

	e.Start("127.0.0.1:8000")

}
