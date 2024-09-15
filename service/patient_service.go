package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/repository"
	"github.com/gocql/gocql"
	"time"
)

type PatientService interface {
	CreatePatient(patient *models.Patient) error
	GetAllPatients() ([]*models.Patient, error)
	GetPatientByID(id gocql.UUID) (*models.Patient, error)
	UpdatePatient(patient *models.Patient) error
	DeletePatient(id gocql.UUID) error
}

type patientService struct {
	repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patient *models.Patient) error {
	if patient.CreatedAt == 0 {
		patient.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return s.repo.Create(patient)
}

func (s *patientService) GetAllPatients() ([]*models.Patient, error) {
	return s.repo.GetAll()
}

func (s *patientService) GetPatientByID(id gocql.UUID) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *patientService) UpdatePatient(patient *models.Patient) error {
	return s.repo.Update(patient)
}

func (s *patientService) DeletePatient(id gocql.UUID) error {
	return s.repo.Delete(id)
}
