package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/eduufreire/url-shortner/internal/auth"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")[1]

		user, err := auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(402)
			w.Write([]byte(err.Error()))
			return
		}

		r.Header.Set("userId", strconv.Itoa(user.ID))
		next.ServeHTTP(w, r)
	})
}
