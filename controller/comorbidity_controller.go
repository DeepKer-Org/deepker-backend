package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const invalidComorbidityIDMsg = "Invalid comorbidity ComorbidityID: %v"
const invalidComorbidityID = "Invalid comorbidity ComorbidityID"

type ComorbidityController struct {
	ComorbidityService service.ComorbidityService
}

func NewComorbidityController(comorbidityService service.ComorbidityService) *ComorbidityController {
	return &ComorbidityController{
		ComorbidityService: comorbidityService,
	}
}

// CreateComorbidity handles the creation of a new comorbidity
func (cc *ComorbidityController) CreateComorbidity(c *gin.Context) {
	var comorbidityDTO dto.ComorbidityCreateDTO
	if err := c.ShouldBindJSON(&comorbidityDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cc.ComorbidityService.CreateComorbidity(&comorbidityDTO)
	if err != nil {
		log.Printf("Failed to create comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comorbidity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comorbidity created successfully", "comorbidity": comorbidityDTO})
}

// GetComorbidityByID handles retrieving a comorbidity by their ComorbidityID
func (cc *ComorbidityController) GetComorbidityByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf(invalidComorbidityIDMsg, idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidComorbidityID})
		return
	}

	comorbidity, err := cc.ComorbidityService.GetComorbidityByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comorbidity"})
		return
	}

	if comorbidity == nil {
		log.Printf("Comorbidity not found with ComorbidityID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comorbidity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comorbidity": comorbidity})
}

// GetAllComorbidities handles retrieving all comorbidities
func (cc *ComorbidityController) GetAllComorbidities(c *gin.Context) {
	comorbidities, err := cc.ComorbidityService.GetAllComorbidities()
	if err != nil {
		log.Printf("Error retrieving comorbidities: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comorbidities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comorbidities": comorbidities})
}

// UpdateComorbidity handles updating an existing comorbidity
func (cc *ComorbidityController) UpdateComorbidity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf(invalidComorbidityIDMsg, idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidComorbidityID})
		return
	}

	var comorbidityDTO dto.ComorbidityUpdateDTO
	if err := c.ShouldBindJSON(&comorbidityDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = cc.ComorbidityService.UpdateComorbidity(uint(id), &comorbidityDTO)
	if err != nil {
		log.Printf("Failed to update comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comorbidity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comorbidity updated successfully", "comorbidity": comorbidityDTO})
}

// DeleteComorbidity handles deleting a comorbidity by their ComorbidityID
func (cc *ComorbidityController) DeleteComorbidity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Printf(invalidComorbidityIDMsg, idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidComorbidityID})
		return
	}

	err = cc.ComorbidityService.DeleteComorbidity(uint(id))
	if err != nil {
		log.Printf("Failed to delete comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comorbidity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comorbidity deleted successfully"})
}
