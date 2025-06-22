package service

import (
	"biometric-data-backend/redis"  
	"biometric-data-backend/repository"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CacheableService interface defines caching operations
type CacheableService[T any, CreateDTO any, UpdateDTO any] interface {
	GetCacheKey(id uuid.UUID) string
	GetAllCacheKey() string
	MapEntityToDTO(entity *T) interface{}
	MapCreateDTOToEntity(dto *CreateDTO) *T
	MapUpdateDTOToEntity(dto *UpdateDTO, entity *T) *T
}

// BaseService provides common CRUD operations with caching
type BaseService[T any, CreateDTO any, UpdateDTO any] struct {
	repo  repository.BaseRepositoryInterface[T]
	cache *redis.CacheManager
}

// NewBaseService creates a new base service
func NewBaseService[T any, CreateDTO any, UpdateDTO any](
	repo repository.BaseRepositoryInterface[T],
	cache *redis.CacheManager,
) *BaseService[T, CreateDTO, UpdateDTO] {
	return &BaseService[T, CreateDTO, UpdateDTO]{
		repo:  repo,
		cache: cache,
	}
}

// Create creates a new entity
func (s *BaseService[T, CreateDTO, UpdateDTO]) Create(dto *CreateDTO, mapper func(*CreateDTO) *T) error {
	entity := mapper(dto)
	
	err := s.repo.Create(entity)
	if err != nil {
		log.Printf("Failed to create entity: %v", err)
		return err
	}

	// Invalidate cache
	s.invalidateCache("")
	
	log.Printf("Entity created successfully")
	return nil
}

// GetByID retrieves an entity by ID with caching
func (s *BaseService[T, CreateDTO, UpdateDTO]) GetByID(
	id uuid.UUID,
	idField string,
	cacheKey string,
	mapper func(*T) interface{},
) (interface{}, error) {
	ctx := context.Background()
	
	// Try cache first
	if s.cache != nil && s.cache.IsEnabled() {
		var cached interface{}
		found, err := s.cache.Get(ctx, cacheKey, &cached)
		if err != nil {
			log.Printf("Error accessing cache for ID %s: %v", id, err)
		} else if found {
			log.Printf("Cache hit for entity with ID: %s", id)
			return cached, nil
		}
	}

	// Fetch from database
	log.Printf("Fetching entity with ID: %s", id)
	entity, err := s.repo.GetByID(id, idField)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}

	// Map to DTO
	dto := mapper(entity)

	// Cache the result
	if s.cache != nil && s.cache.IsEnabled() {
		if err := s.cache.Set(ctx, cacheKey, dto); err != nil {
			log.Printf("Failed to cache entity: %v", err)
		}
	}

	return dto, nil
}

// GetAll retrieves all entities with caching
func (s *BaseService[T, CreateDTO, UpdateDTO]) GetAll(
	cacheKey string,
	mapper func([]*T) interface{},
) (interface{}, error) {
	ctx := context.Background()
	
	// Try cache first
	if s.cache != nil && s.cache.IsEnabled() {
		var cached interface{}
		found, err := s.cache.Get(ctx, cacheKey, &cached)
		if err != nil {
			log.Printf("Error accessing cache for GetAll: %v", err)
		} else if found {
			log.Printf("Cache hit for GetAll entities")
			return cached, nil
		}
	}

	// Fetch from database
	log.Printf("Fetching all entities")
	entities, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Map to DTOs
	dtos := mapper(entities)

	// Cache the result
	if s.cache != nil && s.cache.IsEnabled() {
		if err := s.cache.Set(ctx, cacheKey, dtos); err != nil {
			log.Printf("Failed to cache entities: %v", err)
		}
	}

	return dtos, nil
}

// Update updates an entity
func (s *BaseService[T, CreateDTO, UpdateDTO]) Update(
	id uuid.UUID,
	dto *UpdateDTO,
	idField string,
	mapper func(*UpdateDTO, *T) *T,
) error {
	log.Printf("Updating entity with ID: %s", id)
	
	// Get existing entity
	entity, err := s.repo.GetByID(id, idField)
	if err != nil {
		log.Printf("Error fetching entity: %v", err)
		return err
	}
	if entity == nil {
		log.Printf("Entity not found with ID: %v", id)
		return gorm.ErrRecordNotFound
	}

	// Map update
	updatedEntity := mapper(dto, entity)
	
	// Update in database
	err = s.repo.Update(updatedEntity, idField, id)
	if err != nil {
		log.Printf("Failed to update entity: %v", err)
		return err
	}

	// Invalidate cache
	s.invalidateCache(id.String())
	
	log.Printf("Entity updated successfully with ID: %s", id)
	return nil
}

// Delete deletes an entity
func (s *BaseService[T, CreateDTO, UpdateDTO]) Delete(id uuid.UUID, idField string) error {
	log.Printf("Deleting entity with ID: %s", id)
	
	err := s.repo.Delete(id, idField)
	if err != nil {
		log.Printf("Failed to delete entity: %v", err)
		return err
	}

	// Invalidate cache
	s.invalidateCache(id.String())
	
	log.Printf("Entity deleted successfully with ID: %s", id)
	return nil
}

// invalidateCache removes cached entries
func (s *BaseService[T, CreateDTO, UpdateDTO]) invalidateCache(entityID string) {
	if s.cache == nil || !s.cache.IsEnabled() {
		return
	}

	ctx := context.Background()
	keys := []string{"entities:all"}
	
	if entityID != "" {
		keys = append(keys, fmt.Sprintf("entity:%s", entityID))
	}

	if err := s.cache.Delete(ctx, keys...); err != nil {
		log.Printf("Failed to invalidate cache: %v", err)
	}
}

// Exists checks if an entity exists
func (s *BaseService[T, CreateDTO, UpdateDTO]) Exists(field string, value interface{}) (bool, error) {
	return s.repo.Exists(field, value)
}

// GetByField retrieves entities by a specific field
func (s *BaseService[T, CreateDTO, UpdateDTO]) GetByField(field string, value interface{}) ([]*T, error) {
	return s.repo.GetByField(field, value)
}