package services

import (
	"github.com/wailman24/Go-chi-starter.git/internal/models"
	"github.com/wailman24/Go-chi-starter.git/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepositorie
}

func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepositorie(),
	}
}

func (us *UserService) CreateUser(user *models.User) error {
	err := us.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByEmail(user models.UserLogin) (*models.UserLogin, error) {
	res, err := us.repo.GetUserByEmail(&user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
