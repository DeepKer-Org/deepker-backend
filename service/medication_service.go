package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type MedicationService interface {
	CreateMedication(medicationDTO *dto.MedicationCreateDTO) error
	GetMedicationByID(id uuid.UUID) (*dto.MedicationDTO, error)
	GetAllMedications() ([]*dto.MedicationDTO, error)
	UpdateMedication(id uuid.UUID, medicationDTO *dto.MedicationUpdateDTO) error
	DeleteMedication(id uuid.UUID) error
}

type medicationService struct {
	repo  repository.MedicationRepository
	cache *redis.CacheManager
}

func NewMedicationService(repo repository.MedicationRepository, cache *redis.CacheManager) MedicationService {
	return &medicationService{repo: repo, cache: cache}
}

func (s *medicationService) CreateMedication(medicationDTO *dto.MedicationCreateDTO) error {
	medication := &models.Medication{
		PatientID:   medicationDTO.PatientID,
		Name:        medicationDTO.Name,
		StartDate:   medicationDTO.StartDate,
		EndDate:     medicationDTO.EndDate,
		Dosage:      medicationDTO.Dosage,
		Periodicity: medicationDTO.Periodicity,
	}

	err := s.repo.CreateMedication(medication)
	if err != nil {
		log.Printf("Failed to create medication: %v", err)
		return err
	}
	log.Println("Medication created successfully with MedicationID:", medication.MedicationID)

	// Invalidate cache for all medications
	_ = s.cache.Delete(context.Background(), "medications:all")
	return nil
}

func (s *medicationService) GetMedicationByID(id uuid.UUID) (*dto.MedicationDTO, error) {
	ctx := context.Background()
	cacheKey := "medication:" + id.String()

	// Attempt to fetch from cache
	var medication dto.MedicationDTO
	found, err := s.cache.Get(ctx, cacheKey, &medication)
	if err != nil {
		log.Printf("Error accessing cache for MedicationID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for medication with MedicationID:", id)
		return &medication, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching medication with MedicationID:", id)
	dbMedication, err := s.repo.GetMedicationByID(id)
	if err != nil {
		return nil, err
	}
	if dbMedication == nil {
		log.Println("No medication found with MedicationID:", id)
		return nil, nil
	}

	medication = *dto.MapMedicationToDTO(dbMedication)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, medication); err != nil {
		log.Printf("Failed to cache medication: %v", err)
	}

	return &medication, nil
}

func (s *medicationService) GetAllMedications() ([]*dto.MedicationDTO, error) {
	ctx := context.Background()
	cacheKey := "medications:all"

	// Attempt to fetch from cache
	var medications []*dto.MedicationDTO
	found, err := s.cache.Get(ctx, cacheKey, &medications)
	if err != nil {
		log.Printf("Error accessing cache for all medications: %v", err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for all medications")
		return medications, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching all medications")
	dbMedications, err := s.repo.GetAllMedications()
	if err != nil {
		return nil, err
	}

	medications = dto.MapMedicationsToDTOs(dbMedications)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, medications); err != nil {
		log.Printf("Failed to cache all medications: %v", err)
	}

	return medications, nil
}

func (s *medicationService) UpdateMedication(id uuid.UUID, medicationDTO *dto.MedicationUpdateDTO) error {
	log.Println("Updating medication with MedicationID:", id)

	// Fetch medication from database
	medication, err := s.repo.GetMedicationByID(id)
	if err != nil {
		log.Printf("Error retrieving medication: %v", err)
		return err
	}
	if medication == nil {
		log.Printf("Medication not found with MedicationID: %v", id)
		return gorm.ErrRecordNotFound
	}

	// Update medication fields
	medication.PatientID = medicationDTO.PatientID
	medication.Name = medicationDTO.Name
	medication.StartDate = medicationDTO.StartDate
	medication.EndDate = medicationDTO.EndDate
	medication.Dosage = medicationDTO.Dosage
	medication.Periodicity = medicationDTO.Periodicity

	err = s.repo.UpdateMedication(medication)
	if err != nil {
		log.Printf("Failed to update medication: %v", err)
		return err
	}
	log.Println("Medication updated successfully with MedicationID:", medication.MedicationID)

	// Invalidate cache for the updated medication and all medications
	_ = s.cache.Delete(context.Background(), "medication:"+id.String(), "medications:all")
	return nil
}

func (s *medicationService) DeleteMedication(id uuid.UUID) error {
	log.Println("Deleting medication with MedicationID:", id)
	err := s.repo.DeleteMedication(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Medication not found with MedicationID:", id)
			return nil
		}
		log.Printf("Failed to delete medication: %v", err)
		return err
	}
	log.Println("Medication deleted successfully with MedicationID:", id)

	// Invalidate cache for the deleted medication and all medications
	_ = s.cache.Delete(context.Background(), "medication:"+id.String(), "medications:all")
	return nil
}
