package shortner

import "net/http"

func Routes(handler ShortnerHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /", handler.Create)
	router.HandleFunc("GET /{hash}", handler.Get)

	return router
}
