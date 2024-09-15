package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

type PatientController struct {
	service service.PatientService
}

func NewPatientController(service service.PatientService) *PatientController {
	return &PatientController{service: service}
}

func (c *PatientController) CreatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient.ID = gocql.TimeUUID() // Generate a new UUID for the patient
	if err := c.service.CreatePatient(&patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully"})
}

func (c *PatientController) GetAllPatients(ctx *gin.Context) {
	patients, err := c.service.GetAllPatients()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patients)
}

func (c *PatientController) GetPatientByID(ctx *gin.Context) {
	id, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	patient, err := c.service.GetPatientByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

func (c *PatientController) UpdatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patientID, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	patient.ID = patientID
	if err := c.service.UpdatePatient(&patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

func (c *PatientController) DeletePatient(ctx *gin.Context) {
	id, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	if err := c.service.DeletePatient(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
