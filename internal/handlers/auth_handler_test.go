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
