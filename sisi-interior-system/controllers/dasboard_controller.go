package controllers

import (
	"net/http"
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {

	var totalProject int64
	var totalPortfolio int64
	var latestProjects []models.Project

	if err := config.DB.Model(&models.Project{}).Count(&totalProject).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data project"})
		return
	}

	//  total project
	config.DB.Model(&models.Project{}).Count(&totalProject)

	// total portfolio
	config.DB.Model(&models.Portfolio{}).Count(&totalPortfolio)

	//  project terbaru
	config.DB.Order("id desc").Limit(5).Find(&latestProjects)

	// data asli
	var chart []map[string]interface{}
	config.DB.Raw(`
		SELECT DATE(created_at) as date, SUM(harga_proyek) as total
		FROM data_proyek
		GROUP BY DATE(created_at)
		ORDER BY date ASC
`).Scan(&chart)

	c.JSON(http.StatusOK, gin.H{
		"total_project":   totalProject,
		"total_portfolio": totalPortfolio,
		"latest_projects": latestProjects,
		"chart":           chart,
	})
}
