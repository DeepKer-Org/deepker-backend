package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// BiometricDataCreateDTO is used for creating a new biometric record
type BiometricDataCreateDTO struct {
	O2Saturation           int `json:"o2_saturation"`
	HeartRate              int `json:"heart_rate"`
	SystolicBloodPressure  int `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int `json:"diastolic_blood_pressure"`
}

// BiometricDataUpdateDTO is used for updating an existing biometric record
type BiometricDataUpdateDTO struct {
	O2Saturation           int `json:"o2_saturation"`
	HeartRate              int `json:"heart_rate"`
	SystolicBloodPressure  int `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int `json:"diastolic_blood_pressure"`
}

// BiometricDataDTO is used for retrieving a biometric record
type BiometricDataDTO struct {
	BiometricDataDataID    uuid.UUID `json:"biometric_data_id"`
	O2Saturation           int       `json:"o2_saturation"`
	HeartRate              int       `json:"heart_rate"`
	SystolicBloodPressure  int       `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int       `json:"diastolic_blood_pressure"`
}

// MapBiometricDataToDTO maps a BiometricDataData model to a BiometricDataDTO
func MapBiometricDataToDTO(biometric *models.BiometricData) *BiometricDataDTO {
	return &BiometricDataDTO{
		BiometricDataDataID:    biometric.BiometricDataID,
		O2Saturation:           biometric.O2Saturation,
		HeartRate:              biometric.HeartRate,
		SystolicBloodPressure:  biometric.SystolicBloodPressure,
		DiastolicBloodPressure: biometric.DiastolicBloodPressure,
	}
}

// MapBiometricRecordsToDTOs maps a list of BiometricDataData models to a list of BiometricDataDTOs
func MapBiometricRecordsToDTOs(biometrics []*models.BiometricData) []*BiometricDataDTO {
	var biometricDTOs []*BiometricDataDTO
	for _, biometric := range biometrics {
		biometricDTOs = append(biometricDTOs, MapBiometricDataToDTO(biometric))
	}
	return biometricDTOs
}
