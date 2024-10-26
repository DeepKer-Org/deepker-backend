package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type PhoneRepository interface {
	BaseRepository[models.Phone]
}

type phoneRepository struct {
	BaseRepository[models.Phone]
	db *gorm.DB
}

func NewPhoneRepository(db *gorm.DB) PhoneRepository {
	baseRepo := NewBaseRepository[models.Phone](db)
	return &phoneRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}
