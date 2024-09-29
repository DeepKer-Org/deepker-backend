package repository

import (
	"biometric-data-backend/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BiometricDataRepository interface {
	CreateBiometricData(biometric *models.BiometricData) error
	GetBiometricDataByID(id uuid.UUID) (*models.BiometricData, error)
	GetBiometricRecordsByAlertID(id uuid.UUID) ([]*models.BiometricData, error)
	GetAllBiometricRecords() ([]*models.BiometricData, error)
	UpdateBiometricData(biometric *models.BiometricData) error
	DeleteBiometricData(id uuid.UUID) error
}

type biometricRepository struct {
	db *gorm.DB
}

func NewBiometricDataRepository(db *gorm.DB) BiometricDataRepository {
	return &biometricRepository{db}
}

// CreateBiometricData creates a new biometric record in the database.
func (r *biometricRepository) CreateBiometricData(biometric *models.BiometricData) error {
	if err := r.db.Create(biometric).Error; err != nil {
		return err
	}
	return nil
}

// GetBiometricDataByID retrieves a biometric by their BiometricDataID.
func (r *biometricRepository) GetBiometricDataByID(id uuid.UUID) (*models.BiometricData, error) {
	var biometric models.BiometricData
	if err := r.db.First(&biometric).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &biometric, nil
}

// GetBiometricRecordsByAlertID retrieves a biometric by their AlertID.
func (r *biometricRepository) GetBiometricRecordsByAlertID(id uuid.UUID) ([]*models.BiometricData, error) {
	var biometrics []*models.BiometricData
	if err := r.db.Where("alert_id = ?", id).Find(&biometrics).Error; err != nil {
		return nil, err
	}
	return biometrics, nil
}

// GetAllBiometricRecords retrieves all biometrics from the database.
func (r *biometricRepository) GetAllBiometricRecords() ([]*models.BiometricData, error) {
	var biometrics []*models.BiometricData
	if err := r.db.Find(&biometrics).Error; err != nil {
		return nil, err
	}
	return biometrics, nil
}

// UpdateBiometricData updates an existing biometric record in the database.
func (r *biometricRepository) UpdateBiometricData(biometric *models.BiometricData) error {
	if err := r.db.Save(biometric).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBiometricData deletes a biometric by their BiometricDataID.
func (r *biometricRepository) DeleteBiometricData(id uuid.UUID) error {
	if err := r.db.Delete(&models.BiometricData{}).Error; err != nil {
		return err
	}
	return nil
}
