package handlers

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
)

type TrackStatusHandler struct {
	trackStatusService *services.TrackStatusService
}

func NewTrackStatusHandler(ts *services.TrackStatusService) *TrackStatusHandler {
	return &TrackStatusHandler{trackStatusService: ts}
}

func (th *TrackStatusHandler) CreateTrackStatus(c *gin.Context) {
	var trackStatus model.Status
	if err := c.ShouldBindJSON(&trackStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.trackStatusService.CreateTrackStatus(&trackStatus); err != nil {
		//ph.productService.CreateProduct(&trackStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Track created successfully", "Track Status": trackStatus})
}

func (th *TrackStatusHandler) GetAllTrackStatus(c *gin.Context) {
	trackStatus, err := th.trackStatusService.GetAllTrackStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Track Status", "Track Status": trackStatus})
}

func (th *TrackStatusHandler) UpdateTrackStatus(c *gin.Context) {
	var trackStatus model.Status
	if err := c.ShouldBindJSON(&trackStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackStatusIDStr := c.Param("statusId")
	var trackStatusId uint
	_, err := fmt.Sscanf(trackStatusIDStr, "%d", &trackStatusId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track status ID"})
		return
	}

	if err := th.trackStatusService.UpdateTrackStatus(trackStatusId, &trackStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Track Status updated successfully", "Track Status": trackStatus})
}

func (th *TrackStatusHandler) DeleteTrackStatus(c *gin.Context) {
	trackStatusIDStr := c.Param("statusId")
	var trackStatusId uint
	_, err := fmt.Sscanf(trackStatusIDStr, "%d", &trackStatusId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track status ID"})
		return
	}

	if err := th.trackStatusService.DeleteTrackStatus(trackStatusId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trak Status deleted successfully"})
}
