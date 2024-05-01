package handlers

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TrackHandler struct {
	trackService *services.TrackService
}

func NewTrackHandler(th *services.TrackService) *TrackHandler {
	return &TrackHandler{trackService: th}
}
func (th *TrackHandler) CreatetrackService(c *gin.Context) {
	var trackService model.Track
	if err := c.ShouldBindJSON(&trackService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract serviceCategoryID and customerId from the request body
	serviceID := trackService.ServiceId

	// Check if serviceCategoryID is missing or invalid
	if serviceID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service SERVICE ID"})
		return
	}
	// Proceed with service creation
	if err := th.trackService.CreateTrackService(&trackService, serviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Track created successfully", "Track": trackService})
}

func (th *TrackHandler) CreateProgressTrackStatusByTrackNumber(c *gin.Context) {
	var trackService model.Track
	if err := c.ShouldBindJSON(&trackService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trackNumber := c.Param("trackNumber")

	// Parse track number
	tracknumber, err := uuid.Parse(trackNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track number format"})
		return
	}
	// Extract serviceCategoryID and customerId from the request body
	serviceID := trackService.ServiceId

	// Check if serviceCategoryID is missing or invalid
	if serviceID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service category ID"})
		return
	}
	if tracknumber == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing track number"})
		return
	}

	// Proceed with service creation
	if err := th.trackService.CreateProgressTrackStatusByTrackNumber(&trackService, tracknumber, serviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Track created successfully", "Track": trackService})
}

func (th *TrackHandler) GetAllTrack(c *gin.Context) {
	track, err := th.trackService.GetAllTrackService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All Track ", "Tracks": track})
}

func (th *TrackHandler) GetTrackStatusByTrackNumber(c *gin.Context) {
	trackNumber := c.Param("trackNumber")

	// Parse track number
	trackID, err := uuid.Parse(trackNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track number format"})
		return
	}

	track, err := th.trackService.GetTrackStatusByTrackNumber(trackID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Track Status", "Track": track})
}

func (th *TrackHandler) UpdateTrack(c *gin.Context) {
	var track model.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackIdStr := c.Param("trackId")
	var trackId uint
	_, err := fmt.Sscanf(trackIdStr, "%d", &trackId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track status ID"})
		return
	}
	trackStatusID := track.TrackStatusRefer

	// Check if serviceCategoryID is missing or invalid
	if trackStatusID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service category ID"})
		return
	}

	if err := th.trackService.UpdateTrackService(trackId, trackStatusID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status of Track updated successfully", "Track Status": track})
}

func (th *TrackHandler) DeleteTrack(c *gin.Context) {
	trackIdStr := c.Param("trackId")
	// productID, err := uuid.Parse(trackIdStr)
	var trackId uint
	_, err := fmt.Sscanf(trackIdStr, "%d", &trackId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := th.trackService.DeleteTrack(trackId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})

}
