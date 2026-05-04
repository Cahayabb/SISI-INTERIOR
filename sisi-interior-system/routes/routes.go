package routes

import (
	"sisi-interior-system/controllers"
	"sisi-interior-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// ========================
	// ROOT
	// ========================
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SISI Interior System API Running",
		})
	})

	// ========================
	// API GROUP
	// ========================
	api := r.Group("/api")
	{
		// ========================
		// REGISTER
		// ========================
		api.POST("/register", controllers.RegisterUser)
		// ========================
		// PUBLIC ROUTES
		// ========================
		api.GET("/portfolios", controllers.GetPortfolio)

		// ========================
		// AUTH
		// ========================
		api.POST("/login", controllers.Login)

		// ========================
		// USERS ROUTES (ML ONLY)
		// ========================
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware("users"))
		{
			users.POST("/estimasi", controllers.EstimasiOnly)
		}

		// ========================
		// ADMIN ROUTES
		// ========================
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware("admin"))
		{
			// PORTFOLIO
			admin.POST("/portfolios", controllers.CreatePortfolio)
			admin.PUT("/portfolios/:id", controllers.UpdatePortfolio)
			admin.DELETE("/portfolios/:id", controllers.DeletePortfolio)

			// PROJECT
			admin.GET("/projects", controllers.GetProjects)
			admin.POST("/projects", controllers.CreateProject)
			admin.PUT("/projects/:id", controllers.UpdateProject)
			admin.DELETE("/projects/:id", controllers.DeleteProject)

			// DASHBOARD
			admin.GET("/dashboard", controllers.GetDashboard)

			// ESTIMASI (ADMIN DB)
			admin.POST("/estimasi", controllers.EstimasiBiaya)
			admin.GET("/estimasi", controllers.GetEstimasi)
			admin.DELETE("/estimasi/:id", controllers.DeleteEstimasi)
		}
	}
}
