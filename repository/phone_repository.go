package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type PhoneRepository interface {
	BaseRepository[models.Phone]
	ExistsByExponentPushToken(token string) (bool, error)
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

func (r *phoneRepository) ExistsByExponentPushToken(token string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Phone{}).Where("exponent_push_token = ?", token).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
