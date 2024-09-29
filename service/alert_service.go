package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type AlertService interface {
	CreateAlert(alertDTO *dto.AlertCreateDTO) error
	GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error)
	GetAllAlerts() ([]*dto.AlertDTO, error)
	UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error
	DeleteAlert(id uuid.UUID) error
}

type alertService struct {
	alertRepo              repository.AlertRepository
	biometricRepo          repository.BiometricRepository
	computerDiagnosticRepo repository.ComputerDiagnosticRepository
	doctorRepo             repository.DoctorRepository
}

func NewAlertService(alertRepo repository.AlertRepository, biometricRepo repository.BiometricRepository, computerDiagnosticRepo repository.ComputerDiagnosticRepository, doctorRepo repository.DoctorRepository) AlertService {
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

	// Fetch related entities: Biometrics, Computer Diagnostics, and Doctors
	biometrics, err := s.biometricRepo.GetBiometricsByAlertID(id)
	if err != nil {
		log.Printf("Error retrieving biometrics for alert: %v", err)
		return nil, err
	}

	computerDiagnostics, err := s.computerDiagnosticRepo.GetComputerDiagnosticsByAlertID(id)
	if err != nil {
		log.Printf("Error retrieving computer diagnostics for alert: %v", err)
		return nil, err
	}

	doctors, err := s.doctorRepo.GetDoctorsByAlertID(id)
	if err != nil {
		log.Printf("Error retrieving doctors for alert: %v", err)
		return nil, err
	}

	// Map related entities to DTOs
	alertDTO := dto.MapAlertToDTO(alert)
	alertDTO.Biometrics = dto.MapBiometricsToDTOs(biometrics)
	alertDTO.ComputerDiagnostics = dto.MapComputerDiagnosticsToDTOs(computerDiagnostics)
	alertDTO.Doctors = dto.MapDoctorsToDTOs(doctors)

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
		biometrics, err := s.biometricRepo.GetBiometricsByAlertID(alertDTO.AlertID)
		if err != nil {
			log.Printf("Error retrieving biometrics for alert: %v", err)
			return nil, err
		}
		alertDTO.Biometrics = dto.MapBiometricsToDTOs(biometrics)

		computerDiagnostics, err := s.computerDiagnosticRepo.GetComputerDiagnosticsByAlertID(alertDTO.AlertID)
		if err != nil {
			log.Printf("Error retrieving computer diagnostics for alert: %v", err)
			return nil, err
		}
		alertDTO.ComputerDiagnostics = dto.MapComputerDiagnosticsToDTOs(computerDiagnostics)

		doctors, err := s.doctorRepo.GetDoctorsByAlertID(alertDTO.AlertID)
		if err != nil {
			log.Printf("Error retrieving doctors for alert: %v", err)
			return nil, err
		}
		alertDTO.Doctors = dto.MapDoctorsToDTOs(doctors)
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
			return nil
		}
		log.Printf("Failed to delete alert: %v", err)
		return err
	}
	log.Println("Alert deleted successfully with AlertID:", id)
	return nil
}
