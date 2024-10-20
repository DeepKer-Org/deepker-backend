package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// DoctorCreateDTO is used for creating a new doctor
type DoctorCreateDTO struct {
	DNI            string `json:"dni"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Specialization string `json:"specialization"`
}

// DoctorUpdateDTO is used for updating an existing doctor
type DoctorUpdateDTO struct {
	DNI            string `json:"dni"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Specialization string `json:"specialization"`
}

// DoctorDTO is used for retrieving a doctor
type DoctorDTO struct {
	DoctorID       uuid.UUID `json:"doctor_id"`
	DNI            string    `json:"dni"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
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
	return &models.Doctor{
		DNI:            dto.DNI,
		Name:           dto.Name,
		Password:       dto.Password,
		Specialization: dto.Specialization,
	}
}

// MapUpdateDTOToDoctor maps a DoctorUpdateDTO to a Doctor model
func MapUpdateDTOToDoctor(dto *DoctorUpdateDTO, doctor *models.Doctor) *models.Doctor {
	doctor.DNI = dto.DNI
	doctor.Name = dto.Name
	doctor.Password = dto.Password
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
