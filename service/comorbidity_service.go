package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
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
	repo repository.ComorbidityRepository
}

func NewComorbidityService(repo repository.ComorbidityRepository) ComorbidityService {
	return &comorbidityService{repo: repo}
}

func (s *comorbidityService) CreateComorbidity(comorbidityDTO *dto.ComorbidityCreateDTO) error {
	comorbidity := dto.MapCreateDTOToComorbidity(comorbidityDTO)
	err := s.repo.Create(comorbidity)
	if err != nil {
		log.Printf("Failed to create comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity created successfully with ComorbidityID:", comorbidity.ComorbidityID)
	return nil
}

func (s *comorbidityService) GetComorbidityByID(id uuid.UUID) (*dto.ComorbidityDTO, error) {
	log.Println("Fetching comorbidity with ComorbidityID:", id)
	comorbidity, err := s.repo.GetByID(id, "comorbidity_id")
	if err != nil {
		log.Printf("Error fetching comorbidity: %v", err)
		return nil, err
	}
	if comorbidity == nil {
		log.Println("No comorbidity found with ComorbidityID:", id)
		return nil, nil
	}

	return dto.MapComorbidityToDTO(comorbidity), nil
}

func (s *comorbidityService) GetAllComorbidities() ([]*dto.ComorbidityDTO, error) {
	log.Println("Fetching all comorbidities")
	comorbidities, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching comorbidities: %v", err)
		return nil, err
	}

	return dto.MapComorbiditiesToDTOs(comorbidities), nil
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
	return nil
}
