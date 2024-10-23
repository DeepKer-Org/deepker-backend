package dto

type LoginRequest struct {
	UserType   string `json:"user_type"`
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type ChangePasswordDTO struct {
	DNI          string `json:"dni"`
	IssuanceDate string `json:"issuance_date"`
	NewPassword  string `json:"new_password"`
}
