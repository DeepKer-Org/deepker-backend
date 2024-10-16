package dto

type PatientForAlertDTO struct {
	DNI           string                `json:"dni"`
	Name          string                `json:"name"`
	Location      string                `json:"location"`
	Age           int                   `json:"age"`
	Sex           string                `json:"sex"`
	Doctors       []string              `json:"doctors"`
	Comorbidities []string              `json:"comorbidities"`
	Medications   []*ShortMedicationDTO `json:"medications"`
}
