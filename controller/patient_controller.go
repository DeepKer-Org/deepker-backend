package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

type PatientController struct {
	PatientService service.PatientService
}

func NewPatientController(patientService service.PatientService) *PatientController {
	return &PatientController{
		PatientService: patientService,
	}
}

// CreatePatient handles the creation of a new patient
func (pc *PatientController) CreatePatient(c *gin.Context) {
	var patientDTO dto.PatientCreateDTO
	if err := c.ShouldBindJSON(&patientDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := pc.PatientService.CreatePatient(&patientDTO)
	if err != nil {
		log.Printf("Failed to create patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully", "patient": patientDTO})
}

// GetPatientByID handles retrieving a patient by their PatientID
func (pc *PatientController) GetPatientByID(c *gin.Context) {
	id := c.Param("id")

	patientID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	patient, err := pc.PatientService.GetPatientByID(patientID)
	if err != nil {
		log.Printf("Error retrieving patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient"})
		return
	}

	if patient == nil {
		log.Printf("Patient not found with PatientID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

// GetPatientByDNI handles retrieving a patient by their DNI
func (pc *PatientController) GetPatientByDNI(c *gin.Context) {
	dni := c.Param("dni")

	patient, err := pc.PatientService.GetPatientByDNI(dni)
	if err != nil {
		log.Printf("Error retrieving patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient"})
		return
	}

	if patient == nil {
		log.Printf("Patient not found with DNI: %v", dni)
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

// GetAllPatients handles retrieving all patients
func (pc *PatientController) GetAllPatients(c *gin.Context) {
	var patients []*dto.PatientDTO
	var totalCount int
	var err error

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Advanced Filters
	name := c.Query("name")
	dni := c.Query("dni")
	age, _ := strconv.Atoi(c.DefaultQuery("age", "0"))
	doctorID := c.Query("doctor_id")
	location := c.Query("location")
	deviceID := c.Query("device_id")
	comorbidityName := c.Query("comorbidity")
	entryDate := c.Query("entry_date")
	dischargeDate := c.Query("discharge_date")

	filters := dto.PatientFilter{
		Name:            name,
		DNI:             dni,
		Age:             age,
		DoctorID:        doctorID,
		Location:        location,
		DeviceID:        deviceID,
		ComorbidityName: comorbidityName,
		EntryDate:       entryDate,
		DischargeDate:   dischargeDate,
	}

	patients, totalCount, err = pc.PatientService.GetAllPatients(page, limit, filters)
	if err != nil {
		log.Printf("Error retrieving patients: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"patients":   patients,
		"totalCount": totalCount,
	})
	return
}

// UpdatePatient handles updating an existing patient
func (pc *PatientController) UpdatePatient(c *gin.Context) {
	id := c.Param("id")

	patientID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patientDTO dto.PatientUpdateDTO
	if err := c.ShouldBindJSON(&patientDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = pc.PatientService.UpdatePatient(patientID, &patientDTO)
	if err != nil {
		log.Printf("Failed to update patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully", "patient": patientDTO})
}

// DeletePatient handles deleting a patient by their PatientID
func (pc *PatientController) DeletePatient(c *gin.Context) {
	id := c.Param("id")

	patientID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	err = pc.PatientService.DeletePatient(patientID)
	if err != nil {
		log.Printf("Failed to delete patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
