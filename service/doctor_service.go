package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type DoctorService interface {
	CreateDoctor(doctorDTO *dto.DoctorCreateDTO) error
	GetDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error)
	GetDoctorsByAlertID(alertID uuid.UUID) ([]*dto.DoctorDTO, error)
	GetShortDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error)
	GetAllDoctors() ([]*dto.DoctorDTO, error)
	UpdateDoctor(id uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error
	DeleteDoctor(id uuid.UUID) error
}

type doctorService struct {
	repo repository.DoctorRepository
}

func NewDoctorService(repo repository.DoctorRepository) DoctorService {
	return &doctorService{repo: repo}
}

// Create a new doctor using the DoctorCreateDTO
func (s *doctorService) CreateDoctor(doctorDTO *dto.DoctorCreateDTO) error {
	log.Println("Creating a new doctor with DNI:", doctorDTO.DNI)
	// Map the DTO to the Doctor entity
	doctor := dto.MapCreateDTOToDoctor(doctorDTO)

	err := s.repo.Create(doctor)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		return err
	}
	log.Println("Doctor created successfully with DoctorID:", doctor.DoctorID)
	return nil
}

// Get a doctor by ID and map the result to DoctorDTO
func (s *doctorService) GetDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error) {
	log.Println("Fetching doctor with DoctorID:", id)
	doctor, err := s.repo.GetByID(id, "doctor_id")
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return nil, err
	}
	if doctor == nil {
		log.Println("No doctor found with DoctorID:", id)
		return nil, nil
	}

	// Map the Doctor entity to the DoctorDTO
	return dto.MapDoctorToDTO(doctor), nil
}

// Get doctors by AlertID and map to DoctorDTO
func (s *doctorService) GetDoctorsByAlertID(alertID uuid.UUID) ([]*dto.DoctorDTO, error) {
	doctors, err := s.repo.GetDoctorsByAlertID(alertID)
	if err != nil {
		log.Printf("Error fetching doctors: %v", err)
		return nil, err
	}
	log.Println("Doctors fetched successfully by AlertID:", alertID)

	// Map the list of Doctor entities to DoctorDTOs
	return dto.MapDoctorsToDTOs(doctors), nil
}

// Get a short representation of a doctor (DoctorDTO) by ID
func (s *doctorService) GetShortDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error) {
	log.Println("Fetching short version of doctor with DoctorID:", id)
	doctor, err := s.repo.GetByID(id, "doctor_id")
	if err != nil {
		return nil, err
	}
	if doctor == nil {
		return nil, nil
	}

	// Map to DoctorDTO
	return dto.MapDoctorToDTO(doctor), nil
}

// Get all doctors and map them to DoctorDTO
func (s *doctorService) GetAllDoctors() ([]*dto.DoctorDTO, error) {
	log.Println("Fetching all doctors")
	doctors, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching doctors: %v", err)
		return nil, err
	}
	log.Println("Doctors fetched successfully, total count:", len(doctors))

	// Map the list of Doctor entities to DoctorDTOs
	return dto.MapDoctorsToDTOs(doctors), nil
}

// Update an existing doctor using DoctorUpdateDTO
func (s *doctorService) UpdateDoctor(id uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error {
	log.Println("Updating doctor with DoctorID:", id)

	// Fetch the existing doctor entity
	doctor, err := s.repo.GetByID(id, "doctor_id")
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return err
	}
	if doctor == nil {
		log.Printf("Doctor not found with DoctorID: %v", id)
		return gorm.ErrRecordNotFound
	}

	// Map the UpdateDTO to the existing entity
	doctor = dto.MapUpdateDTOToDoctor(doctorDTO, doctor)

	// Update the doctor in the database
	err = s.repo.Update(doctor, "doctor_id", id)
	if err != nil {
		log.Printf("Failed to update doctor: %v", err)
		return err
	}
	log.Println("Doctor updated successfully with DoctorID:", doctor.DoctorID)
	return nil
}

// Delete a doctor by ID
func (s *doctorService) DeleteDoctor(id uuid.UUID) error {
	log.Println("Deleting doctor with DoctorID:", id)
	err := s.repo.Delete(id, "doctor_id")
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
