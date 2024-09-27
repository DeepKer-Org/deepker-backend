package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type MedicationController struct {
	MedicationService service.MedicationService
}

func NewMedicationController(medicationService service.MedicationService) *MedicationController {
	return &MedicationController{
		MedicationService: medicationService,
	}
}

// CreateMedication handles the creation of a new medication
func (mc *MedicationController) CreateMedication(c *gin.Context) {
	var medicationDTO dto.MedicationCreateDTO
	if err := c.ShouldBindJSON(&medicationDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := mc.MedicationService.CreateMedication(&medicationDTO)
	if err != nil {
		log.Printf("Failed to create medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create medication"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Medication created successfully", "medication": medicationDTO})
}

// GetMedicationByID handles retrieving a medication by its MedicationID
func (mc *MedicationController) GetMedicationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf("Invalid medication MedicationID: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication MedicationID"})
		return
	}

	medication, err := mc.MedicationService.GetMedicationByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve medication"})
		return
	}

	if medication == nil {
		log.Printf("Medication not found with MedicationID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medication": medication})
}

// GetAllMedications handles retrieving all medications
func (mc *MedicationController) GetAllMedications(c *gin.Context) {
	medications, err := mc.MedicationService.GetAllMedications()
	if err != nil {
		log.Printf("Error retrieving medications: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve medications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medications": medications})
}

// UpdateMedication handles updating an existing medication
func (mc *MedicationController) UpdateMedication(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf("Invalid medication MedicationID: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication MedicationID"})
		return
	}

	var medicationDTO dto.MedicationUpdateDTO
	if err := c.ShouldBindJSON(&medicationDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = mc.MedicationService.UpdateMedication(uint(id), &medicationDTO)
	if err != nil {
		log.Printf("Failed to update medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medication"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medication updated successfully", "medication": medicationDTO})
}

// DeleteMedication handles deleting a medication by its MedicationID
func (mc *MedicationController) DeleteMedication(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf("Invalid medication MedicationID: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication MedicationID"})
		return
	}

	err = mc.MedicationService.DeleteMedication(uint(id))
	if err != nil {
		log.Printf("Failed to delete medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medication"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medication deleted successfully"})
}
