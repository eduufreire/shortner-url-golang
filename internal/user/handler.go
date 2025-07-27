package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repository *repository
}

func Handler(repository *repository) *handler {
	return &handler{
		repository: repository,
	}
}

func hashPassword(pass string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 1)
	if err != nil {
		fmt.Println("deu erro aqui kk")
	}
	return string(hashedPass)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	userExists, err := h.repository.GetByEmail(body.Email)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	if userExists.ID != 0 {
		http.Error(w, "user already exists", 400)
		return
	}

	hashedPass := hashPassword(body.Password)
	body.Password = hashedPass

	newID, err := h.repository.Save(body)
	fmt.Println(int(*newID))
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	user, err := h.repository.GetByID(int(*newID))
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	userResponse := ResponseDTO{
		ID:   user.ID,
		Name: user.Name,
	}

	response, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	user, err := h.repository.GetByID(parsedID)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	if user.ID == 0 {
		http.Error(w, "not found", 404)
		return
	}

	userResponse := ResponseDTO{
		ID:   user.ID,
		Name: user.Name,
	}

	response, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


