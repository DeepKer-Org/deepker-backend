package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type AlertService interface {
	CreateAlert(alertDTO *dto.AlertCreateDTO) (*dto.AlertCreateResponseDTO, error)
	GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error)
	GetAllAlerts() ([]*dto.AlertDTO, error)
	UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error
	DeleteAlert(id uuid.UUID) error
	GetAllAlertsByStatus(status string, page int, limit int) ([]*dto.AlertDTO, int, error)
}

type alertService struct {
	alertRepo              repository.AlertRepository
	patientRepo            repository.PatientRepository
	biometricRepo          repository.BiometricDataRepository
	computerDiagnosticRepo repository.ComputerDiagnosticRepository
	doctorRepo             repository.DoctorRepository
}

func NewAlertService(alertRepo repository.AlertRepository, biometricRepo repository.BiometricDataRepository, computerDiagnosticRepo repository.ComputerDiagnosticRepository, doctorRepo repository.DoctorRepository, patientRepo repository.PatientRepository) AlertService {
	return &alertService{
		alertRepo:              alertRepo,
		biometricRepo:          biometricRepo,
		computerDiagnosticRepo: computerDiagnosticRepo,
		doctorRepo:             doctorRepo,
		patientRepo:            patientRepo,
	}
}

func (s *alertService) CreateAlert(alertDTO *dto.AlertCreateDTO) (*dto.AlertCreateResponseDTO, error) {
	biometricData, err := s.biometricRepo.GetByID(alertDTO.BiometricDataID, "biometric_data_id")
	if err != nil {
		log.Printf("Error retrieving biometric data with ID: %s, error: %v", alertDTO.BiometricDataID, err)
		return &dto.AlertCreateResponseDTO{
			Message: fmt.Sprintf("Failed to retrieve biometric data with ID: %s", alertDTO.BiometricDataID),
		}, err
	}

	patient, err := s.patientRepo.GetByID(alertDTO.PatientID, "patient_id")
	if err != nil {
		log.Printf("Error retrieving patient with ID: %s, error: %v", alertDTO.PatientID, err)
		return &dto.AlertCreateResponseDTO{
			Message: fmt.Sprintf("Failed to retrieve patient with ID: %s", alertDTO.PatientID),
		}, err
	}

	computerDiagnosticIDs := convertUUIDsToInterface(alertDTO.ComputerDiagnosticIDs)
	computerDiagnostics, err := s.computerDiagnosticRepo.GetByIDs(computerDiagnosticIDs, "diagnostic_id")
	if err != nil {
		log.Printf("Error retrieving computer diagnostics with IDs: %v, error: %v", alertDTO.ComputerDiagnosticIDs, err)
		return &dto.AlertCreateResponseDTO{
			Message: "Failed to retrieve computer diagnostics",
		}, err
	}

	alert := &models.Alert{
		AlertTimestamp:      time.Now(),
		AttendedTimestamp:   nil,
		AttendedBy:          nil,
		BiometricDataID:     alertDTO.BiometricDataID,
		BiometricData:       biometricData,
		PatientID:           alertDTO.PatientID,
		Patient:             patient,
		ComputerDiagnostics: computerDiagnostics,
	}

	err = s.alertRepo.Create(alert)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		return &dto.AlertCreateResponseDTO{
			Message: "Failed to create alert",
		}, err
	}

	alertResponse := &dto.AlertCreateResponseDTO{
		AlertID: alert.AlertID.String(),
		Message: "Alert created successfully",
	}

	log.Printf("Alert created successfully with AlertID: %s", alert.AlertID)
	return alertResponse, nil
}

func (s *alertService) GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error) {
	log.Println("Fetching alert with AlertID:", id)
	alert, err := s.alertRepo.GetByID(id, "alert_id")
	if err != nil {
		log.Printf("Error retrieving alert: %v", err)
		return nil, err
	}
	if alert == nil {
		log.Println("No alert found with AlertID:", id)
		return nil, err
	}
	log.Println("Alert fetched successfully with AlertID:", id)
	return dto.MapAlertToDTO(alert), nil
}

func (s *alertService) GetAllAlerts() ([]*dto.AlertDTO, error) {
	log.Println("Fetching all alerts")
	alerts, err := s.alertRepo.GetAll()
	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, err
	}
	log.Println("Alerts fetched successfully, total count:", len(alerts))
	return dto.MapAlertsToDTOs(alerts), nil
}

func (s *alertService) UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error {
	log.Println("Updating alert with AlertID:", id)

	alert, err := s.alertRepo.GetByID(id, "alert_id")
	if err != nil {
		log.Printf("Error retrieving alert: %v", err)
		return err
	}
	if alert == nil {
		log.Printf("Alert not found with AlertID: %v", id)
		return gorm.ErrRecordNotFound
	}

	log.Println("This is my alertDTO: ", alertDTO)
	log.Println("This is my alert: ", alert)

	if alert.AttendedByID.Valid == false && alertDTO.AttendedByID == uuid.Nil {
		log.Printf("AttendedByID must be set before updating other fields")
		return errors.New("attendedById must be set before updating other fields")
	}

	// Update AttendedByID if provided
	if alertDTO.AttendedByID != uuid.Nil {
		alert.AttendedByID = uuid.NullUUID{
			UUID:  alertDTO.AttendedByID,
			Valid: true,
		}
	}

	// Only allow updates to other fields if AttendedByID is set
	if alert.AttendedByID.Valid {
		alert.AttendedTimestamp = alertDTO.AttendedTimestamp
	} else {
		log.Println("Cannot update fields other than AttendedByID as it has not been set.")
		return errors.New("other fields cannot be updated until attendedById is set")
	}

	err = s.alertRepo.Update(alert, "alert_id", id)
	if err != nil {
		log.Printf("Failed to update alert: %v", err)
		return err
	}
	log.Println("Alert updated successfully with AlertID:", alert.AlertID)
	return nil
}

func (s *alertService) DeleteAlert(id uuid.UUID) error {
	log.Println("Deleting alert with AlertID:", id)
	err := s.alertRepo.Delete(id, "alert_id")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Alert not found with AlertID:", id)
			return errors.New("alert not found")
		}
		log.Printf("Failed to delete alert: %v", err)
		return err
	}
	log.Println("Alert deleted successfully with AlertID:", id)
	return nil
}

func (s *alertService) GetAllAlertsByStatus(status string, page int, limit int) ([]*dto.AlertDTO, int, error) {
	status = strings.ToLower(status)
	log.Printf("Fetching alerts with status: %s, page: %d, limit: %d", status, page, limit)

	offset := (page - 1) * limit
	var alerts []*models.Alert
	var totalCount int64
	var err error

	// Count total alerts with the given status and fetch paginated results
	switch status {
	case "attended":
		// Count attended alerts and then fetch paginated results if no error
		err = s.alertRepo.CountAlertsByStatus("attended", &totalCount)
		if err == nil {
			alerts, err = s.alertRepo.GetAttendedAlerts(offset, limit)
		}
	case "unattended":
		// Count unattended alerts and then fetch paginated results if no error
		err = s.alertRepo.CountAlertsByStatus("unattended", &totalCount)
		if err == nil {
			alerts, err = s.alertRepo.GetUnattendedAlerts(offset, limit)
		}
	default:
		log.Printf("Invalid status: %s", status)
		return nil, 0, errors.New("invalid status: must be 'attended' or 'unattended'")
	}

	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, 0, err
	}

	alertDTOs := dto.MapAlertsToDTOs(alerts)
	log.Printf("Alerts fetched successfully with status: %s, count: %d", status, len(alerts))
	return alertDTOs, int(totalCount), nil
}

// Convert uuid slice to []interface{}
func convertUUIDsToInterface(uuids []uuid.UUID) []interface{} {
	interfaceSlice := make([]interface{}, len(uuids))
	for i, v := range uuids {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
