package service

import (
	"biometric-data-backend/enums"
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type RoleService interface {
	CreateRole(roleDTO *dto.RoleCreateDTO) error
	GetRoleByID(id uuid.UUID) (*dto.RoleDTO, error)
	GetAllRoles() ([]*dto.RoleDTO, error)
	UpdateRole(id uuid.UUID, roleDTO *dto.RoleUpdateDTO) error
	DeleteRole(id uuid.UUID) error
	GetRolesByNames(roleNames []string) ([]*dto.RoleDTO, error)
}

type roleService struct {
	repo repository.RoleRepository
}

// NewRoleService creates a new instance of RoleService
func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

// CreateRole creates a new role in the system
func (s *roleService) CreateRole(roleDTO *dto.RoleCreateDTO) error {
	role := &models.Role{
		RoleName: enums.RoleEnum(roleDTO.RoleName),
	}

	err := s.repo.Create(role)
	if err != nil {
		log.Printf("Failed to create role: %v", err)
		return err
	}
	log.Println("Role created successfully with RoleID:", role.RoleID)
	return nil
}

// GetRoleByID gets a role by its ID
func (s *roleService) GetRoleByID(id uuid.UUID) (*dto.RoleDTO, error) {
	log.Println("Fetching role with RoleID:", id)
	role, err := s.repo.GetByID(id, "role_id")
	if err != nil {
		log.Printf("Error retrieving role: %v", err)
		return nil, err
	}
	if role == nil {
		log.Println("No role found with RoleID:", id)
		return nil, nil
	}

	roleDTO := dto.MapRoleToDTO(role)
	log.Println("Role fetched successfully with RoleID:", id)
	return roleDTO, nil
}

// GetAllRoles retrieves all roles
func (s *roleService) GetAllRoles() ([]*dto.RoleDTO, error) {
	log.Println("Fetching all roles")
	roles, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error retrieving roles: %v", err)
		return nil, err
	}

	roleDTOs := dto.MapRolesToDTOs(roles)
	log.Println("Roles fetched successfully, total count:", len(roleDTOs))
	return roleDTOs, nil
}

// UpdateRole updates an existing role
func (s *roleService) UpdateRole(id uuid.UUID, roleDTO *dto.RoleUpdateDTO) error {
	log.Println("Updating role with RoleID:", id)

	role, err := s.repo.GetByID(id, "role_id")
	if err != nil {
		log.Printf("Error retrieving role: %v", err)
		return err
	}
	if role == nil {
		log.Printf("Role not found with RoleID: %v", id)
		return gorm.ErrRecordNotFound
	}

	role.RoleName = enums.RoleEnum(roleDTO.RoleName)

	err = s.repo.Update(role, "role_id", id)
	if err != nil {
		log.Printf("Failed to update role: %v", err)
		return err
	}
	log.Println("Role updated successfully with RoleID:", role.RoleID)
	return nil
}

// DeleteRole deletes a role by its ID
func (s *roleService) DeleteRole(id uuid.UUID) error {
	log.Println("Deleting role with RoleID:", id)
	err := s.repo.Delete(id, "role_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Role not found with RoleID:", id)
			return nil
		}
		log.Printf("Failed to delete role: %v", err)
		return err
	}
	log.Println("Role deleted successfully with RoleID:", id)
	return nil
}

// GetRolesByNames retrieves roles by their role names
func (s *roleService) GetRolesByNames(roleNames []string) ([]*dto.RoleDTO, error) {
	log.Println("Fetching roles by names:", roleNames)
	roles, err := s.repo.GetRolesByNames(roleNames)
	if err != nil {
		log.Printf("Error retrieving roles by names: %v", err)
		return nil, err
	}

	roleDTOs := dto.MapRolesToDTOs(roles)
	log.Println("Roles fetched successfully by names")
	return roleDTOs, nil
}
