package routes

import (
	"biometric-data-backend/config"
	"biometric-data-backend/controller"
	"biometric-data-backend/enums"
	"biometric-data-backend/middleware"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"biometric-data-backend/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DoctorsResource             = "doctors"
	ComorbiditiesResource       = "comorbidities"
	MedicationsResource         = "medications"
	BiometricDataResource       = "biometrics"
	ComputerDiagnosticsResource = "computer-diagnostics"
	MonitoringDevicesResource   = "monitoring-devices"
	AlertsResource              = "alerts"
	PatientsResource            = "patients"
	RolesResource               = "roles"
	AuthorizationResource       = "authorization"
	PhoneResource               = "phones"
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
		origin := c.Request.Header.Get("Origin")
		appOrigin := c.Request.Header.Get("X-App-Origin")

		// Allow the specific origin from environment variable or React Native app (with custom header)
		if origin == allowedOrigin || appOrigin == "ReactNativeApp" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-App-Origin")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// registerCrudRoutesWithMiddleware registers CRUD routes for a given resource
func registerCrudRoutesWithMiddleware(router *gin.Engine, resource string, createFunc gin.HandlerFunc, getByIdFunc gin.HandlerFunc, getAllFunc gin.HandlerFunc, updateFunc gin.HandlerFunc, deleteFunc gin.HandlerFunc, requiredRoles []string) {
	authorized := router.Group("/")
	authorized.Use(middleware.RoleAuthorization(requiredRoles))

	authorized.POST("/"+resource, createFunc)
	authorized.GET("/"+resource+"/:id", getByIdFunc)
	authorized.GET("/"+resource, getAllFunc)
	authorized.PATCH("/"+resource+"/:id", updateFunc)
	authorized.DELETE("/"+resource+"/:id", deleteFunc)
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize Redis cache manager
	cacheManager := redis.NewCacheManager(config.RedisClient, 5*time.Minute)
	// Apply CORS middleware to the router
	router.Use(CORSMiddleware())
	// JWT Auth
	authController := controller.NewAuthController()

	// Register JWT auth routes
	router.POST("/generate-token", authController.GenerateTokenEndpoint)

	// Role
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo, cacheManager)
	roleController := controller.NewRoleController(roleService)

	// Register role routes
	registerCrudRoutesWithMiddleware(
		router,
		RolesResource,
		roleController.CreateRole,
		roleController.GetRoleByID,
		roleController.GetAllRoles,
		roleController.UpdateRole,
		roleController.DeleteRole,
		enums.ToStringArray(enums.Admin),
	)
	// Additional role-specific route
	router.POST("/"+RolesResource+"/names", roleController.GetRolesByNames)

	// Authorization
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, roleRepo)
	authorizationController := controller.NewAuthorizationController(userService)

	// Register authorization routes
	router.POST("/"+AuthorizationResource+"/login", authorizationController.AuthenticateUser)

	// Register authorization routes with middleware
	registerCrudRoutesWithMiddleware(
		router,
		AuthorizationResource,
		authorizationController.RegisterUser,
		authorizationController.GetUserById,
		authorizationController.GetAllUsers,
		authorizationController.UpdateUser,
		authorizationController.DeleteUser,
		enums.ToStringArray(enums.Admin),
	)

	// Doctor
	doctorRepo := repository.NewDoctorRepository(db)
	doctorService := service.NewDoctorService(doctorRepo, userRepo, roleRepo, userService, cacheManager)
	doctorController := controller.NewDoctorController(doctorService)

	// Register doctor routes
	registerCrudRoutesWithMiddleware(
		router,
		DoctorsResource,
		doctorController.CreateDoctor,
		doctorController.GetDoctorByID,
		doctorController.GetAllDoctors,
		doctorController.UpdateDoctor,
		doctorController.DeleteDoctor,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)
	// Additional doctor-specific route
	router.GET("/"+DoctorsResource+"/:id/short", doctorController.GetShortDoctorByID)
	router.GET("/"+DoctorsResource+"/alertID/:alertID", doctorController.GetDoctorsByAlertID)
	router.GET("/"+DoctorsResource+"/userID/:userID", doctorController.GetDoctorByUserID)
	router.PATCH("/"+DoctorsResource+"/userID/:userID", doctorController.UpdateDoctorByUserID)

	// Change password for doctor
	router.PATCH("/"+AuthorizationResource+"/change-password", doctorController.ChangePassword)
	// Patient
	patientRepo := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepo, cacheManager)
	patientController := controller.NewPatientController(patientService)

	// Register patient routes
	registerCrudRoutesWithMiddleware(
		router,
		PatientsResource,
		patientController.CreatePatient,
		patientController.GetPatientByID,
		patientController.GetAllPatients,
		patientController.UpdatePatient,
		patientController.DeletePatient,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)
	// Additional patient-specific route
	router.GET("/"+PatientsResource+"/dni/:dni", patientController.GetPatientByDNI)

	// Comorbidity
	comorbidityRepo := repository.NewComorbidityRepository(db)
	comorbidityService := service.NewComorbidityService(comorbidityRepo, cacheManager)
	comorbidityController := controller.NewComorbidityController(comorbidityService)

	// Register comorbidity routes
	registerCrudRoutesWithMiddleware(
		router,
		ComorbiditiesResource,
		comorbidityController.CreateComorbidity,
		comorbidityController.GetComorbidityByID,
		comorbidityController.GetAllComorbidities,
		comorbidityController.UpdateComorbidity,
		comorbidityController.DeleteComorbidity,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)

	// Medication
	medicationRepo := repository.NewMedicationRepository(db)
	medicationService := service.NewMedicationService(medicationRepo, cacheManager)
	medicationController := controller.NewMedicationController(medicationService)

	// Register medication routes
	registerCrudRoutesWithMiddleware(
		router,
		MedicationsResource,
		medicationController.CreateMedication,
		medicationController.GetMedicationByID,
		medicationController.GetAllMedications,
		medicationController.UpdateMedication,
		medicationController.DeleteMedication,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)

	// BiometricData
	biometricRepo := repository.NewBiometricDataRepository(db)
	biometricService := service.NewBiometricDataService(biometricRepo, cacheManager)
	biometricController := controller.NewBiometricDataController(biometricService)

	// Register biometric routes
	registerCrudRoutesWithMiddleware(
		router,
		BiometricDataResource,
		biometricController.CreateBiometricData,
		biometricController.GetBiometricDataByID,
		biometricController.GetAllBiometricData,
		biometricController.UpdateBiometricData,
		biometricController.DeleteBiometricData,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)

	// ComputerDiagnostic
	computerDiagnosticRepo := repository.NewComputerDiagnosticRepository(db)
	computerDiagnosticService := service.NewComputerDiagnosticService(computerDiagnosticRepo, cacheManager)
	computerDiagnosticController := controller.NewComputerDiagnosticController(computerDiagnosticService)

	// Register computer diagnostic routes
	registerCrudRoutesWithMiddleware(
		router,
		ComputerDiagnosticsResource,
		computerDiagnosticController.CreateComputerDiagnostic,
		computerDiagnosticController.GetComputerDiagnosticByID,
		computerDiagnosticController.GetAllComputerDiagnostics,
		computerDiagnosticController.UpdateComputerDiagnostic,
		computerDiagnosticController.DeleteComputerDiagnostic,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)

	// MonitoringDevice
	monitoringDeviceRepo := repository.NewMonitoringDeviceRepository(db)
	monitoringDeviceService := service.NewMonitoringDeviceService(monitoringDeviceRepo, cacheManager)
	monitoringDeviceController := controller.NewMonitoringDeviceController(monitoringDeviceService)

	// Register monitoring device routes
	registerCrudRoutesWithMiddleware(
		router,
		MonitoringDevicesResource,
		monitoringDeviceController.CreateMonitoringDevice,
		monitoringDeviceController.GetMonitoringDeviceByID,
		monitoringDeviceController.GetAllMonitoringDevices,
		monitoringDeviceController.UpdateMonitoringDevice,
		monitoringDeviceController.DeleteMonitoringDevice,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)

	// Phone
	phoneRepo := repository.NewPhoneRepository(db)
	phoneService := service.NewPhoneService(phoneRepo)
	phoneController := controller.NewPhoneController(phoneService)

	// Register phone routes
	router.POST("/"+PhoneResource, phoneController.CreatePhone)

	// Alert
	alertRepo := repository.NewAlertRepository(db)
	alertService := service.NewAlertService(alertRepo, biometricRepo, computerDiagnosticRepo, doctorRepo, monitoringDeviceRepo, phoneRepo, patientRepo, cacheManager)
	alertController := controller.NewAlertController(alertService)

	// Register alert routes
	registerCrudRoutesWithMiddleware(
		router,
		AlertsResource,
		alertController.CreateAlert,
		alertController.GetAlertByID,
		alertController.GetAllAlerts,
		alertController.UpdateAlert,
		alertController.DeleteAlert,
		enums.ToStringArray(enums.Admin, enums.Doctor),
	)
}
