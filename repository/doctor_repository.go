package repository

import (
	"biometric-data-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DoctorRepository interface includes the specific methods and embeds BaseRepository
type DoctorRepository interface {
	BaseRepository[models.Doctor]
	GetDoctorByUserID(userID uuid.UUID) (*models.Doctor, error)
	GetDoctorsByIDs(ids []uuid.UUID) ([]*models.Doctor, error)
	GetDoctorByDNI(dni string) (*models.Doctor, error)
	GetDoctorsByAlertID(alertID uuid.UUID) ([]*models.Doctor, error)
}

// doctorRepository struct embeds the baseRepository for common CRUD operations
type doctorRepository struct {
	BaseRepository[models.Doctor]
	db *gorm.DB
}

// NewDoctorRepository creates a new instance of DoctorRepository
func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	baseRepo := NewBaseRepository[models.Doctor](db)
	return &doctorRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetDoctorByUserID retrieves a doctor by their UserID.
func (r *doctorRepository) GetDoctorByUserID(userID uuid.UUID) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.db.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
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

// GetDoctorByDNI retrieves a doctor by their DNI.
func (r *doctorRepository) GetDoctorByDNI(dni string) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.db.
		Preload("User").
		Where("dni = ?", dni).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

// GetDoctorsByAlertID retrieves doctors associated with a specific alert ID.
func (r *doctorRepository) GetDoctorsByAlertID(alertID uuid.UUID) ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	if err := r.db.Joins("JOIN doctor_alerts ON doctor_alerts.doctor_id = doctors.doctor_id").
		Where("doctor_alerts.alert_id = ?", alertID).Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
