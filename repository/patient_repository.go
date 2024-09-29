package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type PatientRepository interface {
	BaseRepository[models.Patient]
}

type patientRepository struct {
	BaseRepository[models.Patient]
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	baseRepo := NewBaseRepository[models.Patient](db)
	return &patientRepository{baseRepo}
}
