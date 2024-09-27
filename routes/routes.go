package routes

import (
	"biometric-data-backend/controller"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DoctorsResource       = "doctors"
	ComorbiditiesResource = "comorbidities"
)

// registerCrudRoutes registers CRUD routes for a given resource
func registerCrudRoutes(router *gin.Engine, resource string, createFunc gin.HandlerFunc, getByIdFunc gin.HandlerFunc, getAllFunc gin.HandlerFunc, updateFunc gin.HandlerFunc, deleteFunc gin.HandlerFunc) {
	router.POST("/"+resource, createFunc)
	router.GET("/"+resource+"/:id", getByIdFunc)
	router.GET("/"+resource, getAllFunc)
	router.PUT("/"+resource+"/:id", updateFunc)
	router.DELETE("/"+resource+"/:id", deleteFunc)
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Doctor
	doctorRepo := repository.NewDoctorRepository(db)
	doctorService := service.NewDoctorService(doctorRepo)
	doctorController := controller.NewDoctorController(doctorService)

	// Register doctor routes
	registerCrudRoutes(
		router,
		DoctorsResource,
		doctorController.CreateDoctor,
		doctorController.GetDoctorByID,
		doctorController.GetAllDoctors,
		doctorController.UpdateDoctor,
		doctorController.DeleteDoctor,
	)
	// Additional doctor-specific route
	router.GET("/"+DoctorsResource+"/:id/short", doctorController.GetShortDoctorByID)

	// Comorbidity
	comorbidityRepo := repository.NewComorbidityRepository(db)
	comorbidityService := service.NewComorbidityService(comorbidityRepo)
	comorbidityController := controller.NewComorbidityController(comorbidityService)

	// Register comorbidity routes
	registerCrudRoutes(
		router,
		ComorbiditiesResource,
		comorbidityController.CreateComorbidity,
		comorbidityController.GetComorbidityByID,
		comorbidityController.GetAllComorbidities,
		comorbidityController.UpdateComorbidity,
		comorbidityController.DeleteComorbidity,
	)
}
