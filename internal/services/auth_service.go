package services

import (
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(Repo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repo: Repo,
	}
}

func (s *AuthService) Register(userInput *models.Register) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:           userInput.Name,
		Email:          userInput.Email,
		HashedPassword: string(hashedPass),
	}

	err = s.repo.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil

}
