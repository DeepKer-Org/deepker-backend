package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type BiometricRepository interface {
	CreateBiometric(biometric *models.Biometric) error
	GetBiometricByID(id uint) (*models.Biometric, error)
	GetAllBiometrics() ([]*models.Biometric, error)
	UpdateBiometric(biometric *models.Biometric) error
	DeleteBiometric(id uint) error
}

type biometricRepository struct {
	db *gorm.DB
}

func NewBiometricRepository(db *gorm.DB) BiometricRepository {
	return &biometricRepository{db}
}

// CreateBiometric creates a new biometric record in the database.
func (r *biometricRepository) CreateBiometric(biometric *models.Biometric) error {
	if err := r.db.Create(biometric).Error; err != nil {
		return err
	}
	return nil
}

// GetBiometricByID retrieves a biometric by their BiometricID.
func (r *biometricRepository) GetBiometricByID(id uint) (*models.Biometric, error) {
	var biometric models.Biometric
	if err := r.db.First(&biometric, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &biometric, nil
}

// GetAllBiometrics retrieves all biometrics from the database.
func (r *biometricRepository) GetAllBiometrics() ([]*models.Biometric, error) {
	var biometrics []*models.Biometric
	if err := r.db.Find(&biometrics).Error; err != nil {
		return nil, err
	}
	return biometrics, nil
}

// UpdateBiometric updates an existing biometric record in the database.
func (r *biometricRepository) UpdateBiometric(biometric *models.Biometric) error {
	if err := r.db.Save(biometric).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBiometric deletes a biometric by their BiometricID.
func (r *biometricRepository) DeleteBiometric(id uint) error {
	if err := r.db.Delete(&models.Biometric{}, id).Error; err != nil {
		return err
	}
	return nil
}
