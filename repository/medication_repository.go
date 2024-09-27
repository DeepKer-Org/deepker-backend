package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type MedicationRepository interface {
	CreateMedication(medication *models.Medication) error
	GetMedicationByID(id uint) (*models.Medication, error)
	GetAllMedications() ([]*models.Medication, error)
	UpdateMedication(medication *models.Medication) error
	DeleteMedication(id uint) error
}

type medicationRepository struct {
	db *gorm.DB
}

func NewMedicationRepository(db *gorm.DB) MedicationRepository {
	return &medicationRepository{db}
}

// CreateMedication creates a new medication record in the database.
func (r *medicationRepository) CreateMedication(medication *models.Medication) error {
	if err := r.db.Create(medication).Error; err != nil {
		return err
	}
	return nil
}

// GetMedicationByID retrieves a medication by their MedicationID.
func (r *medicationRepository) GetMedicationByID(id uint) (*models.Medication, error) {
	var medication models.Medication
	if err := r.db.First(&medication, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &medication, nil
}

// GetAllMedications retrieves all medications from the database.
func (r *medicationRepository) GetAllMedications() ([]*models.Medication, error) {
	var medications []*models.Medication
	if err := r.db.Find(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

// UpdateMedication updates an existing medication record in the database.
func (r *medicationRepository) UpdateMedication(medication *models.Medication) error {
	if err := r.db.Save(medication).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMedication deletes a medication by their MedicationID.
func (r *medicationRepository) DeleteMedication(id uint) error {
	if err := r.db.Delete(&models.Medication{}, id).Error; err != nil {
		return err
	}
	return nil
}
