package user

import "net/http"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserService interface{
	Create(user RequestDTO) ResponseDTO
}

type UserRepository interface {
	Save(data RequestDTO) (*int64, error)
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}