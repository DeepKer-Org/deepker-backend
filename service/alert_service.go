package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"biometric-data-backend/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlertService interface {
	CreateAlert(alertDTO *dto.AlertCreateDTO) (*dto.AlertCreateResponseDTO, error)
	GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error)
	GetAllAlerts() ([]*dto.AlertDTO, error)
	UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error
	DeleteAlert(id uuid.UUID) error
	GetAllAlertsByPeriod(period string, page int, limit int) ([]*dto.AlertDTO, int, error)
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
	cache                  *redis.CacheManager
}

func NewAlertService(
	alertRepo repository.AlertRepository,
	biometricRepo repository.BiometricDataRepository,
	computerDiagnosticRepo repository.ComputerDiagnosticRepository,
	doctorRepo repository.DoctorRepository,
	monitoringDeviceRepo repository.MonitoringDeviceRepository,
	phoneRepo repository.PhoneRepository,
	patientRepo repository.PatientRepository,
	cache *redis.CacheManager,
) AlertService {
	return &alertService{
		alertRepo:              alertRepo,
		biometricRepo:          biometricRepo,
		computerDiagnosticRepo: computerDiagnosticRepo,
		doctorRepo:             doctorRepo,
		monitoringDeviceRepo:   monitoringDeviceRepo,
		phoneRepo:              phoneRepo,
		patientRepo:            patientRepo,
		cache:                  cache,
	}
}

func (s *alertService) CreateAlert(alertDTO *dto.AlertCreateDTO) (*dto.AlertCreateResponseDTO, error) {
	tx := s.alertRepo.BeginTransaction()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return &dto.AlertCreateResponseDTO{Message: "Transaction start failed"}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to error: %v", tx.Error)
		}
	}()

	device, err := s.monitoringDeviceRepo.GetMonitoringDeviceByID(alertDTO.DeviceID)
	if err != nil {
		log.Printf("Device not found: %v", err)
		return &dto.AlertCreateResponseDTO{Message: "Device not found"}, err
	}

	if device.Status != "In Use" {
		log.Printf("Device is not in use")
		return &dto.AlertCreateResponseDTO{Message: "Device is not in use"}, errors.New("device is not in use")
	}

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

	patient, err := s.patientRepo.GetByID(device.PatientID, "patient_id")
	if err != nil {
		log.Printf("Failed to fetch patient information: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to fetch patient information"}, err
	}

	location, err := time.LoadLocation(alertDTO.Timezone)
	if err != nil {
		fmt.Printf("Invalid timezone provided (%s), defaulting to UTC: %v\n", alertDTO.Timezone, err)
		location = time.UTC
	}

	localTime := time.Now().In(location)
	utcTime := localTime.UTC()

	alert := &models.Alert{
		AlertTimestamp:     utcTime,
		AttendedTimestamp:  nil,
		AttendedBy:         nil,
		BiometricDataID:    biometricData.BiometricDataID,
		BiometricData:      biometricData,
		DiagnosticID:       computerDiagnostic.DiagnosticID,
		ComputerDiagnostic: computerDiagnostic,
		PatientID:          patient.PatientID,
	}

	err = s.alertRepo.CreateInTransaction(alert, tx)
	if err != nil {
		log.Printf("Failed to create alert: %v", err)
		tx.Rollback()
		return &dto.AlertCreateResponseDTO{Message: "Failed to create alert"}, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return &dto.AlertCreateResponseDTO{Message: "Failed to commit transaction"}, err
	}

	pushTokens, err := s.phoneRepo.GetPushTokens()
	if len(pushTokens) != 0 {
		notificationTitle := "Alerta Crítica: " + computerDiagnostic.Diagnosis
		notificationBody := "Paciente: " + patient.Name + ", Ubicación: " + patient.Location

		err = utils.SendExponentPushNotifications(pushTokens, notificationTitle, notificationBody)
		if err != nil {
			return nil, err
		}
	}

	alertResponse := &dto.AlertCreateResponseDTO{
		AlertID: alert.AlertID.String(),
		Message: "Alert created successfully",
	}

	// Invalidate relevant caches
	_ = s.cache.Delete(context.Background(), "alerts:all")

	log.Printf("Alert created successfully with AlertID: %s", alert.AlertID)
	return alertResponse, nil
}

func (s *alertService) GetAlertByID(id uuid.UUID) (*dto.AlertDTO, error) {
	var alert dto.AlertDTO

	log.Println("Fetching alert with AlertID:", id)
	dbAlert, err := s.alertRepo.GetByID(id, "alert_id")
	if err != nil {
		log.Printf("Error retrieving alert: %v", err)
		return nil, err
	}
	if dbAlert == nil {
		log.Println("No alert found with AlertID:", id)
		return nil, nil
	}

	alert = *dto.MapAlertToDTO(dbAlert)

	return &alert, nil
}

func (s *alertService) GetAllAlerts() ([]*dto.AlertDTO, error) {
	ctx := context.Background()
	cacheKey := "alerts:all"

	// Attempt to fetch from cache
	var alerts []*dto.AlertDTO
	found, err := s.cache.Get(ctx, cacheKey, &alerts)
	if err != nil {
		log.Printf("Error accessing cache for all alerts: %v", err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for all alerts")
		return alerts, nil
	}

	log.Println("Fetching all alerts")
	dbAlerts, err := s.alertRepo.GetAll()
	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, err
	}

	alerts = dto.MapAlertsToDTOs(dbAlerts)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, alerts); err != nil {
		log.Printf("Failed to cache alerts: %v", err)
	}

	return alerts, nil
}

func (s *alertService) UpdateAlert(id uuid.UUID, alertDTO *dto.AlertUpdateDTO) error {

	alert, err := s.alertRepo.GetByID(id, "alert_id")
	if err != nil {
		log.Printf("Error retrieving alert: %v", err)
		return err
	}
	if alert == nil {
		log.Printf("Alert not found with AlertID: %v", id)
		return gorm.ErrRecordNotFound
	}

	if alertDTO.FinalDiagnosis != "" {
		alert.FinalDiagnosis = alertDTO.FinalDiagnosis
		err := s.alertRepo.UpdateAlert(alert)
		if err != nil {
			log.Printf("Failed to update alert: %v", err)
			return err
		}
		return nil
	}

	if alertDTO.AttendedByID == uuid.Nil {
		err := s.alertRepo.Liberate(alert)
		if err != nil {
			log.Printf("Failed to liberate alert: %v", err)
			return err
		}

		_ = s.cache.Delete(context.Background(), "alert:"+id.String(), "alerts:all")
		return nil
	}

	// Update AttendedByID if provided
	if alertDTO.AttendedByID != uuid.Nil {
		alert.AttendedByID = uuid.NullUUID{
			UUID:  alertDTO.AttendedByID,
			Valid: true,
		}
	}

	// // Only allow updates to other fields if AttendedByID is set
	if alert.AttendedByID.Valid {
		utcTimestamp := alertDTO.AttendedTimestamp.UTC()
		alert.AttendedTimestamp = &utcTimestamp
	}

	updateAlert := &models.Alert{
		AlertTimestamp:     alert.AlertTimestamp,
		AttendedTimestamp:  alert.AttendedTimestamp,
		AttendedByID:       uuid.NullUUID{UUID: alertDTO.AttendedByID, Valid: true},
		BiometricDataID:    alert.BiometricDataID,
		BiometricData:      alert.BiometricData,
		DiagnosticID:       alert.DiagnosticID,
		ComputerDiagnostic: alert.ComputerDiagnostic,
		PatientID:          alert.PatientID,
	}

	err = s.alertRepo.Update(updateAlert, "alert_id", id)
	if err != nil {
		log.Printf("Failed to update alert: %v", err)
		return err
	}
	log.Println("Alert updated successfully with AlertID:", alert.AlertID)

	// Invalidate cache for updated alert and all alerts
	_ = s.cache.Delete(context.Background(), "alert:"+id.String(), "alerts:all")
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

	// Invalidate cache for deleted alert and all alerts
	_ = s.cache.Delete(context.Background(), "alert:"+id.String(), "alerts:all")
	return nil
}

func (s *alertService) GetAllAlertsByPeriod(period string, page int, limit int) ([]*dto.AlertDTO, int, error) {
	period = strings.ToLower(period)
	log.Printf("Fetching alerts with period: %s, page: %d, limit: %d", period, page, limit)

	offset := (page - 1) * limit
	var alerts []*models.Alert
	var totalCount int64
	var err error

	// Define cache key
	// cacheKey := fmt.Sprintf("alerts:period:%s:page:%d:limit:%d", period, page, limit)

	// Attempt to fetch from cache
	var alertDTOs []*dto.AlertDTO
	// found, cacheErr := s.cache.Get(context.Background(), cacheKey, &alertDTOs)
	// if cacheErr != nil {
	// 	log.Printf("Error accessing cache for period %s: %v", period, cacheErr)
	// }
	// if found {
	// 	log.Println("Cache hit for alerts by period:", period)
	// 	return alertDTOs, int(totalCount), nil
	// }

	// Fetch data based on period
	switch period {
	case "recent":
		err = s.alertRepo.CountAlertsByPeriod("recent", &totalCount)
		if err == nil {
			alerts, err = s.alertRepo.GetRecentAlerts(offset, limit)
		}
	case "past":
		err = s.alertRepo.CountAlertsByPeriod("past", &totalCount)
		if err == nil {
			alerts, err = s.alertRepo.GetPastAlerts(offset, limit)
		}
	default:
		log.Printf("Invalid period: %s", period)
		return nil, 0, errors.New("invalid period: must be 'recent' or 'past'")
	}

	if err != nil {
		log.Printf("Error retrieving alerts: %v", err)
		return nil, 0, err
	}

	alertDTOs = dto.MapAlertsToDTOs(alerts)

	// Store in cache
	// if cacheErr == nil {
	// 	_ = s.cache.Set(context.Background(), cacheKey, alertDTOs)
	// }

	// log.Printf("Alerts fetched successfully with period: %s, count: %d", period, len(alerts))
	return alertDTOs, int(totalCount), nil
}

func (s *alertService) GetAllAlertsByTimezone(timezone string) ([]*dto.AlertDTO, error) {
	cacheKey := "alerts:timezone:" + timezone

	var alerts []*dto.AlertDTO

	dbAlerts, err := s.alertRepo.GetAlertsByTimezone(timezone)
	if err != nil {
		log.Printf("Error retrieving alerts for timezone %s: %v", timezone, err)
		return nil, err
	}

	alerts = dto.MapAlertsToDTOs(dbAlerts)

	_ = s.cache.Set(context.Background(), cacheKey, alerts)

	log.Printf("Alerts fetched successfully for timezone: %s, count: %d", timezone, len(alerts))
	return alerts, nil
}
