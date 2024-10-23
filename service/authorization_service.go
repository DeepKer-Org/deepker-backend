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

type AuthorizationService interface {
	RegisterUser(userDTO *dto.UserRegisterDTO) (*uuid.UUID, error)
	RegisterUserInTransaction(userDTO *dto.UserRegisterDTO, tx *gorm.DB) (*uuid.UUID, error)
	AuthenticateUser(loginDTO *dto.UserLoginDTO) (string, error)
}

type userService struct {
	repo     repository.AuthorizationRepository
	roleRepo repository.RoleRepository
}

func NewUserService(repo repository.AuthorizationRepository, roleRepo repository.RoleRepository) AuthorizationService {
	return &userService{
		repo:     repo,
		roleRepo: roleRepo,
	}
}

// RegisterUser handles user registration
func (s *userService) RegisterUser(userDTO *dto.UserRegisterDTO) (*uuid.UUID, error) {
	roles, err := s.roleRepo.GetRolesByNames(userDTO.Roles)
	if err != nil {
		log.Printf("Failed to fetch roles: %v", err)
		return nil, err
	}
	user := dto.MapRegisterDTOToUser(userDTO, roles)
	err = s.repo.Create(user)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, err
	}
	log.Println("User registered successfully with UserID:", user.UserID)
	return &user.UserID, nil
}

// RegisterUserInTransaction handles user registration within a transaction
func (s *userService) RegisterUserInTransaction(userDTO *dto.UserRegisterDTO, tx *gorm.DB) (*uuid.UUID, error) {
	// Fetch roles within the transaction
	roles, err := s.roleRepo.GetRolesByNames(userDTO.Roles)
	if err != nil {
		log.Printf("Failed to fetch roles: %v", err)
		return nil, err
	}

	// Map the DTO to the User entity with the roles
	user := dto.MapRegisterDTOToUser(userDTO, roles)

	// Create the user within the transaction
	err = s.repo.CreateInTransaction(user, tx)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, err
	}

	log.Println("User registered successfully with UserID:", user.UserID)
	return &user.UserID, nil
}

// AuthenticateUser handles user authentication
func (s *userService) AuthenticateUser(loginDTO *dto.UserLoginDTO) (string, error) {
	user, err := s.repo.GetUserByEmail(loginDTO.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return "", err
	}

	if user.Password != loginDTO.Password {
		log.Println("Incorrect password for user:", loginDTO.Email)
		return "", errors.New("incorrect password")
	}

	token, err := GenerateToken(user.Email, dto.MapRolesToNames(user.Roles), map[string]interface{}{
		"user_id": user.UserID,
	})

	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", err
	}

	log.Println("User authenticated successfully:", loginDTO.Email)
	return token, nil
}

// hasRole helper function to check if a user has a specific role
func hasRole(roles []*models.Role, targetRole enums.RoleEnum) bool {
	for _, role := range roles {
		if role.RoleName == targetRole {
			return true
		}
	}
	return false
}
