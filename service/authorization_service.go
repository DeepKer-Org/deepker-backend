package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	log.Println("Registering a new user with username:", userDTO.Username)

	// Start a new transaction
	tx := s.repo.BeginTransaction()

	roles, err := s.roleRepo.GetRolesByNames(userDTO.Roles)
	if err != nil {
		log.Printf("Failed to fetch roles: %v", err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	// Update the DTO with the hashed password
	userDTO.Password = string(hashedPassword)

	user := dto.MapRegisterDTOToUser(userDTO, roles)
	err = s.repo.CreateInTransaction(user, tx)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	tx.Commit()

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	// Update the DTO with the hashed password
	userDTO.Password = string(hashedPassword)
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
	user, err := s.repo.GetUserByUsername(loginDTO.Username)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return "", err
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		log.Println("Incorrect password for user:", loginDTO.Username)
		return "", errors.New("incorrect password")
	}

	// Generate JWT token with user information
	token, err := GenerateToken(user.Username, dto.MapRolesToNames(user.Roles), map[string]interface{}{
		"user_id": user.UserID,
	})

	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", err
	}

	log.Println("User authenticated successfully:", loginDTO.Username)
	return token, nil
}
