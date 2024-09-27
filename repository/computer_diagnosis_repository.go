package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type ComputerDiagnosisRepository interface {
	CreateComputerDiagnosis(computerDiagnosis *models.ComputerDiagnosis) error
	GetComputerDiagnosisByID(id uint) (*models.ComputerDiagnosis, error)
	GetAllComputerDiagnoses() ([]*models.ComputerDiagnosis, error)
	UpdateComputerDiagnosis(computerDiagnosis *models.ComputerDiagnosis) error
	DeleteComputerDiagnosis(id uint) error
}

type computerDiagnosisRepository struct {
	db *gorm.DB
}

func NewComputerDiagnosisRepository(db *gorm.DB) ComputerDiagnosisRepository {
	return &computerDiagnosisRepository{db}
}

// CreateComputerDiagnosis creates a new computerDiagnosis record in the database.
func (r *computerDiagnosisRepository) CreateComputerDiagnosis(computerDiagnosis *models.ComputerDiagnosis) error {
	if err := r.db.Create(computerDiagnosis).Error; err != nil {
		return err
	}
	return nil
}

// GetComputerDiagnosisByID retrieves a computerDiagnosis by their ComputerDiagnosisID.
func (r *computerDiagnosisRepository) GetComputerDiagnosisByID(id uint) (*models.ComputerDiagnosis, error) {
	var computerDiagnosis models.ComputerDiagnosis
	if err := r.db.First(&computerDiagnosis, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &computerDiagnosis, nil
}

// GetAllComputerDiagnoses retrieves all computerDiagnoses from the database.
func (r *computerDiagnosisRepository) GetAllComputerDiagnoses() ([]*models.ComputerDiagnosis, error) {
	var computerDiagnoses []*models.ComputerDiagnosis
	if err := r.db.Find(&computerDiagnoses).Error; err != nil {
		return nil, err
	}
	return computerDiagnoses, nil
}

// UpdateComputerDiagnosis updates an existing computerDiagnosis record in the database.
func (r *computerDiagnosisRepository) UpdateComputerDiagnosis(computerDiagnosis *models.ComputerDiagnosis) error {
	if err := r.db.Save(computerDiagnosis).Error; err != nil {
		return err
	}
	return nil
}

// DeleteComputerDiagnosis deletes a computerDiagnosis by their ComputerDiagnosisID.
func (r *computerDiagnosisRepository) DeleteComputerDiagnosis(id uint) error {
	if err := r.db.Delete(&models.ComputerDiagnosis{}, id).Error; err != nil {
		return err
	}
	return nil
}
