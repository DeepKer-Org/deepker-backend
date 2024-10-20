package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type MonitoringDeviceController struct {
	MonitoringDeviceService service.MonitoringDeviceService
}

func NewMonitoringDeviceController(deviceService service.MonitoringDeviceService) *MonitoringDeviceController {
	return &MonitoringDeviceController{
		MonitoringDeviceService: deviceService,
	}
}

// CreateMonitoringDevice handles the creation of a new monitoring device
func (mdc *MonitoringDeviceController) CreateMonitoringDevice(c *gin.Context) {
	var deviceDTO dto.MonitoringDeviceCreateDTO
	if err := c.ShouldBindJSON(&deviceDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := mdc.MonitoringDeviceService.CreateMonitoringDevice(&deviceDTO)
	if err != nil {
		log.Printf("Failed to create monitoring device: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create monitoring device"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Monitoring device created successfully", "device": deviceDTO})
}

// GetMonitoringDeviceByID handles retrieving a monitoring device by its DeviceID
func (mdc *MonitoringDeviceController) GetMonitoringDeviceByID(c *gin.Context) {
	id := c.Param("id")

	device, err := mdc.MonitoringDeviceService.GetMonitoringDeviceByID(id)
	if err != nil {
		log.Printf("Error retrieving monitoring device: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve monitoring device"})
		return
	}

	if device == nil {
		log.Printf("Monitoring device not found with DeviceID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitoring device not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"device": device})
}

// GetAllMonitoringDevices handles retrieving all monitoring devices
func (mdc *MonitoringDeviceController) GetAllMonitoringDevices(c *gin.Context) {
	var totalCount int
	var devices []*dto.MonitoringDeviceDTO
	var err error

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	dni := c.Query("dni")
	status := c.Query("status") // Add status as a query parameter

	if status != "" {
		devices, totalCount, err = mdc.MonitoringDeviceService.GetAllMonitoringDevicesByStatus(status)
	} else {
		// Build the filters object
		filters := dto.MonitoringDeviceFilter{
			DNI: dni,
		}

		devices, totalCount, err = mdc.MonitoringDeviceService.GetAllMonitoringDevices(page, limit, filters)
	}

	if err != nil {
		log.Printf("Error retrieving monitoring devices: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve monitoring devices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"devices":    devices,
		"totalCount": totalCount,
	})
	return
}

// UpdateMonitoringDevice handles updating an existing monitoring device
func (mdc *MonitoringDeviceController) UpdateMonitoringDevice(c *gin.Context) {
	id := c.Param("id")

	log.Println("Updating monitoring device with DeviceID:", id)

	var deviceDTO dto.MonitoringDeviceUpdateDTO
	if err := c.ShouldBindJSON(&deviceDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := mdc.MonitoringDeviceService.UpdateMonitoringDevice(id, &deviceDTO)
	if err != nil {
		log.Printf("Failed to update monitoring device: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update monitoring device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Monitoring device updated successfully", "device": deviceDTO})
}

// DeleteMonitoringDevice handles deleting a monitoring device by its DeviceID
func (mdc *MonitoringDeviceController) DeleteMonitoringDevice(c *gin.Context) {
	id := c.Param("id")

	err := mdc.MonitoringDeviceService.DeleteMonitoringDevice(id)
	if err != nil {
		log.Printf("Failed to delete monitoring device: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete monitoring device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Monitoring device deleted successfully"})
}
