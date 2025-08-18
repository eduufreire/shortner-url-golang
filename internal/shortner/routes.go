package shortner

import (
	"net/http"

	"github.com/eduufreire/url-shortner/internal/middleware"
)

func Routes(handler ShortnerHandler) *http.ServeMux {
	router := http.NewServeMux()

	handleCreate := http.HandlerFunc(handler.Create)
	handleGet := http.HandlerFunc(handler.Get)
	router.Handle("POST /", middleware.Log(middleware.Authentication(handleCreate)))
	router.Handle("GET /{hash}", middleware.Authentication(handleGet))

	return router
}
