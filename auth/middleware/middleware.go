package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//grabing the token from the request header
		HeaderValue := c.GetHeader("Authorization")
		if HeaderValue == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized user",
			})
			c.Abort()
			return
		}
		//validating the token
		parts := strings.Split(HeaderValue, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "prefix isnt found in token",
			})
			c.Abort()
		}

		tokenString := parts[1]

		// Takes the header + claims from the incoming token
		// Combines them with your secret key (from the callback)
		// Runs the same HS256 algorithm to produce a new signature
		// Compares the new signature with the signature that came in the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !token.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		userID := int(claims["userID"].(float64))
		c.Set("userID", userID)
		c.Next()
	}

}
