package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"biometric-data-backend/redis"
	"biometric-data-backend/repository"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type MonitoringDeviceService interface {
	CreateMonitoringDevice(deviceDTO *dto.MonitoringDeviceCreateDTO) error
	GetMonitoringDeviceByID(id string) (*dto.MonitoringDeviceDTO, error)
	GetAllMonitoringDevices(page int, limit int, filters dto.MonitoringDeviceFilter) ([]*dto.MonitoringDeviceDTO, int, error)
	GetAllMonitoringDevicesByStatus(status string) ([]*dto.MonitoringDeviceDTO, int, error)
	UpdateMonitoringDevice(id string, deviceDTO *dto.MonitoringDeviceUpdateDTO) error
	DeleteMonitoringDevice(id string) error
}

type monitoringDeviceService struct {
	repo  repository.MonitoringDeviceRepository
	cache *redis.CacheManager
}

func NewMonitoringDeviceService(repo repository.MonitoringDeviceRepository, cache *redis.CacheManager) MonitoringDeviceService {
	return &monitoringDeviceService{repo: repo, cache: cache}
}

func (s *monitoringDeviceService) CreateMonitoringDevice(deviceDTO *dto.MonitoringDeviceCreateDTO) error {
	device := &models.MonitoringDevice{
		Status: deviceDTO.Status,
	}

	err := s.repo.CreateMonitoringDevice(device)
	if err != nil {
		log.Printf("Failed to create monitoring device: %v", err)
		return err
	}
	log.Println("Monitoring device created successfully with DeviceID:", device.DeviceID)

	// Invalidate cache for all devices
	_ = s.cache.Delete(context.Background(), "monitoring_devices:all")
	return nil
}

func (s *monitoringDeviceService) GetMonitoringDeviceByID(id string) (*dto.MonitoringDeviceDTO, error) {
	ctx := context.Background()
	cacheKey := "monitoring_device:" + id

	// Attempt to fetch from cache
	var device dto.MonitoringDeviceDTO
	found, err := s.cache.Get(ctx, cacheKey, &device)
	if err != nil {
		log.Printf("Error accessing cache for DeviceID %s: %v", id, err)
		return nil, err
	}
	if found {
		log.Println("Cache hit for monitoring device with DeviceID:", id)
		return &device, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching monitoring device with DeviceID:", id)
	dbDevice, err := s.repo.GetMonitoringDeviceByID(id)
	if err != nil {
		return nil, err
	}
	if dbDevice == nil {
		log.Println("No monitoring device found with DeviceID:", id)
		return nil, nil
	}

	device = *dto.MapMonitoringDeviceToDTO(dbDevice)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, device); err != nil {
		log.Printf("Failed to cache monitoring device: %v", err)
	}

	return &device, nil
}

func (s *monitoringDeviceService) GetAllMonitoringDevices(page int, limit int, filters dto.MonitoringDeviceFilter) ([]*dto.MonitoringDeviceDTO, int, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("monitoring_devices:all:page=%d:limit=%d:filters=%v", page, limit, filters)

	// Attempt to fetch from cache
	var devices []*dto.MonitoringDeviceDTO
	var totalCount int
	found, err := s.cache.Get(ctx, cacheKey, &devices)
	if err != nil {
		log.Printf("Error accessing cache for all monitoring devices: %v", err)
		return nil, 0, err
	}
	if found {
		log.Println("Cache hit for all monitoring devices")
		return devices, totalCount, nil
	}

	// Fetch from database if not in cache
	log.Println("Fetching all monitoring devices")
	offset := (page - 1) * limit
	dbDevices, err := s.repo.GetAllMonitoringDevices(offset, limit, filters)
	if err != nil {
		return nil, 0, err
	}

	totalCount64, err := s.repo.CountAllMonitoringDevices(filters)
	if err != nil {
		return nil, 0, err
	}
	totalCount = int(totalCount64)

	devices = dto.MapMonitoringDevicesToDTOs(dbDevices)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, devices); err != nil {
		log.Printf("Failed to cache monitoring devices: %v", err)
	}

	return devices, totalCount, nil
}

func (s *monitoringDeviceService) GetAllMonitoringDevicesByStatus(status string) ([]*dto.MonitoringDeviceDTO, int, error) {
	ctx := context.Background()
	cacheKey := "monitoring_devices:status:" + status

	// Attempt to fetch from cache
	var devices []*dto.MonitoringDeviceDTO
	found, err := s.cache.Get(ctx, cacheKey, &devices)
	if err != nil {
		log.Printf("Error accessing cache for devices by status %s: %v", status, err)
		return nil, 0, err
	}
	if found {
		log.Println("Cache hit for devices with status:", status)
		return devices, len(devices), nil
	}

	// Fetch from database if not in cache
	dbDevices, err := s.repo.GetDevicesByStatus(status)
	if err != nil {
		return nil, 0, err
	}

	devices = dto.MapMonitoringDevicesToDTOs(dbDevices)

	// Store in cache
	if err := s.cache.Set(ctx, cacheKey, devices); err != nil {
		log.Printf("Failed to cache devices by status: %v", err)
	}

	return devices, len(devices), nil
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

	// Update fields
	device.Status = deviceDTO.Status
	device.LinkedByID = deviceDTO.LinkedByID
	device.PatientID = deviceDTO.PatientID

	err = s.repo.UpdateMonitoringDevice(device)
	if err != nil {
		log.Printf("Failed to update monitoring device: %v", err)
		return err
	}
	log.Println("Monitoring device updated successfully with DeviceID:", device.DeviceID)

	// Invalidate cache
	_ = s.cache.Delete(context.Background(), "monitoring_device:"+id, "monitoring_devices:all")
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

	// Invalidate cache
	_ = s.cache.Delete(context.Background(), "monitoring_device:"+id, "monitoring_devices:all")
	return nil
}
