package main

import (
	"net/http"

	"github.com/deeep8250/SpendSense/auth/middleware"
	handler "github.com/deeep8250/SpendSense/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler *handler.AuthHandler, categoryHandler *handler.Categoryhandler, expenseHandler *handler.ExpenseHandler, budgetHandler *handler.BudgetHandler) {

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

	{
		//category
		protected.GET("/categories", categoryHandler.GetCategories)
		protected.POST("/categories", categoryHandler.CreateCategory)

		// expenses
		protected.GET("/expenses", expenseHandler.GetAllExpensesByFilters)
		protected.GET("/expenses/:id", expenseHandler.GetSingleExpenseHandler)
		protected.POST("/expenses", expenseHandler.CreateExpenseHandler)
		protected.DELETE("/expenses/:id", expenseHandler.DeleteExpenseHandler)
		protected.POST("expenses/parser", expenseHandler.SetParserExpense)

		//budget
		protected.POST("/budgets", budgetHandler.CreateBudgetHandler)
		protected.GET("/budgets", budgetHandler.GetBudgetHandler)
		protected.GET("/budgets/summary", budgetHandler.SummaryRepoHandler)
		protected.GET("/budget/top_merchant", budgetHandler.TopMerchantHandler)
		protected.GET("/budget/trend", budgetHandler.TrendHandler)
		protected.GET("/budgets/alerts", budgetHandler.AlertHandler)

	}
}
