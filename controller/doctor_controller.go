package controller

import (
	"biometric-data-backend/models/dto"
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
	var doctorDTO dto.DoctorCreateDTO
	if err := c.ShouldBindJSON(&doctorDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := dc.DoctorService.CreateDoctor(&doctorDTO)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Doctor created successfully", "doctor": doctorDTO})
}

// GetDoctorByID handles retrieving a doctor by their DoctorID
func (dc *DoctorController) GetDoctorByID(c *gin.Context) {
	id := c.Param("id")

	doctorID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	doctor, err := dc.DoctorService.GetDoctorByID(doctorID)
	if err != nil {
		log.Printf("Error retrieving doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor"})
		return
	}

	if doctor == nil {
		log.Printf("Doctor not found with DoctorID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// GetDoctorByUserID handles retrieving a doctor by their UserID
func (dc *DoctorController) GetDoctorByUserID(c *gin.Context) {
	id := c.Param("userID")

	userID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	doctor, err := dc.DoctorService.GetDoctorByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor"})
		return
	}

	if doctor == nil {
		log.Printf("Doctor not found with UserID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// GetDoctorsByAlertID handles retrieving all doctors associated with an alert
func (dc *DoctorController) GetDoctorsByAlertID(c *gin.Context) {
	alertID := c.Param("alertID")

	// Parse the string to a UUID
	alertUUID, err := uuid.Parse(alertID)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	doctors, err := dc.DoctorService.GetDoctorsByAlertID(alertUUID)
	if err != nil {
		log.Printf("Error retrieving doctors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
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
	id := c.Param("id")

	doctorID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var doctorDTO dto.DoctorUpdateDTO
	if err := c.ShouldBindJSON(&doctorDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = dc.DoctorService.UpdateDoctor(doctorID, &doctorDTO)
	if err != nil {
		log.Printf("Failed to update doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully", "doctor": doctorDTO})
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

// ChangePassword handles changing a doctor's password
func (dc *DoctorController) ChangePassword(c *gin.Context) {
	var changePasswordDTO dto.ChangePasswordDTO
	if err := c.ShouldBindJSON(&changePasswordDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := dc.DoctorService.ChangePassword(&changePasswordDTO)
	if err != nil {
		log.Printf("Failed to change password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
