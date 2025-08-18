package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type RequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResposeDTO struct {
	Token string `json:"token"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TokenClaims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginService interface {
	Login(email string, pass string) (*string, error)
}
