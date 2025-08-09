package user

import (
	"log"
)

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) Create(user RequestDTO) ResponseDTO {

	userExists, err := us.repository.GetByEmail(user.Email)
	if err != nil {
		log.Fatal("a")
	}

	if userExists.ID != 0 {
		log.Fatal("a")
	}

	hashedPass := hashPassword(user.Password)
	user.Password = hashedPass

	newID, err := us.repository.Save(user)
	if err != nil {
		log.Fatal("a")
	}

	createdUser, err := us.repository.GetByID(int(*newID))
	if err != nil {
		log.Fatal("a")
	}

	return ResponseDTO{
		ID:   createdUser.ID,
		Name: createdUser.Name,
	}
}
