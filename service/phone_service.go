package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"log"
)

type PhoneService interface {
	CreatePhone(phoneDTO *dto.PhoneCreateDTO) error
}

type phoneService struct {
	phoneRepo repository.PhoneRepository
}

func NewPhoneService(phoneRepo repository.PhoneRepository) PhoneService {
	return &phoneService{
		phoneRepo: phoneRepo,
	}
}

func (s *phoneService) CreatePhone(phoneDTO *dto.PhoneCreateDTO) error {
	phone := dto.MapCreateDTOToPhone(phoneDTO)
	err := s.phoneRepo.Create(phone)
	if err != nil {
		log.Printf("Failed to create phone: %v", err)
		return err
	}
	log.Println("Patient created successfully with PhoneID:", phone.PhoneID)
	return nil
}
