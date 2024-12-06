package main

import (
	"log"

	"HerHourlyLiftBackend/config"
	"HerHourlyLiftBackend/db"
	"HerHourlyLiftBackend/middleware"
	"HerHourlyLiftBackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to MongoDB
	db.ConnectMongoDB()

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Register routes
	routes.RegisterQuoteRoutes(router)

	// Start server
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
