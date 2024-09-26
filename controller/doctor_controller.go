package controller

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf("Invalid doctor DoctorID: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor DoctorID"})
		return
	}

	doctor, err := dc.DoctorService.GetDoctorByID(uint(id))
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

// GetShortDoctorByID handles retrieving a doctor by their DoctorID and returning a DTO
func (dc *DoctorController) GetShortDoctorByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	doctorDTO, err := dc.DoctorService.GetShortDoctorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor"})
		return
	}

	if doctorDTO == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
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
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf("Invalid doctor DoctorID: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor DoctorID"})
		return
	}

	err = dc.DoctorService.DeleteDoctor(uint(id))
	if err != nil {
		log.Printf("Failed to delete doctor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}
