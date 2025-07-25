package main

import (
	"net/http"

	"github.com/eduufreire/url-shortner/internal/database"
	"github.com/eduufreire/url-shortner/internal/shortner"
	"github.com/eduufreire/url-shortner/internal/user"
)

func main() {
	db := database.CreateDatabase()
	repo := shortner.Repository(db)
	sh := shortner.Handler(repo)

	userRepo := user.Repository(db)
	uh := user.Handler(userRepo)

	http.HandleFunc("POST /shortners", sh.CreateUrl)
	http.HandleFunc("GET /shortners/{hash}", sh.GetUrl)

	http.HandleFunc("POST /users", uh.CreateUser)
	http.HandleFunc("POST /users/login", uh.LoginTeste)

	http.ListenAndServe(":8080", nil)
}
