package routes

import (
	"biometric-data-backend/controller"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	doctorRepo := repository.NewDoctorRepository(db)
	doctorService := service.NewDoctorService(doctorRepo)
	doctorController := controller.NewDoctorController(doctorService)

	router.POST("/doctors", doctorController.CreateDoctor)
	router.GET("/doctors/:id", doctorController.GetDoctorByID)
	router.GET("/doctors/:id/short", doctorController.GetShortDoctorByID)
	router.GET("/doctors", doctorController.GetAllDoctors)
	router.PUT("/doctors/:id", doctorController.UpdateDoctor)
	router.DELETE("/doctors/:id", doctorController.DeleteDoctor)
}
