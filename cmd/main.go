package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//dependency injection
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "server is running",
		})

	})
	r.Run(":8080")

}
