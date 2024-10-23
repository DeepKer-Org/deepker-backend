package service

import (
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
	//ResetPassword(resetDTO *dto.PasswordResetDTO) error
	//RemoveUser(userID uuid.UUID) error
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

// AuthenticateUser handles user authentication
func (s *userService) AuthenticateUser(loginDTO *dto.UserLoginDTO) (string, error) {
	//user, err := s.repo.GetUserByEmail(loginDTO.Email)
	user, err := s.repo.GetUserByEmail(loginDTO.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return "", err
	}

	if user.Password != loginDTO.Password {
		log.Println("Incorrect password for user:", loginDTO.Email)
		return "", errors.New("incorrect password")
	}

	token, err := GenerateToken(loginDTO.Email, dto.MapRolesToNames(user.Roles))

	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", err
	}

	log.Println("User authenticated successfully:", loginDTO.Email)
	return token, nil
}

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

/*
// ResetPassword restablece la contraseña del usuario
func (s *userService) ResetPassword(resetDTO *dto.PasswordResetDTO) error {
	user, err := s.repo.GetByEmail(resetDTO.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return err
	}

	// Actualizar la contraseña
	user.Password = resetDTO.NewPassword
	err = s.repo.Update(user)
	if err != nil {
		log.Printf("Failed to reset password: %v", err)
		return err
	}

	log.Println("Password reset successfully for user:", resetDTO.Email)
	return nil
}

// RemoveUser elimina un usuario por su ID
func (s *userService) RemoveUser(userID uuid.UUID) error {
	log.Println("Deleting user with UserID:", userID)
	err := s.repo.Delete(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found with UserID:", userID)
			return nil
		}
		log.Printf("Failed to delete user: %v", err)
		return err
	}
	log.Println("User deleted successfully with UserID:", userID)
	return nil
}
*/
