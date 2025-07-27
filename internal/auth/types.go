package auth

type RequestDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ResposeDTO struct {
	Token string `json:"token"`
}
