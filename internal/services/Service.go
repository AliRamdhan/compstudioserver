package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type ServiceComp struct{}

func NewServiceComp() *ServiceComp {
	return &ServiceComp{}
}

func (sc *ServiceComp) CreateServiceComp(service *model.Service, userId uint, categoryId uint) error {
	service.IsCompleteService = "Progress"
	service.CustomerUser = userId
	service.ServiceCategory = categoryId
	service.ServiceEstTime = "3 Days (flexible)"
	service.IsCompleteService = "No"
	service.ServiceDate = time.Now().Format("2006-01-02 15:04:05")
	service.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(service).Error
}

func (sc *ServiceComp) GetAllServiceComp() ([]model.Service, error) {
	var services []model.Service
	if err := config.DB.Preload("CategoryService").Preload("User").Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (sc *ServiceComp) GetDetailServiceById(serviceId uint) (*model.Service, error) {
	var services model.Service
	if err := config.DB.Preload("CategoryService").Preload("User").First(&services, "service_id = ?", serviceId).Error; err != nil {
		return nil, err // Ticket not found
	}
	return &services, nil
}

func (sc *ServiceComp) UpdateServiceComp(updateService *model.Service, serviceId uint, catService uint) error {
	var existingService model.Service
	if err := config.DB.First(&existingService, "service_id = ?", serviceId).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingService.ServiceLaptopName = updateService.ServiceLaptopName
	existingService.ServiceLaptopVersion = updateService.ServiceLaptopVersion
	existingService.ServiceComplaint = updateService.ServiceComplaint
	existingService.ServiceDate = updateService.ServiceDate
	existingService.ServiceEstTime = updateService.ServiceEstTime
	existingService.IsCompleteService = updateService.IsCompleteService
	existingService.CategoryService = updateService.CategoryService
	existingService.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Save(&existingService).Error
}

func (sc *ServiceComp) DeleteService(serviceId uint) error {
	// Find the product with the given ID
	var service model.Service
	if err := config.DB.First(&service, "service_id = ?", serviceId).Error; err != nil {
		return err // service not found
	}
	// Delete the service
	return config.DB.Delete(&service).Error
}
