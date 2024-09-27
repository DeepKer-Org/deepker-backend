package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type BiometricController struct {
	BiometricService service.BiometricService
}

func NewBiometricController(biometricService service.BiometricService) *BiometricController {
	return &BiometricController{
		BiometricService: biometricService,
	}
}

// CreateBiometric handles the creation of a new biometric record
func (bc *BiometricController) CreateBiometric(c *gin.Context) {
	var biometricDTO dto.BiometricCreateDTO
	if !bindJSON(c, &biometricDTO) {
		return
	}

	err := bc.BiometricService.CreateBiometric(&biometricDTO)
	if err != nil {
		log.Printf("Failed to create biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create biometric"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Biometric created successfully", "biometric": biometricDTO})
}

// GetBiometricByID handles retrieving a biometric record by its BiometricsID
func (bc *BiometricController) GetBiometricByID(c *gin.Context) {
	biometric, err := getByID(c, "id", bc.BiometricService.GetBiometricByID, "Biometric not found with BiometricsID: %v")
	if err != nil || biometric == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"biometric": biometric})
}

// GetAllBiometrics handles retrieving all biometric records
func (bc *BiometricController) GetAllBiometrics(c *gin.Context) {
	biometrics, err := bc.BiometricService.GetAllBiometrics()
	if err != nil {
		log.Printf("Error retrieving biometrics: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve biometrics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"biometrics": biometrics})
}

// UpdateBiometric handles updating an existing biometric record
func (bc *BiometricController) UpdateBiometric(c *gin.Context) {
	id := c.Param("id")

	// Parse the string to a UUID
	biometricID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid biometric ID"})
		return
	}

	var biometricDTO dto.BiometricUpdateDTO
	if !bindJSON(c, &biometricDTO) {
		return
	}

	err = bc.BiometricService.UpdateBiometric(biometricID, &biometricDTO)
	if err != nil {
		log.Printf("Failed to update biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update biometric"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biometric updated successfully", "biometric": biometricDTO})
}

// DeleteBiometric handles deleting a biometric record by its BiometricsID
func (bc *BiometricController) DeleteBiometric(c *gin.Context) {
	id := c.Param("id")

	// Parse the string to a UUID
	biometricID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid biometric ID"})
		return
	}

	err = bc.BiometricService.DeleteBiometric(biometricID)
	if err != nil {
		log.Printf("Failed to delete biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete biometric"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biometric deleted successfully"})
}
