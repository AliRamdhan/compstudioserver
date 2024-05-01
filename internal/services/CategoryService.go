package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (sc *CategoryService) CreateServiceCategory(serviceCategory *model.CategoryService) error {
	serviceCategory.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(serviceCategory).Error
}

func (sc *CategoryService) GetAllServiceCategory() ([]model.CategoryService, error) {
	var serviceCategory []model.CategoryService
	if err := config.DB.Find(&serviceCategory).Error; err != nil {
		return nil, err
	}
	return serviceCategory, nil
}

func (sc *CategoryService) UpdateServiceCategory(serviceCategoryId uint, updateServiceCategory *model.CategoryService) error {
	var existingServiceCategory model.CategoryService
	if err := config.DB.First(&existingServiceCategory, "cat_id = ?", serviceCategoryId).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingServiceCategory.CatName = updateServiceCategory.CatName
	existingServiceCategory.CatEstTime = updateServiceCategory.CatEstTime
	existingServiceCategory.CatRangePrice = updateServiceCategory.CatRangePrice
	existingServiceCategory.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated product
	return config.DB.Save(&existingServiceCategory).Error
}

func (ps *CategoryService) DeleteServiceCategory(serviceCategoryId uint) error {
	// Find the product with the given ID
	var serviceCategory model.CategoryService
	if err := config.DB.First(&serviceCategory, "cat_id = ?", serviceCategoryId).Error; err != nil {
		return err // serviceCategory not found
	}
	// Delete the serviceCategory
	return config.DB.Delete(&serviceCategory).Error
}
