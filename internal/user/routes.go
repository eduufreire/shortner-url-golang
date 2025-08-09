package user

import (
	"net/http"
)

func Routes(h UserHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /", h.Create)

	return router
}
