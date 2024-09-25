package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlertController struct {
	alertService service.AlertService
}

// NewAlertController crea una nueva instancia del controlador
func NewAlertController(alertService service.AlertService) *AlertController {
	return &AlertController{
		alertService: alertService,
	}
}

// GetAllAlerts maneja la solicitud para obtener todas las alertas o filtrar por query params
func (c *AlertController) GetAllAlerts(ctx *gin.Context) {
	// Obtener los query params
	status := ctx.Query("status")

	var alerts []models.Alert
	var err error

	// Verificar si hay query params y filtrar si es necesario
	if status != "" {
		alerts, err = c.alertService.GetAlertByStatus(status)
	} else {
		alerts, err = c.alertService.GetAllAlerts()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alerts)
}

// GetAlertByID maneja la solicitud para obtener una alerta por ID
func (c *AlertController) GetAlertByID(ctx *gin.Context) {
	id := ctx.Param("id")
	alert, err := c.alertService.GetAlertByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}
	ctx.JSON(http.StatusOK, alert)
}

// CreateAlert maneja la solicitud para crear una nueva alerta
func (c *AlertController) CreateAlert(ctx *gin.Context) {
	var alert models.Alert
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAlert, err := c.alertService.CreateAlert(alert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdAlert)
}

// UpdateAlert maneja la solicitud para actualizar una alerta
func (c *AlertController) UpdateAlert(ctx *gin.Context) {
	id := ctx.Param("id")
	var alert models.Alert

	// Bind the incoming JSON to the alert struct
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer el ID en la alerta para actualizar el registro correcto
	alert.AlertID = id

	updatedAlert, err := c.alertService.UpdateAlert(alert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedAlert)
}

// DeleteAlert maneja la solicitud para eliminar (lógicamente) una alerta
func (c *AlertController) DeleteAlert(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.alertService.DeleteAlert(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}
