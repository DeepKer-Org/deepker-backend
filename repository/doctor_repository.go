package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	CreateDoctor(doctor *models.Doctor) error
	GetDoctorByID(id uint) (*models.Doctor, error)
	GetAllDoctors() ([]*models.Doctor, error)
	UpdateDoctor(doctor *models.Doctor) error
	DeleteDoctor(id uint) error
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
func (r *doctorRepository) GetDoctorByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.db.First(&doctor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &doctor, nil
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
func (r *doctorRepository) DeleteDoctor(id uint) error {
	if err := r.db.Delete(&models.Doctor{}, id).Error; err != nil {
		return err
	}
	return nil
}
