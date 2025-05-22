package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

// CreateRole handles the creation of a new role
func (rc *RoleController) CreateRole(c *gin.Context) {
	var roleDTO dto.RoleCreateDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := rc.RoleService.CreateRole(&roleDTO)
	if err != nil {
		log.Printf("Failed to create role: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role created successfully", "role": roleDTO})
}

// GetRoleByID handles retrieving a role by its RoleID
func (rc *RoleController) GetRoleByID(c *gin.Context) {
	id := c.Param("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := rc.RoleService.GetRoleByID(roleID)
	if err != nil {
		log.Printf("Error retrieving role: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve role"})
		return
	}

	if role == nil {
		log.Printf("Role not found with RoleID: %v", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// GetAllRoles handles retrieving all roles
func (rc *RoleController) GetAllRoles(c *gin.Context) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		log.Printf("Error retrieving roles: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

// UpdateRole handles updating an existing role
func (rc *RoleController) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var roleDTO dto.RoleUpdateDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = rc.RoleService.UpdateRole(roleID, &roleDTO)
	if err != nil {
		log.Printf("Failed to update role: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully", "role": roleDTO})
}

// DeleteRole handles deleting a role by its RoleID
func (rc *RoleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	err = rc.RoleService.DeleteRole(roleID)
	if err != nil {
		log.Printf("Failed to delete role: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// GetRolesByNames handles retrieving roles by their names
func (rc *RoleController) GetRolesByNames(c *gin.Context) {
	var roleNames []string
	if err := c.ShouldBindJSON(&roleNames); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	roles, err := rc.RoleService.GetRolesByNames(roleNames)
	if err != nil {
		log.Printf("Error retrieving roles by names: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}
