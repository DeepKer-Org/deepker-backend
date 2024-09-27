package controller

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type DoctorController struct {
	DoctorService service.DoctorService
}

func NewDoctorController(doctorService service.DoctorService) *DoctorController {
	return &DoctorController{
		DoctorService: doctorService,
	}
}

// CreateDoctor handles the creation of a new doctor
func (dc *DoctorController) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := dc.DoctorService.CreateDoctor(&doctor)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Doctor created successfully", "doctor": doctor})
}

// GetDoctorByID handles retrieving a doctor by their DoctorID
func (dc *DoctorController) GetDoctorByID(c *gin.Context) {
	doctor, err := getByID(c, "id", dc.DoctorService.GetDoctorByID, "Doctor not found with DoctorID: %v")
	if err != nil || doctor == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// GetShortDoctorByID handles retrieving a doctor by their DoctorID and returning a DTO
func (dc *DoctorController) GetShortDoctorByID(c *gin.Context) {
	doctorDTO, err := getByID(c, "id", dc.DoctorService.GetShortDoctorByID, "Doctor not found with DoctorID: %v")
	if err != nil || doctorDTO == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctor": doctorDTO})
}

// GetAllDoctors handles retrieving all doctors
func (dc *DoctorController) GetAllDoctors(c *gin.Context) {
	doctors, err := dc.DoctorService.GetAllDoctors()
	if err != nil {
		log.Printf("Error retrieving doctors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
}

// UpdateDoctor handles updating an existing doctor
func (dc *DoctorController) UpdateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := dc.DoctorService.UpdateDoctor(&doctor)
	if err != nil {
		log.Printf("Failed to update doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully", "doctor": doctor})
}

// DeleteDoctor handles deleting a doctor by their DoctorID
func (dc *DoctorController) DeleteDoctor(c *gin.Context) {
	id := c.Param("id")

	// Parse the string to a UUID
	doctorID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	err = dc.DoctorService.DeleteDoctor(doctorID)
	if err != nil {
		log.Printf("Failed to delete doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}
