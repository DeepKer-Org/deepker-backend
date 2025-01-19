package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type ComorbidityService interface {
	CreateComorbidity(comorbidityDTO *dto.ComorbidityCreateDTO) error
	GetComorbidityByID(id uuid.UUID) (*dto.ComorbidityDTO, error)
	GetAllComorbidities() ([]*dto.ComorbidityDTO, error)
	UpdateComorbidity(id uuid.UUID, comorbidityDTO *dto.ComorbidityUpdateDTO) error
	DeleteComorbidity(id uuid.UUID) error
}

type comorbidityService struct {
	repo  repository.ComorbidityRepository
	cache *redis.CacheManager
}

func NewComorbidityService(repo repository.ComorbidityRepository, cache *redis.CacheManager) ComorbidityService {
	return &comorbidityService{repo: repo, cache: cache}
}

func (s *comorbidityService) CreateComorbidity(comorbidityDTO *dto.ComorbidityCreateDTO) error {
	comorbidity := dto.MapCreateDTOToComorbidity(comorbidityDTO)
	err := s.repo.Create(comorbidity)
	if err != nil {
		log.Printf("Failed to create comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity created successfully with ComorbidityID:", comorbidity.ComorbidityID)
	_ = s.cache.Delete(context.Background(), "comorbidities:all")
	return nil
}

func (s *comorbidityService) GetComorbidityByID(id uuid.UUID) (*dto.ComorbidityDTO, error) {
	ctx := context.Background()
	cacheKey := "comorbidity:" + id.String()

	// Attempt to fetch from cache
	var comorbidity dto.ComorbidityDTO
	found, err := s.cache.Get(ctx, cacheKey, &comorbidity)
	if err != nil {
		return nil, err
	}
	if found {
		log.Println("Cache hit for comorbidity with ComorbidityID:", id)
		return &comorbidity, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching comorbidity with ComorbidityID:", id)
	dbComorbidity, err := s.repo.GetByID(id, "comorbidity_id")
	if err != nil {
		return nil, err
	}
	if dbComorbidity == nil {
		return nil, nil
	}

	comorbidity = *dto.MapComorbidityToDTO(dbComorbidity)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, comorbidity); err != nil {
		log.Printf("Failed to cache comorbidity: %v", err)
	}

	return &comorbidity, nil
}

func (s *comorbidityService) GetAllComorbidities() ([]*dto.ComorbidityDTO, error) {
	ctx := context.Background()
	cacheKey := "comorbidities:all"
	// Attempt to fetch from cache
	var comorbidities []*dto.ComorbidityDTO
	found, err := s.cache.Get(ctx, cacheKey, &comorbidities)
	if err != nil {
		return nil, err
	}
	if found {
		log.Println("Cache hit for all comorbidities")
		return comorbidities, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching all comorbidities")
	dbComorbidities, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	comorbidities = dto.MapComorbiditiesToDTOs(dbComorbidities)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, comorbidities); err != nil {
		log.Printf("Failed to cache comorbidities: %v", err)
	}

	return comorbidities, nil
}

func (s *comorbidityService) UpdateComorbidity(id uuid.UUID, comorbidityDTO *dto.ComorbidityUpdateDTO) error {
	log.Println("Updating comorbidity with ComorbidityID:", id)

	comorbidity, err := s.repo.GetByID(id, "comorbidity_id")
	if err != nil {
		log.Printf("Error fetching comorbidity: %v", err)
		return err
	}
	if comorbidity == nil {
		log.Printf("Comorbidity not found with ComorbidityID: %v", id)
		return gorm.ErrRecordNotFound
	}

	comorbidity = dto.MapUpdateDTOToComorbidity(comorbidityDTO, comorbidity)
	err = s.repo.Update(comorbidity, "comorbidity_id", id)
	if err != nil {
		log.Printf("Failed to update comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity updated successfully with ComorbidityID:", comorbidity.ComorbidityID)
	_ = s.cache.Delete(context.Background(), "comorbidity:"+id.String(), "comorbidities:all")
	return nil
}

func (s *comorbidityService) DeleteComorbidity(id uuid.UUID) error {
	log.Println("Deleting comorbidity with ComorbidityID:", id)
	err := s.repo.Delete(id, "comorbidity_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Comorbidity not found with ComorbidityID:", id)
			return nil
		}
		log.Printf("Failed to delete comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity deleted successfully with ComorbidityID:", id)
	_ = s.cache.Delete(context.Background(), "comorbidity:"+id.String(), "comorbidities:all")
	return nil
}
