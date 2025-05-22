package service

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
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
	UpdateDoctorByUserID(userID uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error
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
	roleRepo    repository.RoleRepository
	authService AuthorizationService
	cache       *redis.CacheManager
}

func NewDoctorService(
	repo repository.DoctorRepository,
	authRepo repository.AuthorizationRepository,
	roleRepo repository.RoleRepository,
	authService AuthorizationService,
	cache *redis.CacheManager,
) DoctorService {
	return &doctorService{
		repo:        repo,
		authRepo:    authRepo,
		roleRepo:    roleRepo,
		authService: authService,
		cache:       cache,
	}
}

// CreateDoctor creates a new doctor using the DoctorCreateDTO
func (s *doctorService) CreateDoctor(doctorDTO *dto.DoctorCreateDTO) error {
	log.Println("Creating a new doctor with DNI:", doctorDTO.DNI)

	tx := s.repo.BeginTransaction()

	user := &dto.UserRegisterDTO{
		Username: doctorDTO.DNI,
		Password: doctorDTO.Password,
		Roles:    doctorDTO.Roles,
	}

	userID, err := s.authService.RegisterUserInTransaction(user, tx)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		tx.Rollback()
		return err
	}

	doctor := dto.MapCreateDTOToDoctor(doctorDTO)
	doctor.UserID = *userID

	err = s.repo.CreateInTransaction(doctor, tx)
	if err != nil {
		log.Printf("Failed to create doctor: %v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	// Invalidate cache for all doctors
	_ = s.cache.Delete(context.Background(), "doctors:all")

	log.Println("Doctor created successfully with DoctorID:", doctor.DoctorID)
	return nil
}

// GetDoctorByID fetches a doctor by ID and uses cache
func (s *doctorService) GetDoctorByID(id uuid.UUID) (*dto.DoctorDTO, error) {
	ctx := context.Background()
	cacheKey := "doctor:" + id.String()

	// Attempt to fetch from cache
	var doctor dto.DoctorDTO
	found, err := s.cache.Get(ctx, cacheKey, &doctor)
	if err != nil {
		log.Printf("Error accessing cache for DoctorID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for doctor with DoctorID:", id)
		return &doctor, nil
	}

	log.Println("Fetching doctor with DoctorID:", id)
	dbDoctor, err := s.repo.GetByID(id, "doctor_id")
	if err != nil {
		return nil, err
	}
	if dbDoctor == nil {
		return nil, nil
	}

	doctor = *dto.MapDoctorToDTO(dbDoctor)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, doctor); err != nil {
		log.Printf("Failed to cache doctor: %v", err)
	}

	return &doctor, nil
}

// GetAllDoctors fetches all doctors with caching
func (s *doctorService) GetAllDoctors() ([]*dto.DoctorDTO, error) {
	ctx := context.Background()
	cacheKey := "doctors:all"

	// Attempt to fetch from cache
	var doctors []*dto.DoctorDTO
	found, err := s.cache.Get(ctx, cacheKey, &doctors)
	if err != nil {
		log.Printf("Error accessing cache for all doctors: %v", err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for all doctors")
		return doctors, nil
	}

	log.Println("Fetching all doctors")
	dbDoctors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	doctors = dto.MapDoctorsToDTOs(dbDoctors)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, doctors); err != nil {
		log.Printf("Failed to cache doctors: %v", err)
	}

	return doctors, nil
}

// UpdateDoctor updates a doctor and invalidates the cache
func (s *doctorService) UpdateDoctor(id uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error {
	log.Println("Updating doctor with DoctorID:", id)

	doctor, err := s.repo.GetByID(id, "doctor_id")
	if err != nil {
		log.Printf("Error fetching doctor: %v", err)
		return err
	}
	if doctor == nil {
		return gorm.ErrRecordNotFound
	}

	doctor = dto.MapUpdateDTOToDoctor(doctorDTO, doctor)

	err = s.repo.Update(doctor, "doctor_id", id)
	if err != nil {
		log.Printf("Failed to update doctor: %v", err)
		return err
	}

	// Invalidate cache for the updated doctor and all doctors
	_ = s.cache.Delete(context.Background(), "doctor:"+id.String(), "doctors:all")

	log.Println("Doctor updated successfully with DoctorID:", doctor.DoctorID)
	return nil
}

// DeleteDoctor deletes a doctor and invalidates the cache
func (s *doctorService) DeleteDoctor(id uuid.UUID) error {
	log.Println("Deleting doctor with DoctorID:", id)

	err := s.repo.Delete(id, "doctor_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	// Invalidate cache for the deleted doctor and all doctors
	_ = s.cache.Delete(context.Background(), "doctor:"+id.String(), "doctors:all")

	log.Println("Doctor deleted successfully with DoctorID:", id)
	return nil
}

// GetDoctorByUserID fetches a doctor by UserID
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

// UpdateDoctorByUserID updates a doctor by UserID
func (s *doctorService) UpdateDoctorByUserID(userID uuid.UUID, doctorDTO *dto.DoctorUpdateDTO) error {
	log.Println("Updating doctor with UserID:", userID)

	// Start a transaction
	tx := s.repo.BeginTransaction()
	if tx == nil {
		log.Println("Failed to start transaction")
		return errors.New("transaction start failed")
	}

	// Retrieve the doctor by UserID
	doctor, err := s.repo.GetDoctorByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving doctor: %v", err)
		tx.Rollback()
		return err
	}
	if doctor == nil {
		log.Printf("Doctor not found with UserID: %v", userID)
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// Update doctor fields
	const layout = "2006-01-02"
	issuanceDate, err := time.Parse(layout, doctorDTO.IssuanceDate)
	if err != nil {
		log.Printf("Error parsing IssuanceDate: %v", err)
		// Handle error, perhaps by setting a default value or returning nil
		issuanceDate = time.Time{} // Zero value if parsing fails
	}

	doctor.Name = doctorDTO.Name
	doctor.Specialization = doctorDTO.Specialization
	doctor.DNI = doctorDTO.DNI
	doctor.IssuanceDate = issuanceDate

	// Update the doctor entity in the transaction
	if err := tx.Save(&doctor).Error; err != nil {
		log.Printf("Failed to update doctor: %v", err)
		tx.Rollback()
		return err
	}

	// Retrieve the related user entity
	user, err := s.authRepo.GetByID(userID, "user_id")
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		tx.Rollback()
		return err
	}
	if user == nil {
		log.Printf("User not found with UserID: %v", userID)
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// Update user details
	user.Username = doctorDTO.DNI
	if doctorDTO.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doctorDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password: %v", err)
			tx.Rollback()
			return err
		}
		user.Password = string(hashedPassword)
	}

	// Update roles if provided
	if len(doctorDTO.Roles) > 0 {
		roles, err := s.roleRepo.GetRolesByNames(doctorDTO.Roles)
		if err != nil {
			log.Printf("Failed to retrieve roles: %v", err)
			tx.Rollback()
			return err
		}

		// Update user roles within the transaction
		if err := s.authRepo.UpdateUserRoles(user, roles); err != nil {
			log.Printf("Failed to update user roles: %v", err)
			tx.Rollback()
			return err
		}
	}

	// Save the updated user entity
	if err := tx.Save(&user).Error; err != nil {
		log.Printf("Failed to update user: %v", err)
		tx.Rollback()
		return err
	}

	// Commit the transaction if all updates were successful
	if err := tx.Commit().Error; err != nil {
		log.Printf("Transaction commit failed: %v", err)
		return err
	}

	log.Println("Doctor and associated user updated successfully with UserID:", userID)
	return nil
}

// GetDoctorsByAlertID Get doctors by AlertID and map to DoctorDTO
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

// GetShortDoctorByID Get a short representation of a doctor (DoctorDTO) by ID
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
