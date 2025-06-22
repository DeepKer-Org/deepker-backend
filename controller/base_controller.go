package controller

import (
	"biometric-data-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CRUDServiceInterface defines the service methods needed for CRUD operations
type CRUDServiceInterface[CreateDTO any, UpdateDTO any] interface {
	Create(dto *CreateDTO) error
	GetByID(id uuid.UUID) (interface{}, error)
	GetAll() (interface{}, error)
	Update(id uuid.UUID, dto *UpdateDTO) error
	Delete(id uuid.UUID) error
}

// BaseController provides generic CRUD operations
type BaseController[CreateDTO any, UpdateDTO any] struct {
	service     CRUDServiceInterface[CreateDTO, UpdateDTO]
	resourceName string
}

// NewBaseController creates a new base controller
func NewBaseController[CreateDTO any, UpdateDTO any](
	service CRUDServiceInterface[CreateDTO, UpdateDTO],
	resourceName string,
) *BaseController[CreateDTO, UpdateDTO] {
	return &BaseController[CreateDTO, UpdateDTO]{
		service:      service,
		resourceName: resourceName,
	}
}

// Create handles POST requests to create a new entity
func (bc *BaseController[CreateDTO, UpdateDTO]) Create(c *gin.Context) {
	var dto CreateDTO
	if !utils.BindJSONAndValidate(c, &dto) {
		return
	}

	err := bc.service.Create(&dto)
	if err != nil {
		utils.RespondInternalError(c, "create "+bc.resourceName, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, bc.resourceName+" created successfully", dto)
}

// GetByID handles GET requests to retrieve an entity by ID
func (bc *BaseController[CreateDTO, UpdateDTO]) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := utils.ValidateUUID(c, idStr)
	if !ok {
		return
	}

	entity, err := bc.service.GetByID(id)
	if err != nil {
		utils.RespondInternalError(c, "retrieve "+bc.resourceName, err)
		return
	}

	if entity == nil {
		utils.RespondNotFound(c, bc.resourceName)
		return
	}

	utils.RespondWithData(c, http.StatusOK, gin.H{bc.resourceName: entity})
}

// GetAll handles GET requests to retrieve all entities
func (bc *BaseController[CreateDTO, UpdateDTO]) GetAll(c *gin.Context) {
	entities, err := bc.service.GetAll()
	if err != nil {
		utils.RespondInternalError(c, "retrieve "+bc.resourceName+"s", err)
		return
	}

	utils.RespondWithData(c, http.StatusOK, gin.H{bc.resourceName + "s": entities})
}

// Update handles PATCH/PUT requests to update an entity
func (bc *BaseController[CreateDTO, UpdateDTO]) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := utils.ValidateUUID(c, idStr)
	if !ok {
		return
	}

	var dto UpdateDTO
	if !utils.BindJSONAndValidate(c, &dto) {
		return
	}

	err := bc.service.Update(id, &dto)
	if err != nil {
		utils.RespondInternalError(c, "update "+bc.resourceName, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, bc.resourceName+" updated successfully", dto)
}

// Delete handles DELETE requests to remove an entity
func (bc *BaseController[CreateDTO, UpdateDTO]) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := utils.ValidateUUID(c, idStr)
	if !ok {
		return
	}

	err := bc.service.Delete(id)
	if err != nil {
		utils.RespondInternalError(c, "delete "+bc.resourceName, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, bc.resourceName+" deleted successfully")
}

// GetPaginatedAll handles GET requests with pagination support
func (bc *BaseController[CreateDTO, UpdateDTO]) GetPaginatedAll(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)
	sort := utils.GetSortParams(c, "created_at")
	filters := utils.GetFilterParams(c)

	// Note: This would require extending the service interface to support pagination
	// For now, fallback to GetAll
	entities, err := bc.service.GetAll()
	if err != nil {
		utils.RespondInternalError(c, "retrieve "+bc.resourceName+"s", err)
		return
	}

	response := gin.H{
		bc.resourceName + "s": entities,
		"pagination": gin.H{
			"page":  pagination.Page,
			"limit": pagination.Limit,
		},
		"sort":    sort,
		"filters": filters,
	}

	utils.RespondWithData(c, http.StatusOK, response)
}

// Enhanced controller with additional common operations
type EnhancedController[CreateDTO any, UpdateDTO any] struct {
	*BaseController[CreateDTO, UpdateDTO]
}

// NewEnhancedController creates a new enhanced controller with additional features
func NewEnhancedController[CreateDTO any, UpdateDTO any](
	service CRUDServiceInterface[CreateDTO, UpdateDTO],
	resourceName string,
) *EnhancedController[CreateDTO, UpdateDTO] {
	return &EnhancedController[CreateDTO, UpdateDTO]{
		BaseController: NewBaseController(service, resourceName),
	}
}

// CreateMultiple handles POST requests to create multiple entities
func (ec *EnhancedController[CreateDTO, UpdateDTO]) CreateMultiple(c *gin.Context) {
	var dtos []CreateDTO
	if !utils.BindJSONAndValidate(c, &dtos) {
		return
	}

	if len(dtos) == 0 {
		utils.RespondBadRequest(c, "At least one "+ec.resourceName+" is required")
		return
	}

	// Create each entity (could be optimized with batch operations)
	createdCount := 0
	for _, dto := range dtos {
		err := ec.service.Create(&dto)
		if err != nil {
			utils.RespondInternalError(c, "create "+ec.resourceName+" batch", err)
			return
		}
		createdCount++
	}

	utils.RespondWithSuccess(c, http.StatusCreated, 
		"Successfully created "+string(rune(createdCount))+" "+ec.resourceName+"s", 
		gin.H{"created_count": createdCount})
}

// HealthCheck provides a health check endpoint for the controller
func (ec *EnhancedController[CreateDTO, UpdateDTO]) HealthCheck(c *gin.Context) {
	utils.RespondWithSuccess(c, http.StatusOK, ec.resourceName+" controller is healthy")
}