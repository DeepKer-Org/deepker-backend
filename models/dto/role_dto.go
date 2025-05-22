package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

type RoleCreateDTO struct {
	RoleName string `json:"role_name" binding:"required"`
}

type RoleUpdateDTO struct {
	RoleName string `json:"role_name" binding:"required"`
}

type RoleDTO struct {
	RoleID   uuid.UUID `json:"role_id"`
	RoleName string    `json:"role_name"`
}

// MapRoleToDTO maps a Role model to a RoleDTO
func MapRoleToDTO(role *models.Role) *RoleDTO {
	return &RoleDTO{
		RoleID:   role.RoleID,
		RoleName: string(role.RoleName),
	}
}

// MapRolesToDTOs maps a list of Role models to a list of RoleDTOs
func MapRolesToDTOs(roles []*models.Role) []*RoleDTO {
	roleDTOs := make([]*RoleDTO, 0)
	for _, role := range roles {
		roleDTOs = append(roleDTOs, MapRoleToDTO(role))
	}
	return roleDTOs
}

// MapRolesToNames maps a list of Role models to a list of role names
func MapRolesToNames(roles []*models.Role) []string {
	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, string(role.RoleName))
	}
	return roleNames
}
