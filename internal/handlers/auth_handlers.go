package handler

import (
	"net/http"

	"github.com/deeep8250/SpendSense/internal/services"
	"github.com/deeep8250/SpendSense/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHanler(Service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: Service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	// recieve the input
	var user models.Register
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "user register successfully",
	})

}
