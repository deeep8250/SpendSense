package main

import (
	"net/http"

	authHandler "github.com/deeep8250/SpendSense/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler *authHandler.AuthHandler) {

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "server is running",
		})

	})

	r.POST("/register", authHandler.Register)
}
