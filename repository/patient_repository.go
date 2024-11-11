package repository

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"errors"
	"gorm.io/gorm"
	"strings"
)

// PatientRepository interface includes the specific methods and embeds BaseRepository
type PatientRepository interface {
	BaseRepository[models.Patient]
	GetPatientByDNI(dni string) (*models.Patient, error)
	GetAllPaginatedWithFilters(offset int, limit int, filters dto.PatientFilter) ([]*models.Patient, int64, error)
}

type patientRepository struct {
	BaseRepository[models.Patient]
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	baseRepo := NewBaseRepository[models.Patient](db)
	return &patientRepository{
		BaseRepository: baseRepo,
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
		Preload("MedicalVisits").
		Preload("MonitoringDevice").
		Where(primaryKey+" = ?", id).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
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
		Preload("MedicalVisits").
		Preload("MonitoringDevice").
		Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *patientRepository) GetPatientByDNI(dni string) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.
		Preload("Comorbidities").
		Preload("Medications").
		Preload("Doctors").
		Preload("Alerts").
		Where("dni = ?", dni).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &patient, nil
}

func applyPatientFilters(query *gorm.DB, filters dto.PatientFilter) *gorm.DB {
	// Apply filters dynamically

	// Basic filters from the patient table
	if filters.Name != "" {
		query = query.Where("LOWER(patients.name) LIKE ?", "%"+strings.ToLower(filters.Name)+"%")
	}
	if filters.DNI != "" {
		query = query.Where("patients.dni = ?", filters.DNI)
	}
	if filters.Age != 0 {
		query = query.Where("patients.age = ?", filters.Age)
	}
	if filters.Location != "" {
		query = query.Where("patients.location LIKE ?", "%"+filters.Location+"%")
	}

	// Join the doctor_patients table to filter by doctor_id
	if filters.DoctorID != "" {
		query = query.Joins("JOIN doctor_patients dp ON dp.patient_id = patients.patient_id").
			Where("dp.doctor_id = ?", filters.DoctorID)
	}

	// Join the monitoring_devices table to filter by device_id
	if filters.DeviceID != "" {
		query = query.Joins("JOIN monitoring_devices md ON md.patient_id = patients.patient_id").
			Where("md.device_id = ?", filters.DeviceID)
	}

	// Join the comorbidities table to filter by comorbidity name
	if filters.ComorbidityName != "" {
		query = query.Joins("JOIN comorbidities c ON c.patient_id = patients.patient_id").
			Where("LOWER(c.comorbidity) LIKE ?", "%"+strings.ToLower(filters.ComorbidityName)+"%")
	}

	// Join the medical_visits table to filter by entry_date and discharge_date

	// Filter by EntryDate (patients who are currently in the medical center)
	if filters.EntryDate != "" {
		query = query.Joins("JOIN medical_visits mv ON mv.patient_id = patients.patient_id").
			Where("DATE(mv.entry_date) = ? AND mv.discharge_date IS NULL", filters.EntryDate)
	}

	if filters.DischargeDate != "" {
		query = query.Joins("JOIN medical_visits mv2 ON mv2.patient_id = patients.patient_id").
			Where("DATE(mv2.discharge_date) = ? AND mv2.discharge_date = (SELECT MAX(mv3.discharge_date) FROM medical_visits mv3 WHERE mv3.patient_id = patients.patient_id AND mv3.discharge_date IS NOT NULL)", filters.DischargeDate)
	}

	return query
}

func (r *patientRepository) GetAllPaginatedWithFilters(offset int, limit int, filters dto.PatientFilter) ([]*models.Patient, int64, error) {
	var patients []*models.Patient
	var totalCount int64

	// Create base query
	query := r.db.Model(&models.Patient{})

	// Apply filters using the helper function
	query = applyPatientFilters(query, filters)

	// Count the total number of patients with filters
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated results with preloading relationships
	if err := query.
		Preload("Comorbidities").
		Preload("Medications").
		Preload("Doctors").
		Preload("Alerts").
		Preload("MedicalVisits").
		Preload("MonitoringDevice").
		Offset(offset).
		Limit(limit).
		Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, totalCount, nil
}
