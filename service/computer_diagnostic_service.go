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

type ComputerDiagnosticService interface {
	CreateComputerDiagnostic(diagnosisDTO *dto.ComputerDiagnosticCreateDTO) error
	GetComputerDiagnosticByID(id uuid.UUID) (*dto.ComputerDiagnosticDTO, error)
	GetAllComputerDiagnostics() ([]*dto.ComputerDiagnosticDTO, error)
	UpdateComputerDiagnostic(id uuid.UUID, diagnosisDTO *dto.ComputerDiagnosticUpdateDTO) error
	DeleteComputerDiagnostic(id uuid.UUID) error
}

type computerDiagnosticService struct {
	repo  repository.ComputerDiagnosticRepository
	cache *redis.CacheManager
}

func NewComputerDiagnosticService(repo repository.ComputerDiagnosticRepository, cache *redis.CacheManager) ComputerDiagnosticService {
	return &computerDiagnosticService{repo: repo, cache: cache}
}

func (s *computerDiagnosticService) CreateComputerDiagnostic(diagnosisDTO *dto.ComputerDiagnosticCreateDTO) error {
	diagnosis := &models.ComputerDiagnostic{
		Diagnosis:  diagnosisDTO.Diagnosis,
		Percentage: diagnosisDTO.Percentage,
	}

	err := s.repo.Create(diagnosis)
	if err != nil {
		log.Printf("Failed to create computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis created successfully with DiagnosisID:", diagnosis.DiagnosticID)

	// Invalidate cache for all diagnostics
	_ = s.cache.Delete(context.Background(), "computer_diagnostics:all")
	return nil
}

func (s *computerDiagnosticService) GetComputerDiagnosticByID(id uuid.UUID) (*dto.ComputerDiagnosticDTO, error) {
	ctx := context.Background()
	cacheKey := "computer_diagnostic:" + id.String()

	// Attempt to fetch from cache
	var diagnosis dto.ComputerDiagnosticDTO
	found, err := s.cache.Get(ctx, cacheKey, &diagnosis)
	if err != nil {
		log.Printf("Error accessing cache for DiagnosisID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for computer diagnosis with DiagnosisID:", id)
		return &diagnosis, nil
	}

	log.Println("Fetching computer diagnosis with DiagnosisID:", id)
	dbDiagnosis, err := s.repo.GetByID(id, "diagnostic_id")
	if err != nil {
		log.Printf("Error retrieving computer diagnosis: %v", err)
		return nil, err
	}
	if dbDiagnosis == nil {
		log.Println("No computer diagnosis found with DiagnosisID:", id)
		return nil, nil
	}
	diagnosis = *dto.MapComputerDiagnosticToDTO(dbDiagnosis)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, diagnosis); err != nil {
		log.Printf("Failed to cache computer diagnosis: %v", err)
	}

	return &diagnosis, nil
}

func (s *computerDiagnosticService) GetAllComputerDiagnostics() ([]*dto.ComputerDiagnosticDTO, error) {
	ctx := context.Background()
	cacheKey := "computer_diagnostics:all"

	// Attempt to fetch from cache
	var diagnostics []*dto.ComputerDiagnosticDTO
	found, err := s.cache.Get(ctx, cacheKey, &diagnostics)
	if err != nil {
		log.Printf("Error accessing cache for all computer diagnostics: %v", err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for all computer diagnostics")
		return diagnostics, nil
	}

	log.Println("Fetching all computer diagnostics")
	dbDiagnostics, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error retrieving computer diagnostics: %v", err)
		return nil, err
	}
	diagnostics = dto.MapComputerDiagnosticsToDTOs(dbDiagnostics)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, diagnostics); err != nil {
		log.Printf("Failed to cache computer diagnostics: %v", err)
	}

	return diagnostics, nil
}

func (s *computerDiagnosticService) UpdateComputerDiagnostic(id uuid.UUID, diagnosisDTO *dto.ComputerDiagnosticUpdateDTO) error {
	log.Println("Updating computer diagnosis with DiagnosisID:", id)
	diagnosis, err := s.repo.GetByID(id, "diagnostic_id")
	if err != nil {
		log.Printf("Error retrieving computer diagnosis: %v", err)
		return err
	}
	if diagnosis == nil {
		log.Printf("Computer diagnosis not found with DiagnosisID: %v", id)
		return gorm.ErrRecordNotFound
	}

	diagnosis.Diagnosis = diagnosisDTO.Diagnosis
	diagnosis.Percentage = diagnosisDTO.Percentage

	err = s.repo.Update(diagnosis, "diagnostic_id", id)
	if err != nil {
		log.Printf("Failed to update computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis updated successfully with DiagnosisID:", diagnosis.DiagnosticID)

	// Invalidate cache for the updated diagnosis and all diagnostics
	_ = s.cache.Delete(context.Background(), "computer_diagnostic:"+id.String(), "computer_diagnostics:all")
	return nil
}

func (s *computerDiagnosticService) DeleteComputerDiagnostic(id uuid.UUID) error {
	log.Println("Deleting computer diagnosis with DiagnosisID:", id)
	err := s.repo.Delete(id, "diagnostic_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Computer diagnosis not found with DiagnosisID:", id)
			return nil
		}
		log.Printf("Failed to delete computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis deleted successfully with DiagnosisID:", id)

	// Invalidate cache for the deleted diagnosis and all diagnostics
	_ = s.cache.Delete(context.Background(), "computer_diagnostic:"+id.String(), "computer_diagnostics:all")
	return nil
}
