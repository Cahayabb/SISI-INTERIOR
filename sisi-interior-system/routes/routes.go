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

	api := r.Group("/api")
	{
		// ========================
		// PUBLIC ROUTES
		// ========================
		// untuk company profile
		api.GET("/portfolios", controllers.GetPortfolio)

		// estimasi boleh public
		api.POST("/estimasi-biaya", controllers.EstimasiBiaya)

		// AUTH
		api.POST("/login", controllers.LoginAdmin)

		// ========================
		// ADMIN ROUTES
		// ========================
		admin := api.Group("/admin")
		//admin.Use(middleware.AuthMiddleware())
		{
			// PORTFOLIO MANAGEMENT
			admin.POST("/portfolios", controllers.CreatePortfolio)
			admin.PUT("/portfolios/:id", controllers.UpdatePortfolio)
			admin.DELETE("/portfolios/:id", controllers.DeletePortfolio)

			// DASHBOARD
			admin.GET("/dashboard", controllers.GetDashboard)
		}
	}
}
