package handler

import (
	"net/http"

	"github.com/deeep8250/SpendSense/internal/services"
	"github.com/deeep8250/SpendSense/models"
	"github.com/gin-gonic/gin"
)

type Categoryhandler struct {
	services *services.CategoryService
}

func NewCategoryHandler(Service *services.CategoryService) *Categoryhandler {
	return &Categoryhandler{
		services: Service,
	}
}

func (h *Categoryhandler) GetCategories(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	categories, err := h.services.GetCategories(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *Categoryhandler) CreateCategory(c *gin.Context) {

	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	uid := userID.(int)
	category.UserID = &uid
	err = h.services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "category created",
	})

}
