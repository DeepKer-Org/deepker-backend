package service

import (
	"biometric-data-backend/enums"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type DoctorService interface {
	CreateDoctor(doctorDTO *dto.DoctorCreateDTO) error
	GetDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error)
	GetDoctorByUserID(userID uuid.UUID) (*dto.DoctorDTO, error)
	GetDoctorsByAlertID(alertID uuid.UUID) ([]*dto.DoctorDTO, error)
	GetShortDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error)
	GetAllDoctors() ([]*dto.DoctorDTO, error)
	UpdateDoctor(id uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error
	DeleteDoctor(id uuid.UUID) error
	ChangePassword(changePasswordDTO *dto.ChangePasswordDTO) error
}

type doctorService struct {
	repo        repository.DoctorRepository
	authRepo    repository.AuthorizationRepository
	authService AuthorizationService
}

func NewDoctorService(repo repository.DoctorRepository, authRepo repository.AuthorizationRepository, authService AuthorizationService) DoctorService {
	return &doctorService{
		repo:        repo,
		authRepo:    authRepo,
		authService: authService,
	}
}

// CreateDoctor creates a new doctor using the DoctorCreateDTO
func (s *doctorService) CreateDoctor(doctorDTO *dto.DoctorCreateDTO) error {
	log.Println("Creating a new doctor with DNI:", doctorDTO.DNI)

	// Start a new transaction
	tx := s.repo.BeginTransaction()

	// Create the user inside the transaction
	userRegisterDTO := &dto.UserRegisterDTO{
		Username: doctorDTO.DNI,
		Password: doctorDTO.Password,
		Roles:    enums.ToStringArray(enums.Doctor),
	}
	userID, err := s.authService.RegisterUserInTransaction(userRegisterDTO, tx)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		tx.Rollback() // Rollback the transaction if an error occurs
		return err
	}

	// Map the DTO to the Doctor entity
	doctor := dto.MapCreateDTOToDoctor(doctorDTO)
	doctor.UserID = *userID // Assign the UserID to the Doctor

	// Create the doctor inside the transaction
	err = s.repo.CreateInTransaction(doctor, tx)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		tx.Rollback() // Rollback the transaction if an error occurs
		return err
	}

	// Commit the transaction if everything is successful
	tx.Commit()

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

// Get a doctor by UserID and map the result to DoctorDTO
func (s *doctorService) GetDoctorByUserID(userID uuid.UUID) (*dto.DoctorDTO, error) {
	log.Println("Fetching doctor with UserID:", userID)
	doctor, err := s.repo.GetDoctorByUserID(userID)
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return nil, err
	}
	if doctor == nil {
		log.Println("No doctor found with UserID:", userID)
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

// ChangePassword handles user password change
func (s *doctorService) ChangePassword(changePasswordDTO *dto.ChangePasswordDTO) error {
	log.Println("Changing password for user with DNI:", changePasswordDTO.DNI)

	// Fetch the doctor
	doctor, err := s.repo.GetDoctorByDNI(changePasswordDTO.DNI)
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return err
	}
	if doctor == nil {
		log.Printf("Doctor not found with DNI: %v", changePasswordDTO.DNI)
		return gorm.ErrRecordNotFound
	}

	parseDate, err := time.Parse("2006-01-02", changePasswordDTO.IssuanceDate)
	if !doctor.IssuanceDate.Equal(parseDate) {
		log.Println("Incorrect issuance date for user with DNI:", changePasswordDTO.DNI)
		return errors.New("incorrect issuance date")
	}

	// Fetch the user by ID
	user, err := s.authRepo.GetByID(doctor.User.UserID, "user_id")
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return err
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordDTO.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash new password: %v", err)
		return err
	}

	// Update the user's password
	user.Password = string(hashedPassword)
	err = s.authRepo.Update(user, "user_id", user.UserID)
	if err != nil {
		log.Printf("Failed to update password for user ID %v: %v", user.UserID, err)
		return err
	}

	log.Println("Password changed successfully for user ID:", user.UserID)
	return nil
}
