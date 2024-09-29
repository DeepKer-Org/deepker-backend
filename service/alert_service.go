package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

type AlertService interface {
	CreateAlert(alertDTO *dto.AlertCreateDTO) error
	GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error)
	GetAllAlerts() ([]*dto.AlertDTO, error)
	UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error
	DeleteAlert(id uuid.UUID) error
	GetAllAlertsByStatus(status string) ([]*dto.AlertDTO, error)
}

type alertService struct {
	alertRepo              repository.AlertRepository
	biometricRepo          repository.BiometricDataRepository
	computerDiagnosticRepo repository.ComputerDiagnosticRepository
	doctorRepo             repository.DoctorRepository
}

func NewAlertService(alertRepo repository.AlertRepository, biometricRepo repository.BiometricDataRepository, computerDiagnosticRepo repository.ComputerDiagnosticRepository, doctorRepo repository.DoctorRepository) AlertService {
	return &alertService{
		alertRepo:              alertRepo,
		biometricRepo:          biometricRepo,
		computerDiagnosticRepo: computerDiagnosticRepo,
		doctorRepo:             doctorRepo,
	}
}

func (s *alertService) CreateAlert(alertDTO *dto.AlertCreateDTO) error {
	alert := &models.Alert{
		AlertStatus:    alertDTO.AlertStatus,
		Room:           alertDTO.Room,
		AlertTimestamp: alertDTO.AlertTimestamp,
		PatientID:      alertDTO.PatientID,
	}

	err := s.alertRepo.Create(alert)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		return err
	}
	log.Println("Alert created successfully with AlertID:", alert.AlertID)
	return nil
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
		return nil, nil
	}

	computerDiagnostics, err := s.computerDiagnosticRepo.GetComputerDiagnosticsByAlertID(id)
	if err != nil {
		log.Printf("Error retrieving computer diagnostics for alert: %v", err)
		return nil, err
	}

	/* Reconsider whether to include doctors in alerts
	doctors, err := s.doctorRepo.GetDoctorsByAlertID(id)
	if err != nil {
		log.Printf("Error retrieving doctors for alert: %v", err)
		return nil, err
	}
	*/

	// Map related entities to DTOs
	alertDTO := dto.MapAlertToDTO(alert)
	alertDTO.ComputerDiagnostics = dto.MapComputerDiagnosticsToDTOs(computerDiagnostics)
	//alertDTO.Doctors = dto.MapDoctorsToDTOs(doctors)

	log.Println("Alert fetched successfully with AlertID:", id)
	return alertDTO, nil
}

func (s *alertService) GetAllAlerts() ([]*dto.AlertDTO, error) {
	log.Println("Fetching all alerts")
	alerts, err := s.alertRepo.GetAll()
	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, err
	}

	alertDTOs := dto.MapAlertsToDTOs(alerts)

	// Fetch related entities for each alert
	for _, alertDTO := range alertDTOs {
		computerDiagnostics, err := s.computerDiagnosticRepo.GetComputerDiagnosticsByAlertID(alertDTO.AlertID)
		if err != nil {
			log.Printf("Error retrieving computer diagnostics for alert: %v", err)
			return nil, err
		}
		alertDTO.ComputerDiagnostics = dto.MapComputerDiagnosticsToDTOs(computerDiagnostics)

		/*
			doctors, err := s.doctorRepo.GetDoctorsByAlertID(alertDTO.AlertID)
			if err != nil {
				log.Printf("Error retrieving doctors for alert: %v", err)
				return nil, err
			}
			alertDTO.Doctors = dto.MapDoctorsToDTOs(doctors)*/
	}

	log.Println("Alerts fetched successfully, total count:", len(alertDTOs))
	return alertDTOs, nil
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

	alert.AlertStatus = alertDTO.AlertStatus
	alert.Room = alertDTO.Room
	alert.AttendedTimestamp = alertDTO.AttendedTimestamp
	alert.PatientID = alertDTO.PatientID

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

func (s *alertService) GetAllAlertsByStatus(status string) ([]*dto.AlertDTO, error) {
	status = strings.ToLower(status)
	log.Println("Fetching alerts with status:", status)

	var alerts []*models.Alert
	var err error

	switch status {
	case "attended":
		alerts, err = s.alertRepo.GetAttendedAlerts()
	case "unattended":
		alerts, err = s.alertRepo.GetUnattendedAlerts()
	default:
		log.Printf("Invalid status: %s", status)
		return nil, errors.New("invalid status: must be 'attended' or 'unattended'")
	}

	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, err
	}

	alertDTOs := []*dto.AlertDTO{}

	if len(alerts) > 0 {
		alertDTOs = dto.MapAlertsToDTOs(alerts)
	}
	// Fetch related entities for each alert
	for _, alertDTO := range alertDTOs {
		computerDiagnostics, err := s.computerDiagnosticRepo.GetComputerDiagnosticsByAlertID(alertDTO.AlertID)
		if err != nil {
			log.Printf("Error retrieving computer diagnostics for alert: %v", err)
			return nil, err
		}
		alertDTO.ComputerDiagnostics = dto.MapComputerDiagnosticsToDTOs(computerDiagnostics)
	}

	log.Println("Alerts fetched successfully with status:", status)
	return alertDTOs, nil
}
