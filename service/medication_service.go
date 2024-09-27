package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"gorm.io/gorm"
	"log"
)

type MedicationService interface {
	CreateMedication(medicationDTO *dto.MedicationCreateDTO) error
	GetMedicationByID(id uint) (*dto.MedicationDTO, error)
	GetAllMedications() ([]*dto.MedicationDTO, error)
	UpdateMedication(id uint, medicationDTO *dto.MedicationUpdateDTO) error
	DeleteMedication(id uint) error
}

type medicationService struct {
	repo repository.MedicationRepository
}

func NewMedicationService(repo repository.MedicationRepository) MedicationService {
	return &medicationService{repo: repo}
}

func (s *medicationService) CreateMedication(medicationDTO *dto.MedicationCreateDTO) error {
	medication := &models.Medication{
		PatientID:   medicationDTO.PatientID,
		Medication:  medicationDTO.Medication,
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
	return nil
}

func (s *medicationService) GetMedicationByID(id uint) (*dto.MedicationDTO, error) {
	log.Println("Fetching medication with MedicationID:", id)
	medication, err := s.repo.GetMedicationByID(id)
	if err != nil {
		log.Printf("Error retrieving medication: %v", err)
		return nil, err
	}
	if medication == nil {
		log.Println("No medication found with MedicationID:", id)
		return nil, nil
	}

	medicationDTO := dto.MapMedicationToDTO(medication)
	log.Println("Medication fetched successfully with MedicationID:", id)
	return medicationDTO, nil
}

func (s *medicationService) GetAllMedications() ([]*dto.MedicationDTO, error) {
	log.Println("Fetching all medications")
	medications, err := s.repo.GetAllMedications()
	if err != nil {
		log.Printf("Error retrieving medications: %v", err)
		return nil, err
	}

	medicationDTOs := dto.MapMedicationsToDTOs(medications)
	log.Println("Medications fetched successfully, total count:", len(medicationDTOs))
	return medicationDTOs, nil
}

func (s *medicationService) UpdateMedication(id uint, medicationDTO *dto.MedicationUpdateDTO) error {
	log.Println("Updating medication with MedicationID:", id)

	medication, err := s.repo.GetMedicationByID(id)
	if err != nil {
		log.Printf("Error retrieving medication: %v", err)
		return err
	}
	if medication == nil {
		log.Printf("Medication not found with MedicationID: %v", id)
		return gorm.ErrRecordNotFound
	}

	medication.PatientID = medicationDTO.PatientID
	medication.Medication = medicationDTO.Medication
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
	return nil
}

func (s *medicationService) DeleteMedication(id uint) error {
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
	return nil
}
