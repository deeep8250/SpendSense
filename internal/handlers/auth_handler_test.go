package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deeep8250/SpendSense/internal/mocks"
	"github.com/deeep8250/SpendSense/models"
	"github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {

	tests := []struct {
		Name               string
		UserInput          models.Register
		ServiceError       error
		ExpectedStatusCode int
	}{

		{
			Name: "Success",
			UserInput: models.Register{
				Name:     "Deep mondal",
				Email:    "email@gmail.com",
				Password: "Deep@123",
			},
			ServiceError:       nil,
			ExpectedStatusCode: 201,
		},

		{
			Name: "User already exist",
			UserInput: models.Register{
				Name:     "Deep mondal",
				Email:    "email@gmail.com",
				Password: "Deep@123",
			},
			ServiceError:       errors.New("email already exists"),
			ExpectedStatusCode: 409,
		},

		{
			Name: "Password is too small",
			UserInput: models.Register{
				Name:     "Deep mondal",
				Email:    "email@gmail.com",
				Password: "Deep",
			},
			ServiceError:       nil,
			ExpectedStatusCode: 400,
		},

		{
			Name: "Invalid Email",
			UserInput: models.Register{
				Name:     "Deep mondal",
				Email:    "emailgmail.com",
				Password: "Deep",
			},
			ServiceError:       nil,
			ExpectedStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			mockService := mocks.ServiceMocking{
				RegisterFunc: func(userInput *models.Register) error {
					return tt.ServiceError
				},
			}

			handler := NewAuthHanler(&mockService)
			r := gin.Default()
			r.POST("/register", handler.Register)

			bodyBytes, _ := json.Marshal(tt.UserInput)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(string(bodyBytes)))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.ExpectedStatusCode {
				t.Errorf("expected this %d got this %d", tt.ExpectedStatusCode, w.Code)
			}

		})
	}

}

func TestLogin(t *testing.T) {
	tests := []struct {
		Name               string
		UserInput          models.Login
		Token              string
		ServiceError       error
		ExpectedStatusCode int
	}{

		{
			Name: "login Success",
			UserInput: models.Login{

				Email:    "deep@gmail.com",
				Password: "deep@123",
			},
			Token: "1234567890",

			ServiceError:       nil,
			ExpectedStatusCode: 200,
		},
		{
			Name: "incorrect password",
			UserInput: models.Login{

				Email:    "deep@gmail.com",
				Password: "de@123",
			},
			Token:              "",
			ServiceError:       errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"),
			ExpectedStatusCode: 401,
		},
		{
			Name: "invalid email",
			UserInput: models.Login{

				Email:    "deepgmail.com",
				Password: "deep@123",
			},
			Token:              "",
			ServiceError:       nil,
			ExpectedStatusCode: 400,
		},
		{
			Name: "email is not exist",
			UserInput: models.Login{

				Email:    "deep@gmail.com",
				Password: "deep@123",
			},
			Token:              "",
			ServiceError:       errors.New("user not found"),
			ExpectedStatusCode: 401,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			mockingService := mocks.ServiceMocking{
				LoginFunc: func(userInput *models.Login) (string, error) {
					return tt.Token, tt.ServiceError
				},
			}

			handler := NewAuthHanler(&mockingService)
			r := gin.Default()
			r.POST("/login", handler.Login)

			bodyBytes, _ := json.Marshal(tt.UserInput)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(string(bodyBytes)))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.ExpectedStatusCode {
				t.Errorf("expected %d got %d ", tt.ExpectedStatusCode, w.Code)
			}

		})

	}

}
