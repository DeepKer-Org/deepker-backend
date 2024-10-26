package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PhoneController struct {
	PhoneService service.PhoneService
}

func NewPhoneController(phoneService service.PhoneService) *PhoneController {
	return &PhoneController{
		PhoneService: phoneService,
	}
}

func (pc *PhoneController) CreatePhone(c *gin.Context) {
	var phoneDTO dto.PhoneCreateDTO
	if err := c.ShouldBindJSON(&phoneDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := pc.PhoneService.CreatePhone(&phoneDTO)
	if err != nil {
		log.Printf("Failed to create patient: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create phone"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Phone created successfully", "phone": phoneDTO})
}
