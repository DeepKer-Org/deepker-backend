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

type DoctorService interface {
	CreateDoctor(doctor *models.Doctor) error
	GetDoctorByID(id uuid.UUID) (*models.Doctor, error)
	GetShortDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error)
	GetAllDoctors() ([]*models.Doctor, error)
	UpdateDoctor(doctor *models.Doctor) error
	DeleteDoctor(id uuid.UUID) error
}

type doctorService struct {
	repo repository.DoctorRepository
}

func NewDoctorService(repo repository.DoctorRepository) DoctorService {
	return &doctorService{repo: repo}
}

func (s *doctorService) CreateDoctor(doctor *models.Doctor) error {
	log.Println("Creating a new doctor with DNI:", doctor.DNI)
	err := s.repo.CreateDoctor(doctor)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		return err
	}
	log.Println("Doctor created successfully with DoctorID:", doctor.DoctorID)
	return nil
}

func (s *doctorService) GetDoctorByID(id uuid.UUID) (*models.Doctor, error) {
	log.Println("Fetching doctor with DoctorID:", id)
	doctor, err := s.repo.GetDoctorByID(id)
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return nil, err
	}
	if doctor == nil {
		log.Println("No doctor found with DoctorID:", id)
		return nil, nil
	}
	log.Println("Doctor fetched successfully with DoctorID:", id)
	return doctor, nil
}

func (s *doctorService) GetShortDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error) {
	doctor, err := s.repo.GetDoctorByID(id)
	if err != nil {
		return nil, err
	}
	return dto.MapDoctorToDTO(doctor), nil
}

func (s *doctorService) GetAllDoctors() ([]*models.Doctor, error) {
	log.Println("Fetching all doctors")
	doctors, err := s.repo.GetAllDoctors()
	if err != nil {
		log.Printf("Error fetching doctors: %v", err)
		return nil, err
	}
	log.Println("Doctors fetched successfully, total count:", len(doctors))
	return doctors, nil
}

func (s *doctorService) UpdateDoctor(doctor *models.Doctor) error {
	log.Println("Updating doctor with DoctorID:", doctor.DoctorID)
	err := s.repo.UpdateDoctor(doctor)
	if err != nil {
		log.Printf("Failed to update doctor: %v", err)
		return err
	}
	log.Println("Doctor updated successfully with DoctorID:", doctor.DoctorID)
	return nil
}

func (s *doctorService) DeleteDoctor(id uuid.UUID) error {
	log.Println("Deleting doctor with DoctorID:", id)
	err := s.repo.DeleteDoctor(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Doctor not found with DoctorID:", id)
			return nil
		}
		log.Printf("Failed to delete doctor: %v", err)
		return err
	}
	log.Println("Doctor deleted successfully with DoctorID:", id)
	return nil
}
