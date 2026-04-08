package main

import (
	"sisi-interior-system/config"
	"sisi-interior-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8081")
}
