package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format token salah",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalid",
			})
			c.Abort()
			return
		}

		// OPTIONAL: simpan role ke context (kalau mau dipakai nanti)
		c.Set("role", role)

		c.Next()
	}
}
