package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"gorm.io/gorm"
	"log"
)

type BiometricService interface {
	CreateBiometric(biometricDTO *dto.BiometricCreateDTO) error
	GetBiometricByID(id string) (*dto.BiometricDTO, error)
	GetAllBiometrics() ([]*dto.BiometricDTO, error)
	UpdateBiometric(id string, biometricDTO *dto.BiometricUpdateDTO) error
	DeleteBiometric(id string) error
}

type biometricService struct {
	repo repository.BiometricRepository
}

func NewBiometricService(repo repository.BiometricRepository) BiometricService {
	return &biometricService{repo: repo}
}

func (s *biometricService) CreateBiometric(biometricDTO *dto.BiometricCreateDTO) error {
	biometric := &models.Biometric{
		AlertID:                biometricDTO.AlertID,
		O2Saturation:           biometricDTO.O2Saturation,
		HeartRate:              biometricDTO.HeartRate,
		SystolicBloodPressure:  biometricDTO.SystolicBloodPressure,
		DiastolicBloodPressure: biometricDTO.DiastolicBloodPressure,
	}

	err := s.repo.CreateBiometric(biometric)
	if err != nil {
		log.Printf("Failed to create biometric: %v", err)
		return err
	}
	log.Println("Biometric created successfully with BiometricsID:", biometric.BiometricsID)
	return nil
}

func (s *biometricService) GetBiometricByID(id string) (*dto.BiometricDTO, error) {
	log.Println("Fetching biometric with BiometricsID:", id)
	biometric, err := s.repo.GetBiometricByID(id)
	if err != nil {
		log.Printf("Error retrieving biometric: %v", err)
		return nil, err
	}
	if biometric == nil {
		log.Println("No biometric found with BiometricsID:", id)
		return nil, nil
	}

	biometricDTO := dto.MapBiometricToDTO(biometric)
	log.Println("Biometric fetched successfully with BiometricsID:", id)
	return biometricDTO, nil
}

func (s *biometricService) GetAllBiometrics() ([]*dto.BiometricDTO, error) {
	log.Println("Fetching all biometrics")
	biometrics, err := s.repo.GetAllBiometrics()
	if err != nil {
		log.Printf("Error retrieving biometrics: %v", err)
		return nil, err
	}

	biometricDTOs := dto.MapBiometricsToDTOs(biometrics)
	log.Println("Biometrics fetched successfully, total count:", len(biometricDTOs))
	return biometricDTOs, nil
}

func (s *biometricService) UpdateBiometric(id string, biometricDTO *dto.BiometricUpdateDTO) error {
	log.Println("Updating biometric with BiometricsID:", id)

	biometric, err := s.repo.GetBiometricByID(id)
	if err != nil {
		log.Printf("Error retrieving biometric: %v", err)
		return err
	}
	if biometric == nil {
		log.Printf("Biometric not found with BiometricsID: %v", id)
		return gorm.ErrRecordNotFound
	}

	biometric.AlertID = biometricDTO.AlertID
	biometric.O2Saturation = biometricDTO.O2Saturation
	biometric.HeartRate = biometricDTO.HeartRate
	biometric.SystolicBloodPressure = biometricDTO.SystolicBloodPressure
	biometric.DiastolicBloodPressure = biometricDTO.DiastolicBloodPressure

	err = s.repo.UpdateBiometric(biometric)
	if err != nil {
		log.Printf("Failed to update biometric: %v", err)
		return err
	}
	log.Println("Biometric updated successfully with BiometricsID:", biometric.BiometricsID)
	return nil
}

func (s *biometricService) DeleteBiometric(id string) error {
	log.Println("Deleting biometric with BiometricsID:", id)
	err := s.repo.DeleteBiometric(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Biometric not found with BiometricsID:", id)
			return nil
		}
		log.Printf("Failed to delete biometric: %v", err)
		return err
	}
	log.Println("Biometric deleted successfully with BiometricsID:", id)
	return nil
}
