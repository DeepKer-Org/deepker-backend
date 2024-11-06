package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"log"
	"time"
)

// DoctorCreateDTO is used for creating a new doctor
type DoctorCreateDTO struct {
	DNI            string   `json:"dni" binding:"required"`
	Password       string   `json:"password" binding:"required,min=12"`
	Name           string   `json:"name"`
	Specialization string   `json:"specialization"`
	Roles          []string `json:"roles" binding:"required"`
	IssuanceDate   string   `json:"issuance_date"`
}

// DoctorUpdateDTO is used for updating an existing doctor
type DoctorUpdateDTO struct {
	DNI            string   `json:"dni"`
	Password       string   `json:"password"`
	Name           string   `json:"name"`
	Specialization string   `json:"specialization"`
	Roles          []string `json:"roles"`
	IssuanceDate   string   `json:"issuance_date"`
}

// DoctorDTO is used for retrieving a doctor
type DoctorDTO struct {
	DoctorID       uuid.UUID `json:"doctor_id"`
	DNI            string    `json:"dni"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
	IssuanceDate   string    `json:"issuance_date"`
}

// MapDoctorToDTO maps a Doctor model to a DoctorDTO
func MapDoctorToDTO(doctor *models.Doctor) *DoctorDTO {
	if doctor == nil {
		return &DoctorDTO{}
	}
	return &DoctorDTO{
		DoctorID:       doctor.DoctorID,
		DNI:            doctor.DNI,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		IssuanceDate:   doctor.IssuanceDate.Format("2006-01-02"),
	}
}

// MapDoctorsToDTOs maps a list of Doctor models to a list of DoctorDTOs
func MapDoctorsToDTOs(doctors []*models.Doctor) []*DoctorDTO {
	doctorDTOs := make([]*DoctorDTO, 0)
	for _, doctor := range doctors {
		doctorDTOs = append(doctorDTOs, MapDoctorToDTO(doctor))
	}
	return doctorDTOs
}

// MapCreateDTOToDoctor maps a DoctorCreateDTO to a Doctor model
func MapCreateDTOToDoctor(dto *DoctorCreateDTO) *models.Doctor {
	// Define the expected date format
	const layout = "2006-01-02"

	// Parse the issuance date string to time.Time
	issuanceDate, err := time.Parse(layout, dto.IssuanceDate)
	if err != nil {
		log.Printf("Error parsing IssuanceDate: %v", err)
		issuanceDate = time.Time{}
	}

	return &models.Doctor{
		DNI:            dto.DNI,
		Name:           dto.Name,
		Specialization: dto.Specialization,
		IssuanceDate:   issuanceDate,
	}
}

// MapUpdateDTOToDoctor maps a DoctorUpdateDTO to a Doctor model
func MapUpdateDTOToDoctor(dto *DoctorUpdateDTO, doctor *models.Doctor) *models.Doctor {
	doctor.DNI = dto.DNI
	doctor.Name = dto.Name
	doctor.Specialization = dto.Specialization
	return doctor
}

// MapDoctorsToNames maps a list of Doctor models to a list of strings (names of doctors)
func MapDoctorsToNames(doctors []*models.Doctor) []string {
	doctorNames := make([]string, 0)
	for _, doctor := range doctors {
		doctorNames = append(doctorNames, doctor.Name)
	}
	return doctorNames
}
