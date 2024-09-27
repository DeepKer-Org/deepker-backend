package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
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
	if !bindJSON(c, &medicationDTO) {
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
	medication, err := getByID(c, "id", mc.MedicationService.GetMedicationByID, "Medication not found with MedicationID: %v")
	if err != nil || medication == nil {
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
	id := c.Param("id")

	medicationID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication ID"})
		return
	}

	var medicationDTO dto.MedicationUpdateDTO
	if !bindJSON(c, &medicationDTO) {
		return
	}

	err = mc.MedicationService.UpdateMedication(medicationID, &medicationDTO)
	if err != nil {
		log.Printf("Failed to update medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medication"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medication updated successfully", "medication": medicationDTO})
}

// DeleteMedication handles deleting a medication by its MedicationID
func (mc *MedicationController) DeleteMedication(c *gin.Context) {
	id := c.Param("id")

	medicationID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication ID"})
		return
	}

	err = mc.MedicationService.DeleteMedication(medicationID)
	if err != nil {
		log.Printf("Failed to delete medication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medication"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medication deleted successfully"})
}
