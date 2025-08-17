package http

import (
	"QA-Game/services/auth"
	"net/http"
)

type Server struct {
}

func (server Server) Handle() {

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/register", func(res http.ResponseWriter, req *http.Request) {
		auth.NewAuthService().Register(req)
	})

}
