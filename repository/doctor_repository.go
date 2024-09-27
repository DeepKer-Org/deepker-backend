package repository

import (
	"biometric-data-backend/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	CreateDoctor(doctor *models.Doctor) error
	GetDoctorByID(id uuid.UUID) (*models.Doctor, error)
	GetDoctorsByIDs(ids []uuid.UUID) ([]*models.Doctor, error)
	GetDoctorsByAlertID(alertID uuid.UUID) ([]*models.Doctor, error)
	GetAllDoctors() ([]*models.Doctor, error)
	UpdateDoctor(doctor *models.Doctor) error
	DeleteDoctor(id uuid.UUID) error
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{db}
}

// CreateDoctor creates a new doctor record in the database.
func (r *doctorRepository) CreateDoctor(doctor *models.Doctor) error {
	if err := r.db.Create(doctor).Error; err != nil {
		return err
	}
	return nil
}

// GetDoctorByID retrieves a doctor by their DoctorID.
func (r *doctorRepository) GetDoctorByID(id uuid.UUID) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.db.First(&doctor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &doctor, nil
}

// GetDoctorsByIDs retrieves doctors by their DoctorIDs.
func (r *doctorRepository) GetDoctorsByIDs(ids []uuid.UUID) ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	if err := r.db.Where("doctor_id IN (?)", ids).Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

// GetDoctorsByAlertID retrieves doctors associated with a specific alert ID.
func (r *doctorRepository) GetDoctorsByAlertID(alertID uuid.UUID) ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	// Joining the doctor_alerts table to get the doctors for the given alert ID
	if err := r.db.Joins("JOIN doctor_alerts ON doctor_alerts.doctor_id = doctors.doctor_id").
		Where("doctor_alerts.alert_id = ?", alertID).Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

// GetAllDoctors retrieves all doctors from the database.
func (r *doctorRepository) GetAllDoctors() ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	if err := r.db.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

// UpdateDoctor updates an existing doctor record in the database.
func (r *doctorRepository) UpdateDoctor(doctor *models.Doctor) error {
	if err := r.db.Save(doctor).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDoctor deletes a doctor by their DoctorID.
func (r *doctorRepository) DeleteDoctor(id uuid.UUID) error {
	if err := r.db.Delete(&models.Doctor{}).Error; err != nil {
		return err
	}
	return nil
}
