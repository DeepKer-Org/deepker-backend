package service

import (
	"biometric-data-backend/models"
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
	GetUserById(id uuid.UUID) (*dto.UserDTO, error)
	GetAllUsers(page int, limit int) ([]*dto.UserDTO, int, error)
	UpdateUser(id uuid.UUID, userDTO *dto.UserUpdateDTO) error
	DeleteUser(id uuid.UUID) error
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

func (s *userService) GetUserById(id uuid.UUID) (*dto.UserDTO, error) {
	log.Println("Fetching user with UserID:", id)

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		log.Printf("Failed to fetch user: %v", err)
		return nil, err
	}

	userDTO := dto.MapUserToDTO(user)
	log.Println("User fetched successfully with UserID:", user.UserID)
	return userDTO, nil
}

func (s *userService) GetAllUsers(page int, limit int) ([]*dto.UserDTO, int, error) {
	offset := (page - 1) * limit
	var users []*models.User
	var totalCount int64
	var err error

	users, err = s.repo.GetAllUsers(offset, limit)
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return nil, 0, err
	}

	totalCount, err = s.repo.CountAllUsers()
	if err != nil {
		log.Printf("Failed to count users: %v", err)
		return nil, 0, err
	}

	userDTOs := dto.MapUsersToDTOs(users)
	log.Println("Users fetched successfully, total count:", len(users))

	return userDTOs, int(totalCount), nil
}

func (s *userService) UpdateUser(id uuid.UUID, userDTO *dto.UserUpdateDTO) error {
	log.Println("Updating user with UserID:", id)

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return err
	}
	if user == nil {
		log.Printf("User not found with UserID: %v", id)
		return gorm.ErrRecordNotFound
	}

	// Update the user entity with the new values
	user.Username = userDTO.Username

	// If new password, hash it and update the user entity
	if userDTO.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password: %v", err)
			return err
		}
		user.Password = string(hashedPassword)
	}

	err = s.repo.Update(user, "user_id", id)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}

	log.Println("User updated successfully with UserID:", user.UserID)
	return nil
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	log.Println("Deleting user with UserID:", id)

	err := s.repo.DeleteUserAndUserRoles(id)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	log.Println("User deleted successfully with UserID:", id)
	return nil
}
