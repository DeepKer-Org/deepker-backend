package repository

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComputerDiagnosticRepository interface {
	BaseRepository[models.ComputerDiagnostic]
	GetComputerDiagnosticsByAlertID(alertID uuid.UUID) ([]*models.ComputerDiagnostic, error)
}

type computerDiagnosticRepository struct {
	BaseRepository[models.ComputerDiagnostic]
	db *gorm.DB
}

func NewComputerDiagnosticRepository(db *gorm.DB) ComputerDiagnosticRepository {
	baseRepo := NewBaseRepository[models.ComputerDiagnostic](db)
	return &computerDiagnosticRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetComputerDiagnosticsByAlertID retrieves a computerDiagnosis by their AlertID.
func (r *computerDiagnosticRepository) GetComputerDiagnosticsByAlertID(alertID uuid.UUID) ([]*models.ComputerDiagnostic, error) {
	var computerDiagnostics []*models.ComputerDiagnostic
	if err := r.db.Where("alert_id = ?", alertID).Find(&computerDiagnostics).Error; err != nil {
		return nil, err
	}
	return computerDiagnostics, nil
}
