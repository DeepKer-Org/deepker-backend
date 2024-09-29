package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type BiometricDataController struct {
	BiometricDataService service.BiometricDataService
}

func NewBiometricDataController(biometricService service.BiometricDataService) *BiometricDataController {
	return &BiometricDataController{
		BiometricDataService: biometricService,
	}
}

// CreateBiometricData handles the creation of a new biometric record
func (bc *BiometricDataController) CreateBiometricData(c *gin.Context) {
	var biometricDTO dto.BiometricDataCreateDTO
	if !bindJSON(c, &biometricDTO) {
		return
	}

	err := bc.BiometricDataService.CreateBiometricData(&biometricDTO)
	if err != nil {
		log.Printf("Failed to create biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create biometric"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "BiometricData created successfully", "biometric": biometricDTO})
}

// GetBiometricDataByID handles retrieving a biometric record by its BiometricDatasID
func (bc *BiometricDataController) GetBiometricDataByID(c *gin.Context) {
	biometric, err := getByID(c, "id", bc.BiometricDataService.GetBiometricDataByID, "BiometricData not found with BiometricDatasID: %v")
	if err != nil || biometric == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"biometric": biometric})
}

// GetAllBiometricRecords handles retrieving all biometric records
func (bc *BiometricDataController) GetAllBiometricRecords(c *gin.Context) {
	biometrics, err := bc.BiometricDataService.GetAllBiometricRecords()
	if err != nil {
		log.Printf("Error retrieving biometrics: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve biometrics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"biometrics": biometrics})
}

// UpdateBiometricData handles updating an existing biometric record
func (bc *BiometricDataController) UpdateBiometricData(c *gin.Context) {
	id := c.Param("id")

	// Parse the string to a UUID
	biometricID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid biometric ID"})
		return
	}

	var biometricDTO dto.BiometricDataUpdateDTO
	if !bindJSON(c, &biometricDTO) {
		return
	}

	err = bc.BiometricDataService.UpdateBiometricData(biometricID, &biometricDTO)
	if err != nil {
		log.Printf("Failed to update biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update biometric"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BiometricData updated successfully", "biometric": biometricDTO})
}

// DeleteBiometricData handles deleting a biometric record by its BiometricDatasID
func (bc *BiometricDataController) DeleteBiometricData(c *gin.Context) {
	id := c.Param("id")

	// Parse the string to a UUID
	biometricID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid biometric ID"})
		return
	}

	err = bc.BiometricDataService.DeleteBiometricData(biometricID)
	if err != nil {
		log.Printf("Failed to delete biometric: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete biometric"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BiometricData deleted successfully"})
}
