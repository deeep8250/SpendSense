package services

import (
	"fmt"
	"strings"

	"github.com/deeep8250/SpendSense/auth/jwt"
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
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("email already exists")
		}
		return err
	}
	return nil

}

func (s *AuthService) Login(userInput *models.Login) (string, error) {
	user, err := s.repo.LoginUser(userInput)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userInput.Password))
	if err != nil {
		return "", err
	}

	jwt, err := jwt.CreateJWT(user.Id)
	if err != nil {
		return "", nil
	}

	return jwt, nil

}
