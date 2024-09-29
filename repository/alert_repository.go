package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type AlertRepository interface {
	BaseRepository[models.Alert]
}

type alertRepository struct {
	BaseRepository[models.Alert]
}

func NewAlertRepository(db *gorm.DB) AlertRepository {
	baseRepo := NewBaseRepository[models.Alert](db)
	return &alertRepository{baseRepo}
}
