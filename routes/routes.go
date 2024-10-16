package routes

import (
	"biometric-data-backend/controller"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DoctorsResource             = "doctors"
	ComorbiditiesResource       = "comorbidities"
	MedicationsResource         = "medications"
	BiometricRecordsResource    = "biometrics"
	ComputerDiagnosticsResource = "computer-diagnostics"
	MonitoringDevicesResource   = "monitoring-devices"
	AlertsResource              = "alerts"
	PatientsResource            = "patients"
)

func CORSMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default configuration")
	}

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000"
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// registerCrudRoutes registers CRUD routes for a given resource
func registerCrudRoutes(router *gin.Engine, resource string, createFunc gin.HandlerFunc, getByIdFunc gin.HandlerFunc, getAllFunc gin.HandlerFunc, updateFunc gin.HandlerFunc, deleteFunc gin.HandlerFunc) {
	router.POST("/"+resource, createFunc)
	router.GET("/"+resource+"/:id", getByIdFunc)
	router.GET("/"+resource, getAllFunc)
	router.PATCH("/"+resource+"/:id", updateFunc)
	router.DELETE("/"+resource+"/:id", deleteFunc)
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Apply CORS middleware to the router
	router.Use(CORSMiddleware())

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
	router.GET("/"+DoctorsResource+"/alertID/:alertID", doctorController.GetDoctorsByAlertID)

	// Patient
	patientRepo := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepo)
	patientController := controller.NewPatientController(patientService)

	// Register patient routes
	registerCrudRoutes(
		router,
		PatientsResource,
		patientController.CreatePatient,
		patientController.GetPatientByID,
		patientController.GetAllPatients,
		patientController.UpdatePatient,
		patientController.DeletePatient,
	)
	// Additional patient-specific route
	router.GET("/"+PatientsResource+"/dni/:dni", patientController.GetPatientByDNI)

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

	// Medication
	medicationRepo := repository.NewMedicationRepository(db)
	medicationService := service.NewMedicationService(medicationRepo)
	medicationController := controller.NewMedicationController(medicationService)

	// Register medication routes
	registerCrudRoutes(
		router,
		MedicationsResource,
		medicationController.CreateMedication,
		medicationController.GetMedicationByID,
		medicationController.GetAllMedications,
		medicationController.UpdateMedication,
		medicationController.DeleteMedication,
	)

	// BiometricData
	biometricRepo := repository.NewBiometricDataRepository(db)
	biometricService := service.NewBiometricDataService(biometricRepo)
	biometricController := controller.NewBiometricDataController(biometricService)

	// Register biometric routes
	registerCrudRoutes(
		router,
		BiometricRecordsResource,
		biometricController.CreateBiometricData,
		biometricController.GetBiometricDataByID,
		biometricController.GetAllBiometricRecords,
		biometricController.UpdateBiometricData,
		biometricController.DeleteBiometricData,
	)

	// ComputerDiagnostic
	computerDiagnosticRepo := repository.NewComputerDiagnosticRepository(db)
	computerDiagnosticService := service.NewComputerDiagnosticService(computerDiagnosticRepo)
	computerDiagnosticController := controller.NewComputerDiagnosticController(computerDiagnosticService)

	// Register computer diagnostic routes
	registerCrudRoutes(
		router,
		ComputerDiagnosticsResource,
		computerDiagnosticController.CreateComputerDiagnostic,
		computerDiagnosticController.GetComputerDiagnosticByID,
		computerDiagnosticController.GetAllComputerDiagnostics,
		computerDiagnosticController.UpdateComputerDiagnostic,
		computerDiagnosticController.DeleteComputerDiagnostic,
	)

	// MonitoringDevice
	monitoringDeviceRepo := repository.NewMonitoringDeviceRepository(db)
	monitoringDeviceService := service.NewMonitoringDeviceService(monitoringDeviceRepo)
	monitoringDeviceController := controller.NewMonitoringDeviceController(monitoringDeviceService)

	// Register monitoring device routes
	registerCrudRoutes(
		router,
		MonitoringDevicesResource,
		monitoringDeviceController.CreateMonitoringDevice,
		monitoringDeviceController.GetMonitoringDeviceByID,
		monitoringDeviceController.GetAllMonitoringDevices,
		monitoringDeviceController.UpdateMonitoringDevice,
		monitoringDeviceController.DeleteMonitoringDevice,
	)

	// Alert
	alertRepo := repository.NewAlertRepository(db)
	alertService := service.NewAlertService(alertRepo, biometricRepo, computerDiagnosticRepo, doctorRepo, patientRepo)
	alertController := controller.NewAlertController(alertService)

	// Register alert routes
	registerCrudRoutes(
		router,
		AlertsResource,
		alertController.CreateAlert,
		alertController.GetAlertByID,
		alertController.GetAllAlerts,
		alertController.UpdateAlert,
		alertController.DeleteAlert,
	)
}
