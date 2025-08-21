package main

import (
	"QA-Game/delivery/httpdelivery"
)

func main() {

	httpServer := httpdelivery.NewHttpServer()

	httpServer.Serve()
}
