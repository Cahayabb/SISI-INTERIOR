package controllers

import (
	"net/http"
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginAdmin(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var admin models.Admin

	// ambil data dari frontend
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// ambil data dari DB pakai GORM
	if err := config.DB.Where("username = ?", loginData.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// cek password
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte((loginData).Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user": gin.H{
			"username": admin.Username,
			"role":     "admin",
		},
	})
}
