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

type BiometricDataService interface {
	CreateBiometricData(biometricDTO *dto.BiometricDataCreateDTO) error
	GetBiometricDataByID(id uuid.UUID) (*dto.BiometricDataDTO, error)
	GetAllBiometricData() ([]*dto.BiometricDataDTO, error)
	UpdateBiometricData(id uuid.UUID, biometricDTO *dto.BiometricDataUpdateDTO) error
	DeleteBiometricData(id uuid.UUID) error
}

type biometricService struct {
	repo repository.BiometricDataRepository
}

func NewBiometricDataService(repo repository.BiometricDataRepository) BiometricDataService {
	return &biometricService{repo: repo}
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
	log.Println("BiometricDataData created successfully with BiometricDataID:", biometric.BiometricDataID)
	return nil
}

func (s *biometricService) GetBiometricDataByID(id uuid.UUID) (*dto.BiometricDataDTO, error) {
	log.Println("Fetching biometric with BiometricDataID:", id)
	biometric, err := s.repo.GetByID(id, "biometric_data_id")
	if err != nil {
		log.Printf("Error retrieving biometric: %v", err)
		return nil, err
	}
	if biometric == nil {
		log.Println("No biometric found with BiometricDataID:", id)
		return nil, nil
	}

	biometricDTO := dto.MapBiometricDataToDTO(biometric)
	log.Println("BiometricDataData fetched successfully with BiometricDataID:", id)
	return biometricDTO, nil
}

func (s *biometricService) GetAllBiometricData() ([]*dto.BiometricDataDTO, error) {
	log.Println("Fetching all biometrics")
	biometrics, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error retrieving biometrics: %v", err)
		return nil, err
	}

	biometricDTOs := dto.MapBiometricDataToDTOs(biometrics)
	log.Println("BiometricData fetched successfully, total count:", len(biometricDTOs))
	return biometricDTOs, nil
}

func (s *biometricService) UpdateBiometricData(id uuid.UUID, biometricDTO *dto.BiometricDataUpdateDTO) error {
	log.Println("Updating biometric with  BiometricDataID:", id)

	biometric, err := s.repo.GetByID(id, "biometric_data_id")
	if err != nil {
		log.Printf("Error retrieving biometric: %v", err)
		return err
	}
	if biometric == nil {
		log.Printf("BiometricDataData not found with BiometricDataID: %v", id)
		return gorm.ErrRecordNotFound
	}

	biometric.O2Saturation = biometricDTO.O2Saturation
	biometric.HeartRate = biometricDTO.HeartRate

	err = s.repo.Update(biometric, "biometric_data_id", id)
	if err != nil {
		log.Printf("Failed to update biometric: %v", err)
		return err
	}
	log.Println("BiometricDataData updated successfully with BiometricDataID:", biometric.BiometricDataID)
	return nil
}

func (s *biometricService) DeleteBiometricData(id uuid.UUID) error {
	log.Println("Deleting biometric with BiometricDataID:", id)
	err := s.repo.Delete(id, "biometric_data_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("BiometricDataData not found with BiometricDataID:", id)
			return nil
		}
		log.Printf("Failed to delete biometric: %v", err)
		return err
	}
	log.Println("BiometricDataData deleted successfully with BiometricDataID:", id)
	return nil
}
