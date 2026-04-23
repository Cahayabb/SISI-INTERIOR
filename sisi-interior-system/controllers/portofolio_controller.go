package controllers

import (
	"fmt"
	"sisi-interior-system/config"
	"sisi-interior-system/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePortfolio(c *gin.Context) {
	judul := c.PostForm("judul_proyek")
	deskripsi := c.PostForm("deskripsi")

	portfolio := models.Portfolio{
		JudulProyek: judul,
		Deskripsi:   deskripsi,
	}

	// upload file
	file, err := c.FormFile("gambar")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		c.SaveUploadedFile(file, "./uploads/"+filename)
		portfolio.Gambar = filename
	}

	if err := config.DB.Create(&portfolio).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal membuat portfolio: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Berhasil ditambahkan"})
}

func GetPortfolio(c *gin.Context) {
	var portfolio []models.Portfolio
	config.DB.Find(&portfolio)

	c.JSON(200, gin.H{"data": portfolio})
}

func UpdatePortfolio(c *gin.Context) {
	id := c.Param("id")

	var portfolio models.Portfolio

	if err := config.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	judul := c.PostForm("judul_proyek")
	deskripsi := c.PostForm("deskripsi")

	portfolio.JudulProyek = judul
	portfolio.Deskripsi = deskripsi

	file, err := c.FormFile("gambar")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		c.SaveUploadedFile(file, "./uploads/"+filename)
		portfolio.Gambar = filename
	}

	if err := config.DB.Save(&portfolio).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update portfolio: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Berhasil diupdate"})
}

func DeletePortfolio(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Portfolio{}, id)
	c.JSON(200, gin.H{"message": "Berhasil dihapus"})
}
