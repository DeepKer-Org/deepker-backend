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

type BiometricDataService interface {
	CreateBiometricData(biometricDTO *dto.BiometricDataCreateDTO) error
	GetBiometricDataByID(id uuid.UUID) (*dto.BiometricDataDTO, error)
	GetAllBiometricData() ([]*dto.BiometricDataDTO, error)
	UpdateBiometricData(id uuid.UUID, biometricDTO *dto.BiometricDataUpdateDTO) error
	DeleteBiometricData(id uuid.UUID) error
}

type biometricService struct {
	repo  repository.BiometricDataRepository
	cache *redis.CacheManager
}

func NewBiometricDataService(repo repository.BiometricDataRepository, cache *redis.CacheManager) BiometricDataService {
	return &biometricService{repo: repo, cache: cache}
}

func (s *biometricService) CreateBiometricData(biometricDTO *dto.BiometricDataCreateDTO) error {
	biometric := &models.BiometricData{
		O2Saturation: biometricDTO.O2Saturation,
		HeartRate:    biometricDTO.HeartRate,
	}

	err := s.repo.Create(biometric)
	if err != nil {
		log.Printf("Failed to create biometric: %v", err)
		return err
	}
	log.Println("BiometricData created successfully with BiometricDataID:", biometric.BiometricDataID)

	// Invalidate cache for all biometric data
	_ = s.cache.Delete(context.Background(), "biometric_data:all")
	return nil
}

func (s *biometricService) GetBiometricDataByID(id uuid.UUID) (*dto.BiometricDataDTO, error) {
	ctx := context.Background()
	cacheKey := "biometric_data:" + id.String()

	// Attempt to fetch from cache
	var biometric dto.BiometricDataDTO
	found, err := s.cache.Get(ctx, cacheKey, &biometric)
	if err != nil {
		log.Printf("Error accessing cache for BiometricDataID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for biometric with BiometricDataID:", id)
		return &biometric, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching biometric with BiometricDataID:", id)
	dbBiometric, err := s.repo.GetByID(id, "biometric_data_id")
	if err != nil {
		return nil, err
	}
	if dbBiometric == nil {
		log.Println("No biometric found with BiometricDataID:", id)
		return nil, nil
	}

	biometric = *dto.MapBiometricDataToDTO(dbBiometric)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, biometric); err != nil {
		log.Printf("Failed to cache biometric: %v", err)
	}

	return &biometric, nil
}

func (s *biometricService) GetAllBiometricData() ([]*dto.BiometricDataDTO, error) {
	ctx := context.Background()
	cacheKey := "biometric_data:all"

	// Attempt to fetch from cache
	var biometrics []*dto.BiometricDataDTO
	found, err := s.cache.Get(ctx, cacheKey, &biometrics)
	if err != nil {
		log.Printf("Error accessing cache for all biometric data: %v", err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for all biometric data")
		return biometrics, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching all biometric data")
	dbBiometrics, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	biometrics = dto.MapBiometricDataToDTOs(dbBiometrics)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, biometrics); err != nil {
		log.Printf("Failed to cache all biometric data: %v", err)
	}

	return biometrics, nil
}

func (s *biometricService) UpdateBiometricData(id uuid.UUID, biometricDTO *dto.BiometricDataUpdateDTO) error {
	log.Println("Updating biometric with BiometricDataID:", id)

	// Fetch biometric from database
	biometric, err := s.repo.GetByID(id, "biometric_data_id")
	if err != nil {
		log.Printf("Error retrieving biometric: %v", err)
		return err
	}
	if biometric == nil {
		log.Printf("BiometricData not found with BiometricDataID: %v", id)
		return gorm.ErrRecordNotFound
	}

	// Update biometric data
	biometric.O2Saturation = biometricDTO.O2Saturation
	biometric.HeartRate = biometricDTO.HeartRate

	err = s.repo.Update(biometric, "biometric_data_id", id)
	if err != nil {
		log.Printf("Failed to update biometric: %v", err)
		return err
	}
	log.Println("BiometricData updated successfully with BiometricDataID:", biometric.BiometricDataID)

	// Invalidate cache for the updated biometric and all biometric data
	_ = s.cache.Delete(context.Background(), "biometric_data:"+id.String(), "biometric_data:all")
	return nil
}

func (s *biometricService) DeleteBiometricData(id uuid.UUID) error {
	log.Println("Deleting biometric with BiometricDataID:", id)
	err := s.repo.Delete(id, "biometric_data_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("BiometricData not found with BiometricDataID:", id)
			return nil
		}
		log.Printf("Failed to delete biometric: %v", err)
		return err
	}
	log.Println("BiometricData deleted successfully with BiometricDataID:", id)

	// Invalidate cache for the deleted biometric and all biometric data
	_ = s.cache.Delete(context.Background(), "biometric_data:"+id.String(), "biometric_data:all")
	return nil
}
