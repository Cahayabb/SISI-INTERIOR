package controllers

import (
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "input tidak valid"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	user := models.User{
		Username: input.Username,
		Password: string(hash),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "register berhasil"})
}
