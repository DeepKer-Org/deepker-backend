package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// ComorbidityCreateDTO is used for the creation of a new comorbidity
type ComorbidityCreateDTO struct {
	PatientID   uuid.UUID `json:"patient_id"`
	Comorbidity string    `json:"comorbidity"`
}

// ComorbidityUpdateDTO is used for updating an existing comorbidity
type ComorbidityUpdateDTO struct {
	PatientID   uuid.UUID `json:"patient_id"`
	Comorbidity string    `json:"comorbidity"`
}

// ComorbidityDTO is used for retrieving a comorbidity
type ComorbidityDTO struct {
	ComorbidityID uuid.UUID `json:"comorbidity_id"`
	PatientID     uuid.UUID `json:"patient_id"`
	Comorbidity   string    `json:"comorbidity"`
}

// MapComorbidityToDTO maps a Comorbidity model to a ComorbidityDTO
func MapComorbidityToDTO(comorbidity *models.Comorbidity) *ComorbidityDTO {
	return &ComorbidityDTO{
		ComorbidityID: comorbidity.ComorbidityID,
		PatientID:     comorbidity.PatientID,
		Comorbidity:   comorbidity.Comorbidity,
	}
}

// MapComorbiditiesToDTOs maps a list of Comorbidity models to a list of ComorbidityDTOs
func MapComorbiditiesToDTOs(comorbidities []*models.Comorbidity) []*ComorbidityDTO {
	var comorbidityDTOs []*ComorbidityDTO
	for _, comorbidity := range comorbidities {
		comorbidityDTOs = append(comorbidityDTOs, MapComorbidityToDTO(comorbidity))
	}
	return comorbidityDTOs
}

// MapCreateDTOToComorbidity maps a ComorbidityCreateDTO to a Comorbidity model
func MapCreateDTOToComorbidity(dto *ComorbidityCreateDTO) *models.Comorbidity {
	return &models.Comorbidity{
		PatientID:   dto.PatientID,
		Comorbidity: dto.Comorbidity,
	}
}

// MapUpdateDTOToComorbidity maps a ComorbidityUpdateDTO to a Comorbidity model
func MapUpdateDTOToComorbidity(dto *ComorbidityUpdateDTO, comorbidity *models.Comorbidity) *models.Comorbidity {
	comorbidity.PatientID = dto.PatientID
	comorbidity.Comorbidity = dto.Comorbidity
	return comorbidity
}

// MapComorbiditiesToNames maps a list of Comorbidity models to a list of strings (names of comorbidities)
func MapComorbiditiesToNames(comorbidities []*models.Comorbidity) []string {
	comorbidityNames := make([]string, 0)
	for _, comorbidity := range comorbidities {
		comorbidityNames = append(comorbidityNames, comorbidity.Comorbidity)
	}
	return comorbidityNames
}
