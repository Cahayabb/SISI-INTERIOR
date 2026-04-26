package main

import (
	"sisi-interior-system/config"
	"sisi-interior-system/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}))

	r.Static("/uploads", "./uploads")
	routes.SetupRoutes(r)
	r.Run(":8081")
}
