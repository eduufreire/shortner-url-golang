package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

func hashPassword(pass string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 1)
	if err != nil {
		fmt.Println("deu erro aqui kk")
	}
	return string(hashedPass)
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := RequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	createdUser := uh.service.Create(body)

	response, err := json.Marshal(createdUser)
	if err != nil {
		http.Error(w, "erro ao cadastrar novo usuario", 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// func (uh *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

// 	id := r.PathValue("id")
// 	parsedID, err := strconv.Atoi(id)
// 	if err != nil {
// 		http.Error(w, "erro ao cadastrar novo usuario", 422)
// 		return
// 	}

// 	user, err := uh.repository.GetByID(parsedID)
// 	if err != nil {
// 		http.Error(w, "erro ao cadastrar novo usuario", 422)
// 		return
// 	}

// 	if user.ID == 0 {
// 		http.Error(w, "not found", 404)
// 		return
// 	}

// 	userResponse := ResponseDTO{
// 		ID:   user.ID,
// 		Name: user.Name,
// 	}

// 	response, err := json.Marshal(userResponse)
// 	if err != nil {
// 		http.Error(w, "erro ao cadastrar novo usuario", 422)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(response)
// }
