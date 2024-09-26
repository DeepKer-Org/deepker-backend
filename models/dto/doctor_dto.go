package dto

import "biometric-data-backend/models"

type DoctorDTO struct {
	DNI            string `json:"dni"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}

// MapDoctorToDTO maps a Doctor model to a DoctorDTO
func MapDoctorToDTO(doctor *models.Doctor) *DoctorDTO {
	return &DoctorDTO{
		DNI:            doctor.DNI,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
	}
}

// MapDoctorsToDTOs maps a list of Doctor models to a list of DoctorDTOs
func MapDoctorsToDTOs(doctors []*models.Doctor) []*DoctorDTO {
	var doctorDTOs []*DoctorDTO
	for _, doctor := range doctors {
		doctorDTOs = append(doctorDTOs, MapDoctorToDTO(doctor))
	}
	return doctorDTOs
}
