package dtos

type CreatePatientDTO struct {
	DNI            string   `json:"dni" binding:"required"`
	Name           string   `json:"name" binding:"required"`
	Age            int      `json:"age"`
	Weight         float64  `json:"weight"`
	Height         float64  `json:"height"`
	Sex            string   `json:"sex"`
	Location       string   `json:"location"`
	CurrentState   string   `json:"current_state"`
	FinalDiagnosis string   `json:"final_diagnosis"`
	Comorbidities  []string `json:"comorbidities"`
	Medications    []string `json:"medications"`
	MedicalStaff   []string `json:"medical_staff"`
}

type UpdatePatientDTO struct {
	Name           *string  `json:"name"`
	Age            *int     `json:"age"`
	Weight         *float64 `json:"weight"`
	Height         *float64 `json:"height"`
	Location       *string  `json:"location"`
	CurrentState   *string  `json:"current_state"`
	FinalDiagnosis *string  `json:"final_diagnosis"`
}
