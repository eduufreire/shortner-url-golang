package middleware

import (
	"net/http"

	"github.com/eduufreire/url-shortner/internal/logger"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("estou na area da requisicoa")
		next.ServeHTTP(w, r)
	})
}
