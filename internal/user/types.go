package user

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
