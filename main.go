package main

import (
	"QA-Game/services/auth"
	"net/http"
)

func main() {

	authService := auth.NewAuthService()

	http.HandleFunc("/auth/register", func(res http.ResponseWriter, req *http.Request) {
		result := authService.Register(req)
		res.Write([]byte(result))
	})

	

	http.ListenAndServe("127.0.0.1:8000", nil)

}
