package dto

type PatientUpdateDTO struct {
	DNI      string  `json:"dni"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	Sex      string  `json:"sex"`
	Location string  `json:"location,omitempty"`
}
