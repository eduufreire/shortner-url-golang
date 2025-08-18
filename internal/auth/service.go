package auth

import (
	"fmt"

	"github.com/eduufreire/url-shortner/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type loginService struct {
	userRepository user.UserRepository
}

func NewLoginService(userRepository user.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (ls *loginService) Login(email string, pass string) (*string, error) {
	user, err := ls.userRepository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("erro ao pegar usuario no banco no momento do login")
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("usuario nao existe")
	}

	isValidPassword := checkPassword(user.Password, pass)
	if !isValidPassword {
		return nil, fmt.Errorf("senhas diferentes")
	}

	token := createToken(user.ID, user.Name)

	return &token, nil
}

func checkPassword(hashedPass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		fmt.Println("deu erro aqui kk")
		return false
	}
	return true
}
