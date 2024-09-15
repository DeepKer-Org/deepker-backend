package main

import (
	"biometric-data-backend/config"
	"biometric-data-backend/docs"
	"biometric-data-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	// Load database configuration and create the session
	config.LoadConfig()
	defer config.CloseSession()

	// Create a new Gin router
	router := gin.Default()

	// Initialize swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes and pass the session
	routes.RegisterRoutes(router, config.Session)

	// Start the server
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
