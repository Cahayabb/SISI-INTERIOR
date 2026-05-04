package controllers

import (
	"net/http"
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LOGIN GABUNGAN ADMIN + USER
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// bind request
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// =========================
	// 1. CEK ADMIN
	// =========================
	var admin models.Admin
	errAdmin := config.DB.Where("username = ?", loginData.Username).First(&admin).Error

	if errAdmin == nil {
		// cek password admin
		err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginData.Password))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login berhasil",
				"user": gin.H{
					"username": admin.Username,
					"role":     "admin",
				},
			})
			return
		}
	}

	// =========================
	// 2. CEK USER
	// =========================
	var user models.User
	errUser := config.DB.Where("username = ?", loginData.Username).First(&user).Error

	if errUser == nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login berhasil",
				"user": gin.H{
					"username": user.Username,
					"role":     "users",
				},
			})
			return
		}
	}

	// =========================
	// 3. GAGAL LOGIN
	// =========================
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Username atau password salah",
	})
}
