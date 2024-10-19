package dto

import (
	"biometric-data-backend/models"
	"time"
)

type MedicalVisitDTO struct {
	Reason        string     `json:"reason"`
	Diagnosis     string     `json:"diagnosis"`
	Treatment     string     `json:"treatment"`
	EntryDate     *time.Time `json:"entry_date"`
	DischargeDate string     `json:"discharge_date"`
}

// MapMedicalVisitToDTO maps a MedicalVisit model to a MedicalVisitDTO
func MapMedicalVisitToDTO(medicalVisit *models.MedicalVisit) *MedicalVisitDTO {
	var dischargeDate string
	if medicalVisit.DischargeDate != nil {
		dischargeDate = medicalVisit.DischargeDate.Format(time.RFC3339)
	}
	return &MedicalVisitDTO{
		Reason:        medicalVisit.Reason,
		Diagnosis:     medicalVisit.Diagnosis,
		Treatment:     medicalVisit.Treatment,
		EntryDate:     medicalVisit.EntryDate,
		DischargeDate: dischargeDate,
	}
}

// MapMedicalVisitsToDTOs maps a slice of MedicalVisit models to a slice of MedicalVisitDTOs
func MapMedicalVisitsToDTOs(medicalVisits []*models.MedicalVisit) []*MedicalVisitDTO {
	medicalVisitDTOs := make([]*MedicalVisitDTO, 0)
	for _, medicalVisit := range medicalVisits {
		medicalVisitDTOs = append(medicalVisitDTOs, MapMedicalVisitToDTO(medicalVisit))
	}
	return medicalVisitDTOs
}

func FindEntryDate(medicalVisits []*models.MedicalVisit) string {
	if len(medicalVisits) == 0 {
		return ""
	}

	// Loop through the medical visits to find one with an open discharge date
	for _, visit := range medicalVisits {
		if visit.DischargeDate == nil { // DischargeDate is open if it's nil
			return visit.EntryDate.Format(time.RFC3339)
		}
	}

	// No current medical visit found (i.e., all have discharge dates)
	return ""
}
