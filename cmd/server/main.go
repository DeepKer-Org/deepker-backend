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
	// Cargar el archivo .env
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	// Cargar la configuración de la base de datos y establecer la conexión
	config.LoadConfig()
	defer config.CloseDB() // Asegurarse de cerrar la conexión a PostgreSQL al finalizar

	// Crear un nuevo router de Gin
	router := gin.Default()

	// Inicializar Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registrar las rutas y pasar la conexión de la base de datos (PostgreSQL)
	routes.RegisterRoutes(router, config.DB)

	// Iniciar el servidor
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
