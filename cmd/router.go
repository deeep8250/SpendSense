package main

import (
	"net/http"

	"github.com/deeep8250/SpendSense/auth/middleware"
	handler "github.com/deeep8250/SpendSense/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler *handler.AuthHandler, categoryHandler *handler.Categoryhandler, expenseHandler *handler.ExpenseHandler) {

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "server is running",
		})

	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.Middleware())
	//category
	{
		protected.GET("/category", categoryHandler.GetCategories)
		protected.POST("/category", categoryHandler.CreateCategory)

		// expenses
		protected.GET("/expense", expenseHandler.GetAllExpensesHandler)
		protected.GET("/expense/:id", expenseHandler.GetSingleExpenseHandler)
		protected.POST("/expense", expenseHandler.CreateExpenseHandler)
		protected.DELETE("/expense/:id", expenseHandler.DeleteExpenseHandler)
	}
}
