package handlers

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
)

type ServiceCompHandler struct {
	serviceComp *services.ServiceComp
}

func NewServiceCompHandler(sc *services.ServiceComp) *ServiceCompHandler {
	return &ServiceCompHandler{serviceComp: sc}
}
func (sh *ServiceCompHandler) CreateserviceComp(c *gin.Context) {

	var serviceComp model.Service
	if err := c.ShouldBindJSON(&serviceComp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract serviceCategoryID and customerId from the request body
	serviceCategoryID := serviceComp.ServiceCategory
	customerID := serviceComp.CustomerUser

	// Check if serviceCategoryID is missing or invalid
	if serviceCategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service category ID"})
		return
	}

	// Check if customerId is missing or invalid
	if customerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing customer ID"})
		return
	}

	// Proceed with service creation
	if err := sh.serviceComp.CreateServiceComp(&serviceComp, customerID, serviceCategoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Service created successfully", "Service": serviceComp})
}

func (sh *ServiceCompHandler) GetAllService(c *gin.Context) {
	serviceComp, err := sh.serviceComp.GetAllServiceComp()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Service ", "Service Category": serviceComp})
}
func (sh *ServiceCompHandler) GetDetailServiceById(c *gin.Context) {
	serviceIdStr := c.Param("serviceId")

	var serviceId uint
	_, err := fmt.Sscanf(serviceIdStr, "%d", &serviceId)
	// trackID, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket id format"})
		return
	}
	services, err := sh.serviceComp.GetDetailServiceById(serviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Details tickets", "Service": services})
}

func (sh *ServiceCompHandler) UpdateService(c *gin.Context) {
	var service model.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serviceCat := service.ServiceCategory
	// Check if serviceCategoryID is missing or invalid
	if serviceCat == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service category ID"})
		return
	}
	serviceIdStr := c.Param("serviceId")
	// productID, err := uuid.Parse(serviceIdStr)
	var serviceId uint
	_, err := fmt.Sscanf(serviceIdStr, "%d", &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	if err := sh.serviceComp.UpdateServiceComp(&service, serviceId, serviceCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": service})
}

func (sh *ServiceCompHandler) DeleteService(c *gin.Context) {
	serviceIdStr := c.Param("serviceId")
	// productID, err := uuid.Parse(serviceIdStr)
	var serviceId uint
	_, err := fmt.Sscanf(serviceIdStr, "%d", &serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := sh.serviceComp.DeleteService(serviceId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
