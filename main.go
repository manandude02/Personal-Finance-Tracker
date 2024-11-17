package main

import (
	"personal-finance-tracker/config"
	"personal-finance-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDatabase()

	// Set up Gin router
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
