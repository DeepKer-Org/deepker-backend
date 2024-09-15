package routes

import (
	"biometric-data-backend/controllers"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// RegisterRoutes sets up the application routes
func RegisterRoutes(router *gin.Engine, session *gocql.Session) {
	// Initialize repository, service, and controller
	patientRepo := repository.NewPatientRepository(session)
	patientService := service.NewPatientService(patientRepo)
	patientController := controllers.NewPatientController(patientService)

	// Define your routes here
	router.POST("/patients", patientController.CreatePatient)
	router.GET("/patients", patientController.GetAllPatients)
	router.GET("/patients/:id", patientController.GetPatientByID)
	router.PUT("/patients/:id", patientController.UpdatePatient)
	router.DELETE("/patients/:id", patientController.DeletePatient)
}
