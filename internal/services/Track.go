package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/google/uuid"
)

type TrackService struct{}

func NewTrackService() *TrackService {
	return &TrackService{}
}

func (ts *TrackService) CreateTrackService(track *model.Track, serviceId uint) error {
	track.TrackNumber = uuid.New()
	track.ServiceId = serviceId
	track.TrackStatusRefer = 1 //Preparation
	track.TrackDescription = "Checking"
	track.TrackStaff = "Laduny Staff"
	track.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(track).Error
}

func (ts *TrackService) GetAllTrackService() ([]model.Track, error) {
	var tracks []model.Track
	if err := config.DB.Preload("Service").Preload("Status").Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

func (ts *TrackService) GetTrackStatusByTrackNumber(trackNumber uuid.UUID) ([]model.Track, error) {
	var track []model.Track
	if err := config.DB.Preload("Service").Preload("Status").Where("track_number = ?", trackNumber).Find(&track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func (ts *TrackService) CreateProgressTrackStatusByTrackNumber(track *model.Track, trackNumber uuid.UUID, serviceId uint) error {
	track.TrackNumber = trackNumber
	track.ServiceId = serviceId
	track.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(track).Error
}

func (ts *TrackService) UpdateTrackService(trakId uint, newstatus uint) error {
	var existingTrack model.Track
	if err := config.DB.First(&existingTrack, "track_id = ?", trakId).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingTrack.TrackStatusRefer = newstatus
	existingTrack.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Save(&existingTrack).Error
}

func (ts *TrackService) DeleteTrack(trackId uint) error {
	// Find the product with the given ID
	var track model.Track
	if err := config.DB.First(&track, "track_id = ?", trackId).Error; err != nil {
		return err // track not found
	}
	// Delete the track
	return config.DB.Delete(&track).Error
}
