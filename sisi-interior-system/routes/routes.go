package routes

import (
	"sisi-interior-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SISI Interior System API Running",
		})
	})

	r.POST("api/login", controllers.LoginAdmin)
	r.POST("/api/portfolio", controllers.CreatePortfolio)
	r.GET("/api/portfolio", controllers.GetPortfolio)
	r.PUT("/api/portfolio/:id", controllers.UpdatePortfolio)
	r.DELETE("/api/portfolio/:id", controllers.DeletePortfolio)
	r.POST("/estimasi-biaya", controllers.EstimasiBiaya)
}
