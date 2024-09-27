package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ComputerDiagnosticService interface {
	CreateComputerDiagnostic(diagnosisDTO *dto.ComputerDiagnosticCreateDTO) error
	GetComputerDiagnosticByID(id uint) (*dto.ComputerDiagnosticDTO, error)
	GetAllComputerDiagnostics() ([]*dto.ComputerDiagnosticDTO, error)
	UpdateComputerDiagnostic(id uint, diagnosisDTO *dto.ComputerDiagnosticUpdateDTO) error
	DeleteComputerDiagnostic(id uint) error
}

type computerDiagnosticService struct {
	repo repository.ComputerDiagnosticRepository
}

func NewComputerDiagnosticService(repo repository.ComputerDiagnosticRepository) ComputerDiagnosticService {
	return &computerDiagnosticService{repo: repo}
}

func (s *computerDiagnosticService) CreateComputerDiagnostic(diagnosisDTO *dto.ComputerDiagnosticCreateDTO) error {
	diagnosis := &models.ComputerDiagnostic{
		AlertID:    diagnosisDTO.AlertID,
		Diagnosis:  diagnosisDTO.Diagnosis,
		Percentage: diagnosisDTO.Percentage,
	}

	err := s.repo.CreateComputerDiagnostic(diagnosis)
	if err != nil {
		log.Printf("Failed to create computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis created successfully with DiagnosisID:", diagnosis.DiagnosisID)
	return nil
}

func (s *computerDiagnosticService) GetComputerDiagnosticByID(id uint) (*dto.ComputerDiagnosticDTO, error) {
	log.Println("Fetching computer diagnosis with DiagnosisID:", id)
	diagnosis, err := s.repo.GetComputerDiagnosticByID(id)
	if err != nil {
		log.Printf("Error retrieving computer diagnosis: %v", err)
		return nil, err
	}
	if diagnosis == nil {
		log.Println("No computer diagnosis found with DiagnosisID:", id)
		return nil, nil
	}

	diagnosisDTO := dto.MapComputerDiagnosticToDTO(diagnosis)
	log.Println("Computer diagnosis fetched successfully with DiagnosisID:", id)
	return diagnosisDTO, nil
}

func (s *computerDiagnosticService) GetAllComputerDiagnostics() ([]*dto.ComputerDiagnosticDTO, error) {
	log.Println("Fetching all computer diagnostics")
	diagnostics, err := s.repo.GetAllComputerDiagnostics()
	if err != nil {
		log.Printf("Error retrieving computer diagnostics: %v", err)
		return nil, err
	}

	diagnosisDTOs := dto.MapComputerDiagnosticsToDTOs(diagnostics)
	log.Println("Computer diagnostics fetched successfully, total count:", len(diagnosisDTOs))
	return diagnosisDTOs, nil
}

func (s *computerDiagnosticService) UpdateComputerDiagnostic(id uint, diagnosisDTO *dto.ComputerDiagnosticUpdateDTO) error {
	log.Println("Updating computer diagnosis with DiagnosisID:", id)

	diagnosis, err := s.repo.GetComputerDiagnosticByID(id)
	if err != nil {
		log.Printf("Error retrieving computer diagnosis: %v", err)
		return err
	}
	if diagnosis == nil {
		log.Printf("Computer diagnosis not found with DiagnosisID: %v", id)
		return gorm.ErrRecordNotFound
	}

	diagnosis.AlertID = diagnosisDTO.AlertID
	diagnosis.Diagnosis = diagnosisDTO.Diagnosis
	diagnosis.Percentage = diagnosisDTO.Percentage

	err = s.repo.UpdateComputerDiagnostic(diagnosis)
	if err != nil {
		log.Printf("Failed to update computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis updated successfully with DiagnosisID:", diagnosis.DiagnosisID)
	return nil
}

func (s *computerDiagnosticService) DeleteComputerDiagnostic(id uint) error {
	log.Println("Deleting computer diagnosis with DiagnosisID:", id)
	err := s.repo.DeleteComputerDiagnostic(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Computer diagnosis not found with DiagnosisID:", id)
			return nil
		}
		log.Printf("Failed to delete computer diagnosis: %v", err)
		return err
	}
	log.Println("Computer diagnosis deleted successfully with DiagnosisID:", id)
	return nil
}
