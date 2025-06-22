package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"biometric-data-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImprovedPatientController demonstrates the new controller pattern
type ImprovedPatientController struct {
	*EnhancedController[dto.PatientCreateDTO, dto.PatientUpdateDTO]
	patientService *service.ImprovedPatientService
}

// NewImprovedPatientController creates a new improved patient controller
func NewImprovedPatientController(patientService *service.ImprovedPatientService) *ImprovedPatientController {
	enhancedController := NewEnhancedController[dto.PatientCreateDTO, dto.PatientUpdateDTO](
		patientService, 
		"patient",
	)
	
	return &ImprovedPatientController{
		EnhancedController: enhancedController,
		patientService:     patientService,
	}
}

// All basic CRUD operations (Create, GetByID, GetAll, Update, Delete) are inherited
// from EnhancedController and work automatically with proper error handling and validation

// Domain-specific endpoints
func (pc *ImprovedPatientController) GetPatientByDNI(c *gin.Context) {
	dni := c.Param("dni")
	if dni == "" {
		utils.RespondBadRequest(c, "DNI parameter is required")
		return
	}

	patient, err := pc.patientService.GetPatientByDNI(dni)
	if err != nil {
		utils.RespondInternalError(c, "retrieve patient by DNI", err)
		return
	}

	if patient == nil {
		utils.RespondNotFound(c, "patient")
		return
	}

	utils.RespondWithData(c, http.StatusOK, gin.H{"patient": patient})
}

func (pc *ImprovedPatientController) GetAllPatientLocations(c *gin.Context) {
	locations, err := pc.patientService.GetAllPatientLocations()
	if err != nil {
		utils.RespondInternalError(c, "retrieve patient locations", err)
		return
	}

	utils.RespondWithData(c, http.StatusOK, gin.H{"locations": locations})
}

// Example of how pagination would work with the enhanced controller
func (pc *ImprovedPatientController) GetPatientsWithPagination(c *gin.Context) {
	// This uses the inherited GetPaginatedAll method from EnhancedController
	pc.GetPaginatedAll(c)
}

// Example of bulk operations
func (pc *ImprovedPatientController) CreateMultiplePatients(c *gin.Context) {
	// This uses the inherited CreateMultiple method from EnhancedController
	pc.CreateMultiple(c)
}

// Health check endpoint
func (pc *ImprovedPatientController) HealthCheck(c *gin.Context) {
	// This uses the inherited HealthCheck method from EnhancedController
	pc.EnhancedController.HealthCheck(c)
}