package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
}

type userRepository struct {
	BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	baseRepo := NewBaseRepository[models.User](db)
	return &userRepository{baseRepo}
}
