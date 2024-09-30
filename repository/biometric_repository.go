package repository

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BiometricDataRepository interface {
	BaseRepository[models.BiometricData]
	GetBiometricRecordsByAlertID(id uuid.UUID) ([]*models.BiometricData, error)
}

type biometricRepository struct {
	BaseRepository[models.BiometricData]
	db *gorm.DB
}

func NewBiometricDataRepository(db *gorm.DB) BiometricDataRepository {
	baseRepo := NewBaseRepository[models.BiometricData](db)
	return &biometricRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetBiometricRecordsByAlertID retrieves a biometric by their AlertID.
func (r *biometricRepository) GetBiometricRecordsByAlertID(id uuid.UUID) ([]*models.BiometricData, error) {
	var biometrics []*models.BiometricData
	if err := r.db.Where("alert_id = ?", id).Find(&biometrics).Error; err != nil {
		return nil, err
	}
	return biometrics, nil
}
