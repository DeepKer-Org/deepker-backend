package comorbidity

import (
	"biometric-data-backend/internal/comorbidity/dto"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Service interface {
	CreateComorbidity(comorbidityDTO *dto.CreateDTO) error
	GetComorbidityByID(id uuid.UUID) (*dto.ResponseDTO, error)
	GetAllComorbidities() ([]*dto.ResponseDTO, error)
	UpdateComorbidity(id uuid.UUID, comorbidityDTO *dto.UpdateDTO) error
	DeleteComorbidity(id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateComorbidity(comorbidityDTO *dto.CreateDTO) error {
	comorbidity := MapCreateDTOToComorbidity(comorbidityDTO)
	err := s.repo.Create(comorbidity)
	if err != nil {
		log.Printf("Failed to create comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity created successfully with ComorbidityID:", comorbidity.ComorbidityID)
	return nil
}

func (s *service) GetComorbidityByID(id uuid.UUID) (*dto.ResponseDTO, error) {
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

	return MapComorbidityToDTO(comorbidity), nil
}

func (s *service) GetAllComorbidities() ([]*dto.ResponseDTO, error) {
	log.Println("Fetching all comorbidities")
	comorbidities, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching comorbidities: %v", err)
		return nil, err
	}

	return MapComorbiditiesToDTOs(comorbidities), nil
}

func (s *service) UpdateComorbidity(id uuid.UUID, comorbidityDTO *dto.UpdateDTO) error {
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

	comorbidity = MapUpdateDTOToComorbidity(comorbidityDTO, comorbidity)
	err = s.repo.Update(comorbidity, "comorbidity_id", id)
	if err != nil {
		log.Printf("Failed to update comorbidity: %v", err)
		return err
	}
	log.Println("Comorbidity updated successfully with ComorbidityID:", comorbidity.ComorbidityID)
	return nil
}

func (s *service) DeleteComorbidity(id uuid.UUID) error {
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
