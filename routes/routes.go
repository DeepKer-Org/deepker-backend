package routes

import (
	"biometric-data-backend/config"
	"biometric-data-backend/controllers"
	"biometric-data-backend/models"
	"biometric-data-backend/repositories"
	"biometric-data-backend/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	testRepository := repositories.NewRepository[*models.Test](config.DB)
	testService := services.NewService[*models.Test](testRepository)
	testController := controllers.NewController[*models.Test](testService)

	testRoutes := router.Group("/tests")
	{
		testRoutes.GET("/", testController.GetAll)
		testRoutes.GET("/:id", testController.GetByID)
		testRoutes.POST("/", testController.Create)
		testRoutes.PUT("/:id", testController.Update)
		testRoutes.DELETE("/:id", testController.Delete)
	}

	userRepository := repositories.NewRepository[*models.User](config.DB)
	userService := services.NewService[*models.User](userRepository)
	userController := controllers.NewController[*models.User](userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userController.GetAll)
		userRoutes.GET("/:id", userController.GetByID)
		userRoutes.POST("/", userController.Create)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

}
