package dto

type PatientForAlertDTO struct {
	DNI           string                `json:"dni"`
	Name          string                `json:"name"`
	Location      string                `json:"location"`
	Age           int                   `json:"age"`
	Sex           string                `json:"sex"`
	Comorbidities []string              `json:"comorbidities"`
	Medications   []*ShortMedicationDTO `json:"medications"`
}
