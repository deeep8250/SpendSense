package services

import (
	"github.com/deeep8250/SpendSense/models"
)

type ServiceInterface interface {
	Register(userInput *models.Register) error
	Login(userInput *models.Login) (string, error)
}
