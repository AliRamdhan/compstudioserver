package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type TrackStatusService struct{}

func NewTrackStatusService() *TrackStatusService {
	return &TrackStatusService{}
}

func (ts *TrackStatusService) CreateTrackStatus(trackStatus *model.Status) error {
	trackStatus.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(trackStatus).Error
}

func (ts *TrackStatusService) GetAllTrackStatus() ([]model.Status, error) {
	var trackstatus []model.Status
	if err := config.DB.Find(&trackstatus).Error; err != nil {
		return nil, err
	}
	return trackstatus, nil
}

func (ts *TrackStatusService) UpdateTrackStatus(statusID uint, updatedTrackStatus *model.Status) error {
	var existingTrackStatus model.Status
	if err := config.DB.First(&existingTrackStatus, "status_id = ?", statusID).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingTrackStatus.StatusName = updatedTrackStatus.StatusName
	existingTrackStatus.StatusDescription = updatedTrackStatus.StatusDescription
	existingTrackStatus.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated product
	return config.DB.Save(&existingTrackStatus).Error
}

func (ps *TrackStatusService) DeleteTrackStatus(statusID uint) error {
	// Find the product with the given ID
	var trackStatus model.Status
	if err := config.DB.First(&trackStatus, "status_id = ?", statusID).Error; err != nil {
		return err // trackStatus not found
	}
	// Delete the trackStatus
	return config.DB.Delete(&trackStatus).Error
}
