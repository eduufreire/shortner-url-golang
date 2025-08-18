package auth

import (
	"encoding/json"
	"net/http"
)

type loginHandler struct {
	service LoginService
}

func NewLoginHandler(service LoginService) LoginHandler {
	return &loginHandler{
		service: service,
	}
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "erro login", 422)
		return
	}

	token, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		http.Error(w, "erro login", 422)
		return
	}

	parsed, err := json.Marshal(ResposeDTO{
		Token: *token,
	})
	w.Write(parsed)
}
