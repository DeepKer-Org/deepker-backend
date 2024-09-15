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
	alertRepo := repository.NewAlertRepository(session)
	patientService := service.NewPatientService(patientRepo)
	alertService := service.NewAlertService(alertRepo)
	patientController := controllers.NewPatientController(patientService)
	alertController := controllers.NewAlertController(alertService)

	// Define your routes here
	router.POST("/patients", patientController.CreatePatient)
	router.GET("/patients", patientController.GetAllPatients)
	router.GET("/patients/:id", patientController.GetPatientByID)
	router.PUT("/patients/:id", patientController.UpdatePatient)
	router.DELETE("/patients/:id", patientController.DeletePatient)
	// Alerts
	router.POST("/alerts", alertController.CreateAlert)
	router.GET("/alerts", alertController.GetAllAlerts)
	router.GET("/alerts/:id", alertController.GetAlertByID)
	router.PUT("/alerts/:id", alertController.UpdateAlert)
	router.DELETE("/alerts/:id", alertController.DeleteAlert)
	router.GET("/alerts/patient_data/:id", alertController.GetAlertWithPatient)
}
