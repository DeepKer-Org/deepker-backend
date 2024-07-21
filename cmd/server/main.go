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
	config.LoadConfig()

	router := gin.Default()
	// Initialize swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
