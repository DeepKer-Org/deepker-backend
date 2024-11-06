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

type UserUpdateDTO struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
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

func MapUserToDTO(user *models.User) *UserDTO {
	roles := make([]string, 0)
	for _, role := range user.Roles {
		roles = append(roles, string(role.RoleName))
	}

	return &UserDTO{
		UserID:   user.UserID.String(),
		Username: user.Username,
		Roles:    roles,
	}
}

func MapUsersToDTOs(users []*models.User) []*UserDTO {
	userDTOs := make([]*UserDTO, 0)
	for _, user := range users {
		userDTOs = append(userDTOs, MapUserToDTO(user))
	}
	return userDTOs
}
