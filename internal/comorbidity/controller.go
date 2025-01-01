package comorbidity

import (
	"biometric-data-backend/internal/comorbidity/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Controller struct {
	Service Service
}

func NewController(comorbidityService Service) *Controller {
	return &Controller{
		Service: comorbidityService,
	}
}

// CreateComorbidity handles the creation of a new comorbidity
func (cc *Controller) CreateComorbidity(c *gin.Context) {
	var comorbidityDTO dto.CreateDTO
	if err := c.ShouldBindJSON(&comorbidityDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cc.Service.CreateComorbidity(&comorbidityDTO)
	if err != nil {
		log.Printf("Failed to create comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comorbidity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comorbidity created successfully", "comorbidity": comorbidityDTO})
}

// GetComorbidityByID handles retrieving a comorbidity by their ComorbidityID
func (cc *Controller) GetComorbidityByID(c *gin.Context) {
	id := c.Param("id")

	comorbidityID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comorbidity ID"})
		return
	}

	comorbidity, err := cc.Service.GetComorbidityByID(comorbidityID)
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
func (cc *Controller) GetAllComorbidities(c *gin.Context) {
	comorbidities, err := cc.Service.GetAllComorbidities()
	if err != nil {
		log.Printf("Error retrieving comorbidities: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comorbidities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comorbidities": comorbidities})
}

// UpdateComorbidity handles updating an existing comorbidity
func (cc *Controller) UpdateComorbidity(c *gin.Context) {
	id := c.Param("id")

	comorbidityID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comorbidity ID"})
		return
	}

	var comorbidityDTO dto.UpdateDTO
	if err := c.ShouldBindJSON(&comorbidityDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = cc.Service.UpdateComorbidity(comorbidityID, &comorbidityDTO)
	if err != nil {
		log.Printf("Failed to update comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comorbidity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comorbidity updated successfully", "comorbidity": comorbidityDTO})
}

// DeleteComorbidity handles deleting a comorbidity by their ComorbidityID
func (cc *Controller) DeleteComorbidity(c *gin.Context) {
	id := c.Param("id")

	comorbidityID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comorbidity ID"})
		return
	}

	err = cc.Service.DeleteComorbidity(comorbidityID)
	if err != nil {
		log.Printf("Failed to delete comorbidity: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comorbidity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comorbidity deleted successfully"})
}
