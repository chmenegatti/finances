package services

import (
	"finances/src/config"
	"finances/src/models"
	"finances/src/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(user *models.User) error {
	hashedPassword, err := config.EncryptPwd(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepository.Create(user)
}

func (s *userService) GetUser(id string) (*models.User, error) {
	return s.userRepository.Get(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}
