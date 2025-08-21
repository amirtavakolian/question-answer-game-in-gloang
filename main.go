package main

import (
	"QA-Game/services/auth"
	"QA-Game/services/profile"
	"net/http"
)

func main() {

	authService := auth.NewAuthService()
	profileService := profile.NewProfileService()

	http.HandleFunc("/auth/register", func(res http.ResponseWriter, req *http.Request) {
		result := authService.Register(req)
		res.Write([]byte(result))
	})

	http.HandleFunc("/auth/login", func(res http.ResponseWriter, req *http.Request) {
		result := authService.Login(req)
		res.Write([]byte(result))
	})

	http.HandleFunc("/player/profile", func(res http.ResponseWriter, req *http.Request) {
		result := profileService.GetPlayerProfile(req)
		res.Write([]byte(result))
	})

	http.ListenAndServe("127.0.0.1:8000", nil)
}
