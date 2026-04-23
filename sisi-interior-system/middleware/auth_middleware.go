package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//ambil header Authorization
		authHeader := c.GetHeader("Authorization")

		//kalau tidak ada token
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		//format harus: Bearer token
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format token salah",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		//kalau token kosong
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalid",
			})
			c.Abort()
			return
		}

		//lanjut ke controller
		c.Next()
	}
}
