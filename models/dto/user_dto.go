package dto

import (
	"biometric-data-backend/models"
)

// UserRegisterDTO is used for registering a new user
type UserRegisterDTO struct {
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required,min=12"`
	Roles    []string `json:"roles" binding:"required"`
}

// UserLoginDTO is used for authenticating a user
type UserLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PasswordResetDTO is used for resetting a user's password
type PasswordResetDTO struct {
	Username    string `json:"username" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=12"`
}

// UserDTO is used for retrieving a user
type UserDTO struct {
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// MapRegisterDTOToUser maps a UserRegisterDTO to a User model
func MapRegisterDTOToUser(dto *UserRegisterDTO, roles []*models.Role) *models.User {
	return &models.User{
		Username: dto.Username,
		Password: dto.Password,
		Roles:    roles,
	}
}
