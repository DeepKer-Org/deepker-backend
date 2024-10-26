package dto

import "biometric-data-backend/models"

type PhoneCreateDTO struct {
	ExponentPushToken string `json:"exponent_push_token"`
}

func MapCreateDTOToPhone(dto *PhoneCreateDTO) *models.Phone {
	return &models.Phone{
		ExponentPushToken: dto.ExponentPushToken,
	}
}
