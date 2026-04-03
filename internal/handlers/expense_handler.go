package handler

import (
	"net/http"
	"strconv"

	"github.com/deeep8250/SpendSense/internal/services"
	"github.com/deeep8250/SpendSense/models"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	services *services.ExpenseService
}

func NewExpenseHandler(Service *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		services: Service,
	}
}

func (h *ExpenseHandler) CreateExpenseHandler(c *gin.Context) {

	var expense models.Expense
	err := c.ShouldBindJSON(&expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	expense.UserID = userID.(int)
	expense.Source = "manual"
	err = h.services.CreateExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "expense created",
	})

}

func (h *ExpenseHandler) GetSingleExpenseHandler(c *gin.Context) {

	ParamValue := c.Param("id")
	paramValueInt, err := strconv.Atoi(ParamValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	expense, err := h.services.GetExpensesByID(userID.(int), paramValueInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expense": expense,
	})

}

func (h *ExpenseHandler) DeleteExpenseHandler(c *gin.Context) {

	ParamValue := c.Param("id")
	expenseId, err := strconv.Atoi(ParamValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	result, err := h.services.DeleteExpense(userID.(int), expenseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": result,
	})

}

func (h *ExpenseHandler) GetAllExpensesHandler(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	expenses, err := h.services.GetExpenses(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}
