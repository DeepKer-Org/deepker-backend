package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type ComputerDiagnosticRepository interface {
	CreateComputerDiagnostic(computerDiagnosis *models.ComputerDiagnostic) error
	GetComputerDiagnosticByID(id uint) (*models.ComputerDiagnostic, error)
	GetComputerDiagnosticsByAlertID(alertID string) ([]*models.ComputerDiagnostic, error)
	GetAllComputerDiagnostics() ([]*models.ComputerDiagnostic, error)
	UpdateComputerDiagnostic(computerDiagnosis *models.ComputerDiagnostic) error
	DeleteComputerDiagnostic(id uint) error
}

type computerDiagnosticRepository struct {
	db *gorm.DB
}

func NewComputerDiagnosticRepository(db *gorm.DB) ComputerDiagnosticRepository {
	return &computerDiagnosticRepository{db}
}

// CreateComputerDiagnostic creates a new computerDiagnosis record in the database.
func (r *computerDiagnosticRepository) CreateComputerDiagnostic(computerDiagnosis *models.ComputerDiagnostic) error {
	if err := r.db.Create(computerDiagnosis).Error; err != nil {
		return err
	}
	return nil
}

// GetComputerDiagnosticByID retrieves a computerDiagnosis by their ComputerDiagnosticID.
func (r *computerDiagnosticRepository) GetComputerDiagnosticByID(id uint) (*models.ComputerDiagnostic, error) {
	var computerDiagnosis models.ComputerDiagnostic
	if err := r.db.First(&computerDiagnosis, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &computerDiagnosis, nil
}

// GetComputerDiagnosticsByAlertID retrieves a computerDiagnosis by their AlertID.
func (r *computerDiagnosticRepository) GetComputerDiagnosticsByAlertID(alertID string) ([]*models.ComputerDiagnostic, error) {
	var computerDiagnostics []*models.ComputerDiagnostic
	if err := r.db.Where("alert_id = ?", alertID).Find(&computerDiagnostics).Error; err != nil {
		return nil, err
	}
	return computerDiagnostics, nil
}

// GetAllComputerDiagnostics retrieves all computerDiagnostics from the database.
func (r *computerDiagnosticRepository) GetAllComputerDiagnostics() ([]*models.ComputerDiagnostic, error) {
	var computerDiagnostics []*models.ComputerDiagnostic
	if err := r.db.Find(&computerDiagnostics).Error; err != nil {
		return nil, err
	}
	return computerDiagnostics, nil
}

// UpdateComputerDiagnostic updates an existing computerDiagnosis record in the database.
func (r *computerDiagnosticRepository) UpdateComputerDiagnostic(computerDiagnosis *models.ComputerDiagnostic) error {
	if err := r.db.Save(computerDiagnosis).Error; err != nil {
		return err
	}
	return nil
}

// DeleteComputerDiagnostic deletes a computerDiagnosis by their ComputerDiagnosticID.
func (r *computerDiagnosticRepository) DeleteComputerDiagnostic(id uint) error {
	if err := r.db.Delete(&models.ComputerDiagnostic{}, id).Error; err != nil {
		return err
	}
	return nil
}
