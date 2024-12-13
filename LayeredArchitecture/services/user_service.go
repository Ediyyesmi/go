package services

import (
	"errors"
	"project/models"
	"project/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) RegisterUser(name, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     name,
		Password: string(hashedPassword),
	}
	return service.Repo.CreateUser(user)
}

func (service *UserService) AuthenticateUser(name, password string) (bool, error) {
	user, err := service.Repo.GetUserByName(name)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}
	return true, nil
}
