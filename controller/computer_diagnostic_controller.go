package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type ComputerDiagnosticController struct {
	ComputerDiagnosticService service.ComputerDiagnosticService
}

func NewComputerDiagnosticController(computerDiagnosisService service.ComputerDiagnosticService) *ComputerDiagnosticController {
	return &ComputerDiagnosticController{
		ComputerDiagnosticService: computerDiagnosisService,
	}
}

// CreateComputerDiagnostic handles the creation of a new computer diagnosis
func (cdc *ComputerDiagnosticController) CreateComputerDiagnostic(c *gin.Context) {
	var diagnosisDTO dto.ComputerDiagnosticCreateDTO
	if err := c.ShouldBindJSON(&diagnosisDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cdc.ComputerDiagnosticService.CreateComputerDiagnostic(&diagnosisDTO)
	if err != nil {
		log.Printf("Failed to create computer diagnosis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create computer diagnosis"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Computer diagnosis created successfully", "diagnosis": diagnosisDTO})
}

// GetComputerDiagnosticByID handles retrieving a computer diagnosis by its DiagnosisID
func (cdc *ComputerDiagnosticController) GetComputerDiagnosticByID(c *gin.Context) {
	id := c.Param("id")

	diagnosisID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diagnosis ID"})
		return
	}

	diagnosis, err := cdc.ComputerDiagnosticService.GetComputerDiagnosticByID(diagnosisID)
	if err != nil {
		log.Printf("Error retrieving computer diagnosis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve computer diagnosis"})
		return
	}

	if diagnosis == nil {
		log.Printf("Computer diagnosis not found with DiagnosisID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Computer diagnosis not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"diagnosis": diagnosis})
}

// GetAllComputerDiagnostics handles retrieving all computer diagnostics
func (cdc *ComputerDiagnosticController) GetAllComputerDiagnostics(c *gin.Context) {
	diagnostics, err := cdc.ComputerDiagnosticService.GetAllComputerDiagnostics()
	if err != nil {
		log.Printf("Error retrieving computer diagnostics: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve computer diagnostics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"diagnostics": diagnostics})
}

// UpdateComputerDiagnostic handles updating an existing computer diagnosis
func (cdc *ComputerDiagnosticController) UpdateComputerDiagnostic(c *gin.Context) {
	id := c.Param("id")

	diagnosisID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diagnosis ID"})
		return
	}

	var diagnosisDTO dto.ComputerDiagnosticUpdateDTO
	if err := c.ShouldBindJSON(&diagnosisDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = cdc.ComputerDiagnosticService.UpdateComputerDiagnostic(diagnosisID, &diagnosisDTO)
	if err != nil {
		log.Printf("Failed to update computer diagnosis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update computer diagnosis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Computer diagnosis updated successfully", "diagnosis": diagnosisDTO})
}

// DeleteComputerDiagnostic handles deleting a computer diagnosis by its DiagnosisID
func (cdc *ComputerDiagnosticController) DeleteComputerDiagnostic(c *gin.Context) {
	id := c.Param("id")

	diagnosisID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diagnosis ID"})
		return
	}

	err = cdc.ComputerDiagnosticService.DeleteComputerDiagnostic(diagnosisID)
	if err != nil {
		log.Printf("Failed to delete computer diagnosis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete computer diagnosis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Computer diagnosis deleted successfully"})
}
