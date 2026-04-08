package controllers

import "github.com/gin-gonic/gin"

func LoginAdmin(c *gin.Context) {

	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&loginData)

	c.JSON(200, gin.H{
		"message": "Login endpoint ready",
	})
}
