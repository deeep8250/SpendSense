package handler

import (
	"net/http"
	"strconv"

	"github.com/deeep8250/SpendSense/internal/services"
	"github.com/deeep8250/SpendSense/models"
	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	BudgetServices *services.BudgetService
}

func NewBudgetHandler(BS *services.BudgetService) *BudgetHandler {
	return &BudgetHandler{
		BudgetServices: BS,
	}
}

func (h *BudgetHandler) CreateBudgetHandler(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	var Budget models.Budget
	err := c.ShouldBindJSON(&Budget)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inputs",
		})
		return
	}

	Budget.UserID = userID.(int)

	err = h.BudgetServices.CreateBudgetService(Budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "budget created",
	})

}
func (h *BudgetHandler) GetBudgetHandler(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	budgets, err := h.BudgetServices.GetBudgetService(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"budgets": budgets,
	})

}

func (h *BudgetHandler) SummaryRepoHandler(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	Month := c.Query("month")
	MonthInt, err := strconv.Atoi(Month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Year := c.Query("year")
	YearInt, err := strconv.Atoi(Year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	SummaryR, err := h.BudgetServices.SummaryRepoService(userID.(int), MonthInt, YearInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"summary_report": SummaryR,
	})

}

func (h *BudgetHandler) TopMerchantHandler(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	topMer, err := h.BudgetServices.TopMerchantService(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Top_merchat": topMer,
	})

}

func (h *BudgetHandler) TrendHandler(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	Trends, err := h.BudgetServices.TrendService(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"trends": Trends,
	})

}
