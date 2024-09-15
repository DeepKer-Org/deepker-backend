package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

type AlertController struct {
	service service.AlertService
}

func NewAlertController(service service.AlertService) *AlertController {
	return &AlertController{service: service}
}

// Create a new alert
func (c *AlertController) CreateAlert(ctx *gin.Context) {
	var alert models.Alert
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	alert.AlertID = gocql.TimeUUID() // Generate a new UUID for the alert
	if err := c.service.CreateAlert(&alert); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Alert created successfully"})
}

// Get all alerts
func (c *AlertController) GetAllAlerts(ctx *gin.Context) {
	alerts, err := c.service.GetAllAlerts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alerts)
}

// Get an alert by ID
func (c *AlertController) GetAlertByID(ctx *gin.Context) {
	id, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	alert, err := c.service.GetAlertByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alert)
}

// Update an existing alert
func (c *AlertController) UpdateAlert(ctx *gin.Context) {
	var alert models.Alert
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	alertID, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	alert.AlertID = alertID
	if err := c.service.UpdateAlert(&alert); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Alert updated successfully"})
}

// Delete an alert by ID
func (c *AlertController) DeleteAlert(ctx *gin.Context) {
	id, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	if err := c.service.DeleteAlert(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}

// Get an alert along with patient details
func (c *AlertController) GetAlertWithPatient(ctx *gin.Context) {
	id, err := gocql.ParseUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	alertWithPatient, err := c.service.GetAlertWithPatient(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alertWithPatient)
}
