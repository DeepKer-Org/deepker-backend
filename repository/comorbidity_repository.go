package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type ComorbidityRepository interface {
	BaseRepository[models.Comorbidity]
}

type comorbidityRepository struct {
	BaseRepository[models.Comorbidity]
}

func NewComorbidityRepository(db *gorm.DB) ComorbidityRepository {
	baseRepo := NewBaseRepository[models.Comorbidity](db)
	return &comorbidityRepository{baseRepo}
}
