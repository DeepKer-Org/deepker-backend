package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"biometric-data-backend/utils"
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
	GetAllAlertsByTimezone(timezone string) ([]*dto.AlertDTO, error)
}

type alertService struct {
	alertRepo              repository.AlertRepository
	patientRepo            repository.PatientRepository
	biometricRepo          repository.BiometricDataRepository
	computerDiagnosticRepo repository.ComputerDiagnosticRepository
	doctorRepo             repository.DoctorRepository
	monitoringDeviceRepo   repository.MonitoringDeviceRepository
	phoneRepo              repository.PhoneRepository
}

func NewAlertService(alertRepo repository.AlertRepository, biometricRepo repository.BiometricDataRepository, computerDiagnosticRepo repository.ComputerDiagnosticRepository, doctorRepo repository.DoctorRepository, monitoringDeviceRepo repository.MonitoringDeviceRepository, phoneRepo repository.PhoneRepository, patientRepo repository.PatientRepository) AlertService {
	return &alertService{
		alertRepo:              alertRepo,
		biometricRepo:          biometricRepo,
		computerDiagnosticRepo: computerDiagnosticRepo,
		doctorRepo:             doctorRepo,
		monitoringDeviceRepo:   monitoringDeviceRepo,
		phoneRepo:              phoneRepo,
		patientRepo:            patientRepo,
	}
}

func (s *alertService) CreateAlert(alertDTO *dto.AlertCreateDTO) (*dto.AlertCreateResponseDTO, error) {
	tx := s.alertRepo.BeginTransaction()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return &dto.AlertCreateResponseDTO{Message: "Transaction start failed"}, tx.Error
	}

	defer func() {
		// Rollback transaction if it's not committed yet
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to error: %v", tx.Error)
		}
	}()

	// Get Monitoring Device
	device, err := s.monitoringDeviceRepo.GetMonitoringDeviceByID(alertDTO.DeviceID)
	if err != nil {
		log.Printf("Device not found: %v", err)
		return &dto.AlertCreateResponseDTO{Message: "Device not found"}, err
	}

	// Create Biometric Data
	biometricData := &models.BiometricData{
		O2Saturation: alertDTO.O2Saturation,
		HeartRate:    alertDTO.HeartRate,
	}

	err = s.biometricRepo.CreateInTransaction(biometricData, tx)
	if err != nil {
		log.Printf("Failed to create biometric data: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to create biometric data"}, err
	}

	// Create Computer Diagnostic
	computerDiagnostic := &models.ComputerDiagnostic{
		Diagnosis:  alertDTO.Diagnosis,
		Percentage: alertDTO.Percentage,
	}

	err = s.computerDiagnosticRepo.CreateInTransaction(computerDiagnostic, tx)
	if err != nil {
		log.Printf("Failed to create computer diagnostic: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to create computer diagnostic"}, err
	}

	// Fetch Patient Information
	patient, err := s.patientRepo.GetByID(device.PatientID, "patient_id")
	if err != nil {
		log.Printf("Failed to fetch patient information: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to fetch patient information"}, err
	}

	// Parse the timezone from the alertDTO
	location, err := time.LoadLocation(alertDTO.Timezone)
	if err != nil {
		// Fallback to UTC if the provided timezone is invalid
		fmt.Printf("Invalid timezone provided (%s), defaulting to UTC: %v\n", alertDTO.Timezone, err)
		location = time.UTC
	}

	// Get the current time in the specified timezone and convert it to UTC
	localTime := time.Now().In(location)
	utcTime := localTime.UTC()

	// Create Alert
	alert := &models.Alert{
		AlertTimestamp:     utcTime,
		AttendedTimestamp:  nil,
		AttendedBy:         nil,
		BiometricDataID:    biometricData.BiometricDataID,
		BiometricData:      biometricData,
		DiagnosticID:       computerDiagnostic.DiagnosticID,
		ComputerDiagnostic: computerDiagnostic,
		PatientID:          patient.PatientID,
		Patient:            patient,
	}

	err = s.alertRepo.CreateInTransaction(alert, tx)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to create alert"}, err
	}

	// Commit transaction if all operations succeeded
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return &dto.AlertCreateResponseDTO{Message: "Failed to commit transaction"}, err
	}

	// Get all push tokens in the phones table
	pushTokens, err := s.phoneRepo.GetPushTokens()

	if len(pushTokens) != 0 {
		notificationTitle := "Alerta Crítica: " + computerDiagnostic.Diagnosis
		notificationBody := "Paciente: " + patient.Name + ", Ubicación: " + patient.Location

		// Send push notifications to all push tokens
		err = utils.SendExponentPushNotifications(pushTokens, notificationTitle, notificationBody)
		if err != nil {
			return nil, err
		}
	}

	// Build Response DTO
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
		utcTimestamp := alertDTO.AttendedTimestamp.UTC()
		alert.AttendedTimestamp = &utcTimestamp
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

func (s *alertService) GetAllAlertsByTimezone(timezone string) ([]*dto.AlertDTO, error) {
	var err error
	var alerts []*models.Alert

	alerts, err = s.alertRepo.GetAlertsByTimezone(timezone)

	if err != nil {
		log.Printf("Error retrieving alerts for the timezone: %v", err)
		return nil, err
	}

	alertDTOs := dto.MapAlertsToDTOs(alerts)
	log.Printf("Alerts fetched successfully with count: %d", len(alerts))
	return alertDTOs, nil
}
