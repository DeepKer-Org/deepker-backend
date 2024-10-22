package service

import (
	"biometric-data-backend/models/dto"
	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(userDTO *dto.UserRegisterDTO) error
	AuthenticateUser(loginDTO *dto.UserLoginDTO) (string, error)
	ResetPassword(resetDTO *dto.PasswordResetDTO) error
	RemoveUser(userID uuid.UUID) error
}

/*
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// RegisterUser registra un nuevo usuario
func (s *userService) RegisterUser(userDTO *dto.UserRegisterDTO) error {
	user := dto.MapRegisterDTOToUser(userDTO)
	err := s.repo.Create(user)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return err
	}
	log.Println("User registered successfully with UserID:", user.UserID)
	return nil
}

// AuthenticateUser autentica un usuario
func (s *userService) AuthenticateUser(loginDTO *dto.UserLoginDTO) (string, error) {
	user, err := s.repo.GetByUsername(loginDTO.Username)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return "", err
	}

	// Verificar contraseña (ejemplo básico)
	if user.Password != loginDTO.Password {
		log.Println("Incorrect password for user:", loginDTO.Username)
		return "", errors.New("incorrect password")
	}

	// Generar y devolver un token JWT (ejemplo)
	token, err := generateJWT(user)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", err
	}

	log.Println("User authenticated successfully:", loginDTO.Username)
	return token, nil
}

// ResetPassword restablece la contraseña del usuario
func (s *userService) ResetPassword(resetDTO *dto.PasswordResetDTO) error {
	user, err := s.repo.GetByUsername(resetDTO.Username)
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

	log.Println("Password reset successfully for user:", resetDTO.Username)
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
