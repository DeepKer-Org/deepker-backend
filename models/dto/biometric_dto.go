package dto

import "biometric-data-backend/models"

// BiometricCreateDTO is used for creating a new biometric record
type BiometricCreateDTO struct {
	AlertID                string `json:"alert_id"`
	O2Saturation           int    `json:"o2_saturation"`
	HeartRate              int    `json:"heart_rate"`
	SystolicBloodPressure  int    `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int    `json:"diastolic_blood_pressure"`
}

// BiometricUpdateDTO is used for updating an existing biometric record
type BiometricUpdateDTO struct {
	AlertID                string `json:"alert_id"`
	O2Saturation           int    `json:"o2_saturation"`
	HeartRate              int    `json:"heart_rate"`
	SystolicBloodPressure  int    `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int    `json:"diastolic_blood_pressure"`
}

// BiometricDTO is used for retrieving a biometric record
type BiometricDTO struct {
	BiometricsID           string `json:"biometrics_id"`
	AlertID                string `json:"alert_id"`
	O2Saturation           int    `json:"o2_saturation"`
	HeartRate              int    `json:"heart_rate"`
	SystolicBloodPressure  int    `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int    `json:"diastolic_blood_pressure"`
}

// MapBiometricToDTO maps a Biometric model to a BiometricDTO
func MapBiometricToDTO(biometric *models.Biometric) *BiometricDTO {
	return &BiometricDTO{
		BiometricsID:           biometric.BiometricsID,
		AlertID:                biometric.AlertID,
		O2Saturation:           biometric.O2Saturation,
		HeartRate:              biometric.HeartRate,
		SystolicBloodPressure:  biometric.SystolicBloodPressure,
		DiastolicBloodPressure: biometric.DiastolicBloodPressure,
	}
}

// MapBiometricsToDTOs maps a list of Biometric models to a list of BiometricDTOs
func MapBiometricsToDTOs(biometrics []*models.Biometric) []*BiometricDTO {
	var biometricDTOs []*BiometricDTO
	for _, biometric := range biometrics {
		biometricDTOs = append(biometricDTOs, MapBiometricToDTO(biometric))
	}
	return biometricDTOs
}
