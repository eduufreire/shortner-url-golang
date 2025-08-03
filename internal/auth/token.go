package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("SigningTOkenKeySecurity")

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TokenClaims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

func createToken(id int, name string) string {

	claims := TokenClaims{
		User: User{ID: id, Name: name},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "root",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("deu ruim no token")
	}

	return tokenSigned
}

func VerifyToken(tokenString string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("erro em alguma coisa: %s", err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("token invalido")
	}

	user := User{}
	if claims, ok := token.Claims.(*TokenClaims); ok {
		user.ID = claims.User.ID
		user.Name = claims.User.Name
	}
	return &user, nil
}
