package main

import (
	"QA-Game/delivery/httpserver"
)

func main() {

	httpServer := httpserver.NewHttpServer()

	httpServer.Serve()
}
