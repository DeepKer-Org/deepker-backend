package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type PatientService interface {
	CreatePatient(patientDTO *dto.PatientCreateDTO) error
	GetPatientByID(id uuid.UUID) (*dto.PatientDTO, error)
	GetPatientByDNI(dni string) (*dto.PatientDTO, error)
	GetAllPatients(page int, limit int, filters dto.PatientFilter) ([]*dto.PatientDTO, int, error)
	UpdatePatient(id uuid.UUID, patientDTO *dto.PatientUpdateDTO) error
	DeletePatient(id uuid.UUID) error
}

type patientService struct {
	repo  repository.PatientRepository
	cache *redis.CacheManager
}

func NewPatientService(repo repository.PatientRepository, cache *redis.CacheManager) PatientService {
	return &patientService{repo: repo, cache: cache}
}

func (s *patientService) CreatePatient(patientDTO *dto.PatientCreateDTO) error {
	patient := dto.MapCreateDTOToPatient(patientDTO)
	err := s.repo.Create(patient)
	if err != nil {
		log.Printf("Failed to create patient: %v", err)
		return err
	}
	log.Println("Patient created successfully with PatientID:", patient.PatientID)

	// Invalidate cache for all patients
	_ = s.cache.Delete(context.Background(), "patients:all")
	return nil
}

func (s *patientService) GetPatientByID(id uuid.UUID) (*dto.PatientDTO, error) {
	ctx := context.Background()
	cacheKey := "patient:" + id.String()

	// Attempt to fetch from cache
	var patient dto.PatientDTO
	found, err := s.cache.Get(ctx, cacheKey, &patient)
	if err != nil {
		log.Printf("Error accessing cache for PatientID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for patient with PatientID:", id)
		return &patient, nil
	}

	log.Println("Fetching patient with PatientID:", id)
	dbPatient, err := s.repo.GetByID(id, "patient_id")
	if err != nil {
		log.Printf("Error fetching patient: %v", err)
		return nil, err
	}
	if dbPatient == nil {
		log.Println("No patient found with PatientID:", id)
		return nil, nil
	}

	patient = *dto.MapPatientToDTO(dbPatient)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, patient); err != nil {
		log.Printf("Failed to cache patient: %v", err)
	}

	return &patient, nil
}

func (s *patientService) GetPatientByDNI(dni string) (*dto.PatientDTO, error) {
	ctx := context.Background()
	cacheKey := "patient:dni:" + dni

	// Attempt to fetch from cache
	var patient dto.PatientDTO
	found, err := s.cache.Get(ctx, cacheKey, &patient)
	if err != nil {
		log.Printf("Error accessing cache for Patient DNI %s: %v", dni, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for patient with DNI:", dni)
		return &patient, nil
	}

	log.Println("Fetching patient with DNI:", dni)
	dbPatient, err := s.repo.GetPatientByDNI(dni)
	if err != nil {
		log.Printf("Error fetching patient: %v", err)
		return nil, err
	}
	if dbPatient == nil {
		log.Println("No patient found with DNI:", dni)
		return nil, nil
	}

	patient = *dto.MapPatientToDTO(dbPatient)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, patient); err != nil {
		log.Printf("Failed to cache patient: %v", err)
	}

	return &patient, nil
}

func (s *patientService) GetAllPatients(page int, limit int, filters dto.PatientFilter) ([]*dto.PatientDTO, int, error) {
	ctx := context.Background()
	cacheKey := "patients:all:page=%d:limit=%d:filters=%v"

	// Attempt to fetch from cache
	var patients []*dto.PatientDTO
	var totalCount int
	found, err := s.cache.Get(ctx, fmt.Sprintf(cacheKey, page, limit, filters), &patients)
	if err != nil {
		log.Printf("Error accessing cache for patients: %v", err)
		return nil, 0, err
	}
	if found {
		log.Println("Cache hit for all patients")
		return patients, totalCount, nil
	}

	offset := (page - 1) * limit
	var dbPatients []*models.Patient
	var totalCount64 int64

	dbPatients, totalCount64, err = s.repo.GetAllPaginatedWithFilters(offset, limit, filters)
	if err != nil {
		log.Printf("Error fetching filtered patients: %v", err)
		return nil, 0, err
	}
	totalCount = int(totalCount64)
	patients = dto.MapPatientsToDTOs(dbPatients)

	// Store in cache
	if err := s.cache.Set(ctx, fmt.Sprintf(cacheKey, page, limit, filters), patients); err != nil {
		log.Printf("Failed to cache patients: %v", err)
	}

	return patients, totalCount, nil
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

	// Invalidate cache for the updated patient and all patients
	_ = s.cache.Delete(context.Background(), "patient:"+id.String(), "patients:all")
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

	// Invalidate cache for the deleted patient and all patients
	_ = s.cache.Delete(context.Background(), "patient:"+id.String(), "patients:all")
	return nil
}
