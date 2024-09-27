package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/repository"
	"errors"
	"gorm.io/gorm"
	"log"
)

type MonitoringDeviceService interface {
	CreateMonitoringDevice(deviceDTO *dto.MonitoringDeviceCreateDTO) error
	GetMonitoringDeviceByID(id string) (*dto.MonitoringDeviceDTO, error)
	GetAllMonitoringDevices() ([]*dto.MonitoringDeviceDTO, error)
	UpdateMonitoringDevice(id string, deviceDTO *dto.MonitoringDeviceUpdateDTO) error
	DeleteMonitoringDevice(id string) error
}

type monitoringDeviceService struct {
	repo repository.MonitoringDeviceRepository
}

func NewMonitoringDeviceService(repo repository.MonitoringDeviceRepository) MonitoringDeviceService {
	return &monitoringDeviceService{repo: repo}
}

func (s *monitoringDeviceService) CreateMonitoringDevice(deviceDTO *dto.MonitoringDeviceCreateDTO) error {
	device := &models.MonitoringDevice{
		Type:      deviceDTO.Type,
		Status:    deviceDTO.Status,
		PatientID: deviceDTO.PatientID,
		Sensors:   deviceDTO.Sensors,
	}

	err := s.repo.CreateMonitoringDevice(device)
	if err != nil {
		log.Printf("Failed to create monitoring device: %v", err)
		return err
	}
	log.Println("Monitoring device created successfully with DeviceID:", device.DeviceID)
	return nil
}

func (s *monitoringDeviceService) GetMonitoringDeviceByID(id string) (*dto.MonitoringDeviceDTO, error) {
	log.Println("Fetching monitoring device with DeviceID:", id)
	device, err := s.repo.GetMonitoringDeviceByID(id)
	if err != nil {
		log.Printf("Error retrieving monitoring device: %v", err)
		return nil, err
	}
	if device == nil {
		log.Println("No monitoring device found with DeviceID:", id)
		return nil, nil
	}

	deviceDTO := dto.MapMonitoringDeviceToDTO(device)
	log.Println("Monitoring device fetched successfully with DeviceID:", id)
	return deviceDTO, nil
}

func (s *monitoringDeviceService) GetAllMonitoringDevices() ([]*dto.MonitoringDeviceDTO, error) {
	log.Println("Fetching all monitoring devices")
	devices, err := s.repo.GetAllMonitoringDevices()
	if err != nil {
		log.Printf("Error retrieving monitoring devices: %v", err)
		return nil, err
	}

	deviceDTOs := dto.MapMonitoringDevicesToDTOs(devices)
	log.Println("Monitoring devices fetched successfully, total count:", len(deviceDTOs))
	return deviceDTOs, nil
}

func (s *monitoringDeviceService) UpdateMonitoringDevice(id string, deviceDTO *dto.MonitoringDeviceUpdateDTO) error {
	log.Println("Updating monitoring device with DeviceID:", id)

	device, err := s.repo.GetMonitoringDeviceByID(id)
	if err != nil {
		log.Printf("Error retrieving monitoring device: %v", err)
		return err
	}
	if device == nil {
		log.Printf("Monitoring device not found with DeviceID: %v", id)
		return gorm.ErrRecordNotFound
	}

	device.Status = deviceDTO.Status
	device.PatientID = deviceDTO.PatientID
	device.Sensors = deviceDTO.Sensors

	err = s.repo.UpdateMonitoringDevice(device)
	if err != nil {
		log.Printf("Failed to update monitoring device: %v", err)
		return err
	}
	log.Println("Monitoring device updated successfully with DeviceID:", device.DeviceID)
	return nil
}

func (s *monitoringDeviceService) DeleteMonitoringDevice(id string) error {
	log.Println("Deleting monitoring device with DeviceID:", id)
	err := s.repo.DeleteMonitoringDevice(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Monitoring device not found with DeviceID:", id)
			return nil
		}
		log.Printf("Failed to delete monitoring device: %v", err)
		return err
	}
	log.Println("Monitoring device deleted successfully with DeviceID:", id)
	return nil
}
