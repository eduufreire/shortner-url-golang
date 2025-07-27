package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduufreire/url-shortner/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	ur user.UserRepository
}

func Handler(ur user.UserRepository) *handler {
	return &handler{
		ur: ur,
	}
}

func checkPassword(hashedPass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		fmt.Println("deu erro aqui kk")
		return false
	}
	return true
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "erro login", 422)
		return
	}

	user, err := h.ur.GetByEmail(body.Email)
	if err != nil {
		http.Error(w, "erro user", 422)
		return
	}

	if user.ID == 0 {
		http.Error(w, "not found", 404)
		return
	}

	isValidPassword := checkPassword(user.Password, body.Password)
	if !isValidPassword {
		fmt.Println("senhas diferente")
		return
	}

	token := createToken(user.ID, user.Name)

	parsed, err := json.Marshal(ResposeDTO{
		Token: token,
	})
	w.Write(parsed)
}
