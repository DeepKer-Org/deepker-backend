package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

// PatientRepository interface includes the specific methods and embeds BaseRepository
type PatientRepository interface {
	BaseRepository[models.Patient]
}

type patientRepository struct {
	BaseRepository[models.Patient]
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{
		BaseRepository: NewBaseRepository[models.Patient](db),
		db:             db,
	}
}

func (r *patientRepository) GetByID(id interface{}, primaryKey string) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.
		Preload("Comorbidities").
		Preload("Medications").
		Preload("Doctors").
		Preload("Alerts").
		Where(primaryKey+" = ?", id).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil
}

func (r *patientRepository) GetAll() ([]*models.Patient, error) {
	var patients []*models.Patient
	if err := r.db.
		Preload("Comorbidities").
		Preload("Medications").
		Preload("Doctors").
		Preload("Alerts").
		Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
