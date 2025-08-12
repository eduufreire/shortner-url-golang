package main

import (
	"net/http"

	"github.com/eduufreire/url-shortner/internal/dependencies"
)

func main() {
	dependencies := dependencies.Init().Wire()

	// antes da versao 1.22, os devs precisavam controlar o roteamento manualmente
	// incluindo o path, parametros, etc. agora é possivel fazer isso pelo proprio package net/http

	// server := &http.Server{
	// 	Addr: ":8080",
	// }
	http.Handle("/users/", http.StripPrefix("/users", dependencies.UserRoutes))
	http.Handle("/urls/", http.StripPrefix("/urls", dependencies.ShortnerRoutes))

	// http.HandleFunc("POST /auth/login", ah.Login)
	http.ListenAndServe(":8080", nil)
}
