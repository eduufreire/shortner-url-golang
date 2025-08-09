package main

import (
	"net/http"

	"github.com/eduufreire/url-shortner/internal/dependencies"
)

func main() {
	// db := database.CreateDatabase()
	// repo := shortner.Repository(db)
	// sh := shortner.Handler(repo)

	dependencies := dependencies.Init().Wire()

	// antes da versao 1.22, os devs precisavam controlar o roteamento manualmente
	// incluindo o path, parametros, etc. agora é possivel fazer isso pelo proprio package net/http

	// server := &http.Server{
	// 	Addr: ":8080",
	// }
	http.Handle("/users/", http.StripPrefix("/users", dependencies.UserRoutes))

	// http.HandleFunc("POST /shortners", sh.CreateUrl)
	// http.HandleFunc("POST /users", uh.CreateUser)
	// http.HandleFunc("GET /shortners/{hash}", sh.GetUrl)
	// http.HandleFunc("POST /auth/login", ah.Login)
	http.ListenAndServe(":8080", nil)
}
