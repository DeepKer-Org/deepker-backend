package routes

import (
	"biometric-data-backend/controllers"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	alertRepo := repository.NewAlertRepository(db)

	alertService := service.NewAlertService(alertRepo)

	alertController := controllers.NewAlertController(alertService)

	router.POST("/alerts", alertController.CreateAlert)
	router.GET("/alerts", alertController.GetAllAlerts)
	router.GET("/alerts/:id", alertController.GetAlertByID)
	router.PUT("/alerts/:id", alertController.UpdateAlert)
	router.DELETE("/alerts/:id", alertController.DeleteAlert)
}
