package routes

import (
	"biometric-data-backend/controller"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DoctorsResource             = "doctors"
	ComorbiditiesResource       = "comorbidities"
	MedicationsResource         = "medications"
	BiometricsResource          = "biometrics"
	ComputerDiagnosticsResource = "computer-diagnostics"
	MonitoringDevicesResource   = "monitoring-devices"
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

	// Biometric
	biometricRepo := repository.NewBiometricRepository(db)
	biometricService := service.NewBiometricService(biometricRepo)
	biometricController := controller.NewBiometricController(biometricService)

	// Register biometric routes
	registerCrudRoutes(
		router,
		BiometricsResource,
		biometricController.CreateBiometric,
		biometricController.GetBiometricByID,
		biometricController.GetAllBiometrics,
		biometricController.UpdateBiometric,
		biometricController.DeleteBiometric,
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
}
