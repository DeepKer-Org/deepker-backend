package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type AlertController struct {
	AlertService service.AlertService
}

func NewAlertController(alertService service.AlertService) *AlertController {
	return &AlertController{
		AlertService: alertService,
	}
}

// CreateAlert handles the creation of a new alert
func (ac *AlertController) CreateAlert(c *gin.Context) {
	var alertDTO dto.AlertCreateDTO
	if err := c.ShouldBindJSON(&alertDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := ac.AlertService.CreateAlert(&alertDTO)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create alert"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alert created successfully", "alert": alertDTO})
}

// GetAlertByID handles retrieving an alert by its AlertID
func (ac *AlertController) GetAlertByID(c *gin.Context) {
	id := c.Param("id")

	alertID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	alert, err := ac.AlertService.GetAlertByID(alertID)
	if err != nil {
		log.Printf("Error retrieving alert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve alert"})
		return
	}

	if alert == nil {
		log.Printf("Alert not found with AlertID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alert": alert})
}

// GetAllAlerts handles retrieving all alerts
func (ac *AlertController) GetAllAlerts(c *gin.Context) {
	status := c.Query("status")
	alerts, err := func() ([]*dto.AlertDTO, error) {
		if status == "" {
			return ac.AlertService.GetAllAlerts()
		}
		return ac.AlertService.GetAllAlertsByStatus(status)
	}()
	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve alerts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alerts": alerts})
}

// UpdateAlert handles updating an existing alert
func (ac *AlertController) UpdateAlert(c *gin.Context) {
	id := c.Param("id")

	alertID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	var alertDTO dto.AlertUpdateDTO
	if err := c.ShouldBindJSON(&alertDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = ac.AlertService.UpdateAlert(alertID, &alertDTO)
	if err != nil {
		log.Printf("Failed to update alert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update alert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert updated successfully", "alert": alertDTO})
}

// DeleteAlert handles deleting an alert by its AlertID
func (ac *AlertController) DeleteAlert(c *gin.Context) {
	id := c.Param("id")

	alertID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	err = ac.AlertService.DeleteAlert(alertID)
	if err != nil {
		log.Printf("Failed to delete alert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete alert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}
