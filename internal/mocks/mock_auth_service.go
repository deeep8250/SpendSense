package mocks

import (
	"github.com/deeep8250/SpendSense/models"
)

type ServiceMocking struct {
	RegisterFunc func(userInput *models.Register) error
	LoginFunc    func(userInput *models.Login) (string, error)
}

func (m *ServiceMocking) Register(userInput *models.Register) error {
	return m.RegisterFunc(userInput)
}

func (m *ServiceMocking) Login(userInput *models.Login) (string, error) {
	return m.LoginFunc(userInput)
}
