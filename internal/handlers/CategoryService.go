package handlers

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
)

type ServiceCategoryHandler struct {
	serviceCategory *services.CategoryService
}

func NewServiceCategoryHandler(sc *services.CategoryService) *ServiceCategoryHandler {
	return &ServiceCategoryHandler{serviceCategory: sc}
}

func (sh *ServiceCategoryHandler) CreateServiceCategory(c *gin.Context) {
	var serviceCategory model.CategoryService
	if err := c.ShouldBindJSON(&serviceCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sh.serviceCategory.CreateServiceCategory(&serviceCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Service Category created successfully", "ServiceCategory": serviceCategory})
}

func (sh *ServiceCategoryHandler) GetAllServiceCategory(c *gin.Context) {
	serviceCategory, err := sh.serviceCategory.GetAllServiceCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Service Category", "ServiceCategory": serviceCategory})
}

func (sh *ServiceCategoryHandler) UpdateServiceCategory(c *gin.Context) {
	var serviceCategory model.CategoryService
	if err := c.ShouldBindJSON(&serviceCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceCategoryIDStr := c.Param("catId")
	var serviceCategoryId uint
	_, err := fmt.Sscanf(serviceCategoryIDStr, "%d", &serviceCategoryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service category status ID"})
		return
	}

	if err := sh.serviceCategory.UpdateServiceCategory(serviceCategoryId, &serviceCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service Category updated successfully", "Service Category": serviceCategory})
}

func (sh *ServiceCategoryHandler) DeleteServiceCategory(c *gin.Context) {
	serviceCategoryIDStr := c.Param("catId")
	var serviceCategoryId uint
	_, err := fmt.Sscanf(serviceCategoryIDStr, "%d", &serviceCategoryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service category status ID"})
		return
	}

	if err := sh.serviceCategory.DeleteServiceCategory(serviceCategoryId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service Category deleted successfully"})
}
