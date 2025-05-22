package dto

import (
	"biometric-data-backend/models"
)

// BiometricDataCreateDTO is used for creating a new biometric record
type BiometricDataCreateDTO struct {
	O2Saturation float64 `json:"o2_saturation"`
	HeartRate    float64 `json:"heart_rate"`
}

// BiometricDataUpdateDTO is used for updating an existing biometric record
type BiometricDataUpdateDTO struct {
	O2Saturation float64 `json:"o2_saturation"`
	HeartRate    float64 `json:"heart_rate"`
}

// BiometricDataDTO is used for retrieving a biometric record
type BiometricDataDTO struct {
	O2Saturation float64 `json:"o2_saturation"`
	HeartRate    float64 `json:"heart_rate"`
}

// MapBiometricDataToDTO maps a BiometricDataData model to a BiometricDataDTO
func MapBiometricDataToDTO(biometric *models.BiometricData) *BiometricDataDTO {
	return &BiometricDataDTO{
		O2Saturation: biometric.O2Saturation,
		HeartRate:    biometric.HeartRate,
	}
}

// MapBiometricDataToDTOs maps a list of BiometricDataData models to a list of BiometricDataDTOs
func MapBiometricDataToDTOs(biometrics []*models.BiometricData) []*BiometricDataDTO {
	var biometricDTOs []*BiometricDataDTO
	for _, biometric := range biometrics {
		biometricDTOs = append(biometricDTOs, MapBiometricDataToDTO(biometric))
	}
	return biometricDTOs
}
