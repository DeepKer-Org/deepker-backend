package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// ImprovedPatientService demonstrates the new service pattern with generic base
type ImprovedPatientService struct {
	*BaseService[models.Patient, dto.PatientCreateDTO, dto.PatientUpdateDTO]
	repo  repository.BaseRepositoryInterface[models.Patient]
	cache *redis.CacheManager
}

// NewImprovedPatientService creates a new improved patient service
func NewImprovedPatientService(
	repo repository.BaseRepositoryInterface[models.Patient], 
	cacheManager *redis.CacheManager,
) *ImprovedPatientService {
	baseService := NewBaseService[models.Patient, dto.PatientCreateDTO, dto.PatientUpdateDTO](repo, cacheManager)
	
	return &ImprovedPatientService{
		BaseService: baseService,
		repo:        repo,
		cache:       cacheManager,
	}
}

// Create implements the CRUDService interface
func (s *ImprovedPatientService) Create(createDTO *dto.PatientCreateDTO) error {
	return s.BaseService.Create(createDTO, dto.MapCreateDTOToPatient)
}

// GetByID implements the CRUDService interface
func (s *ImprovedPatientService) GetByID(id uuid.UUID) (interface{}, error) {
	cacheKey := s.getCacheKey(id)
	
	return s.BaseService.GetByID(id, "patient_id", cacheKey, func(patient *models.Patient) interface{} {
		return dto.MapPatientToDTO(patient)
	})
}

// GetAll implements the CRUDService interface
func (s *ImprovedPatientService) GetAll() (interface{}, error) {
	cacheKey := s.getAllCacheKey()
	
	return s.BaseService.GetAll(cacheKey, func(patients []*models.Patient) interface{} {
		return dto.MapPatientsToDTOs(patients)
	})
}

// Update implements the CRUDService interface
func (s *ImprovedPatientService) Update(id uuid.UUID, updateDTO *dto.PatientUpdateDTO) error {
	return s.BaseService.Update(id, updateDTO, "patient_id", func(updateDTO *dto.PatientUpdateDTO, existing *models.Patient) *models.Patient {
		return dto.MapUpdateDTOToPatient(updateDTO, existing)
	})
}

// Delete implements the CRUDService interface
func (s *ImprovedPatientService) Delete(id uuid.UUID) error {
	return s.BaseService.Delete(id, "patient_id")
}

// Additional domain-specific methods
func (s *ImprovedPatientService) GetPatientByDNI(dni string) (*dto.PatientDTO, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("patient:dni:%s", dni)
	
	// Try cache first
	if s.cache != nil && s.cache.IsEnabled() {
		var cached dto.PatientDTO
		found, err := s.cache.Get(ctx, cacheKey, &cached)
		if err != nil {
			log.Printf("Error accessing cache for DNI %s: %v", dni, err)
		} else if found {
			log.Printf("Cache hit for patient with DNI: %s", dni)
			return &cached, nil
		}
	}

	// Fetch from database
	patients, err := s.repo.GetByField("dni", dni)
	if err != nil {
		return nil, err
	}
	
	if len(patients) == 0 {
		return nil, nil
	}

	patientDTO := dto.MapPatientToDTO(patients[0])

	// Cache the result
	if s.cache != nil && s.cache.IsEnabled() {
		if err := s.cache.Set(ctx, cacheKey, patientDTO); err != nil {
			log.Printf("Failed to cache patient: %v", err)
		}
	}

	return patientDTO, nil
}

func (s *ImprovedPatientService) GetAllPatientLocations() ([]string, error) {
	ctx := context.Background()
	cacheKey := "patients:locations:all"
	
	// Try cache first
	if s.cache != nil && s.cache.IsEnabled() {
		var cached []string
		found, err := s.cache.Get(ctx, cacheKey, &cached)
		if err != nil {
			log.Printf("Error accessing cache for patient locations: %v", err)
		} else if found {
			log.Printf("Cache hit for patient locations")
			return cached, nil
		}
	}

	// This would need to be implemented based on your patient model structure
	// For now, returning empty slice as placeholder
	locations := []string{}

	// Cache the result
	if s.cache != nil && s.cache.IsEnabled() {
		if err := s.cache.Set(ctx, cacheKey, locations); err != nil {
			log.Printf("Failed to cache patient locations: %v", err)
		}
	}

	return locations, nil
}

// Cache key helpers
func (s *ImprovedPatientService) getCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("patient:%s", id.String())
}

func (s *ImprovedPatientService) getAllCacheKey() string {
	return "patients:all"
}

// Compile-time interface check (commented out until controller package is imported)
// var _ controller.CRUDServiceInterface[dto.PatientCreateDTO, dto.PatientUpdateDTO] = &ImprovedPatientService{}