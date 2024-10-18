package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type PatientService interface {
	CreatePatient(patientDTO *dto.PatientCreateDTO) error
	GetPatientByID(id uuid.UUID) (*dto.PatientDTO, error)
	GetPatientByDNI(dni string) (*dto.PatientDTO, error)
	GetAllPatients(page int, limit int) ([]*dto.PatientDTO, int, error)
	UpdatePatient(id uuid.UUID, patientDTO *dto.PatientUpdateDTO) error
	DeletePatient(id uuid.UUID) error
}

type patientService struct {
	repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patientDTO *dto.PatientCreateDTO) error {
	patient := dto.MapCreateDTOToPatient(patientDTO)
	err := s.repo.Create(patient)
	if err != nil {
		log.Printf("Failed to create patient: %v", err)
		return err
	}
	log.Println("Patient created successfully with PatientID:", patient.PatientID)
	return nil
}

func (s *patientService) GetPatientByID(id uuid.UUID) (*dto.PatientDTO, error) {
	log.Println("Fetching patient with PatientID:", id)
	patient, err := s.repo.GetByID(id, "patient_id")
	if err != nil {
		log.Printf("Error fetching patient: %v", err)
		return nil, err
	}
	if patient == nil {
		log.Println("No patient found with PatientID:", id)
		return nil, nil
	}

	return dto.MapPatientToDTO(patient), nil
}

func (s *patientService) GetPatientByDNI(dni string) (*dto.PatientDTO, error) {
	log.Println("Fetching patient with DNI:", dni)
	patient, err := s.repo.GetPatientByDNI(dni)
	if err != nil {
		log.Printf("Error fetching patient: %v", err)
		return nil, err
	}
	if patient == nil {
		log.Println("No patient found with DNI:", dni)
		return nil, nil
	}

	return dto.MapPatientToDTO(patient), nil
}

func (s *patientService) GetAllPatients(page int, limit int) ([]*dto.PatientDTO, int, error) {
	offset := (page - 1) * limit
	var patients []*models.Patient
	var totalCount int64
	var err error

	log.Println("Fetching all patients")
	err = s.repo.CountPatients(&totalCount)
	if err != nil {
		log.Printf("Error counting patients: %v", err)
		return nil, 0, err
	}
	patients, err = s.repo.GetAllPaginated(offset, limit)

	if err != nil {
		log.Printf("Error fetching patients: %v", err)
		return nil, 0, err
	}

	return dto.MapPatientsToDTOs(patients), int(totalCount), nil
}

func (s *patientService) UpdatePatient(id uuid.UUID, patientDTO *dto.PatientUpdateDTO) error {
	log.Println("Updating patient with PatientID:", id)

	patient, err := s.repo.GetByID(id, "patient_id")
	if err != nil {
		log.Printf("Error fetching patient: %v", err)
		return err
	}
	if patient == nil {
		log.Printf("Patient not found with PatientID: %v", id)
		return gorm.ErrRecordNotFound
	}

	patient = dto.MapUpdateDTOToPatient(patientDTO, patient)
	err = s.repo.Update(patient, "patient_id", id)
	if err != nil {
		log.Printf("Failed to update patient: %v", err)
		return err
	}
	log.Println("Patient updated successfully with PatientID:", patient.PatientID)
	return nil
}

func (s *patientService) DeletePatient(id uuid.UUID) error {
	log.Println("Deleting patient with PatientID:", id)
	err := s.repo.Delete(id, "patient_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Patient not found with PatientID:", id)
			return nil
		}
		log.Printf("Failed to delete patient: %v", err)
		return err
	}
	log.Println("Patient deleted successfully with PatientID:", id)
	return nil
}
