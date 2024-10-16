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

	alertResponse, err := ac.AlertService.CreateAlert(&alertDTO)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		switch err.Error() {
		case "record not found":
			c.JSON(http.StatusNotFound, alertResponse)
		default:
			c.JSON(http.StatusInternalServerError, alertResponse)
		}
		return
	}

	c.JSON(http.StatusCreated, alertResponse) //gin.H{"message": "Alert created successfully", "alert": alertDTO})
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

// GetAllAlerts handles retrieving all alerts with optional status filtering and pagination
func (ac *AlertController) GetAllAlerts(c *gin.Context) {
	status := c.Query("status")
	timezone := c.Query("timezone")

	var alerts []*dto.AlertDTO
	var totalCount int
	var err error

	if status != "" {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}

		// Fetch paginated alerts and total count by status
		alerts, totalCount, err = ac.AlertService.GetAllAlertsByStatus(status, page, limit)
		if err != nil {
			log.Printf("Error retrieving alerts by status: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve alerts by status"})
			return
		}

		// Include totalCount in the response for status-specific queries
		c.JSON(http.StatusOK, gin.H{
			"alerts":     alerts,
			"totalCount": totalCount,
		})
		return
	}

	// Handle case without status if needed (currently returns all alerts without pagination)
	if timezone != "" {
		// Fetch alerts from today
		alerts, err = ac.AlertService.GetAllAlertsByTimezone(timezone)
		if err != nil {
			log.Printf("Error retrieving today's timezone alerts: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve today's timezone alerts"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"alerts": alerts,
		})
	} else {
		alerts, err = ac.AlertService.GetAllAlerts()
		if err != nil {
			log.Printf("Error retrieving all alerts: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve alerts"})
			return
		}
		totalCount = len(alerts)

		c.JSON(http.StatusOK, gin.H{
			"alerts":     alerts,
			"totalCount": totalCount,
		})
	}

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
