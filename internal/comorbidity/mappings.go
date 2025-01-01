package comorbidity

import "biometric-data-backend/internal/comorbidity/dto"

// MapComorbidityToDTO maps a Comorbidity model to a ResponseDTO
func MapComorbidityToDTO(comorbidity *Model) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		ComorbidityID: comorbidity.ComorbidityID,
		PatientID:     comorbidity.PatientID,
		Comorbidity:   comorbidity.Comorbidity,
	}
}

// MapComorbiditiesToDTOs maps a list of Comorbidity models to a list of ResponseDTOs
func MapComorbiditiesToDTOs(comorbidities []*Model) []*dto.ResponseDTO {
	var comorbidityDTOs []*dto.ResponseDTO
	for _, model := range comorbidities {
		comorbidityDTOs = append(comorbidityDTOs, MapComorbidityToDTO(model))
	}
	return comorbidityDTOs
}

// MapCreateDTOToComorbidity maps a CreateDTO to a Comorbidity model
func MapCreateDTOToComorbidity(dto *dto.CreateDTO) *Model {
	return &Model{
		PatientID:   dto.PatientID,
		Comorbidity: dto.Comorbidity,
	}
}

// MapUpdateDTOToComorbidity maps a UpdateDTO to a Comorbidity model
func MapUpdateDTOToComorbidity(dto *dto.UpdateDTO, comorbidity *Model) *Model {
	comorbidity.PatientID = dto.PatientID
	comorbidity.Comorbidity = dto.Comorbidity
	return comorbidity
}

// MapComorbiditiesToNames maps a list of Comorbidity models to a list of strings (names of comorbidities)
func MapComorbiditiesToNames(comorbidities []*Model) []string {
	comorbidityNames := make([]string, 0)
	for _, model := range comorbidities {
		comorbidityNames = append(comorbidityNames, model.Comorbidity)
	}
	return comorbidityNames
}
