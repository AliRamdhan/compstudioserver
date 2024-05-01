package model

import (
	"github.com/google/uuid"
)

type Track struct {
	// gorm.Model
	TrackID          uint      `gorm:"primaryKey"`
	TrackNumber      uuid.UUID `gorm:"not null"`
	TrackStatusRefer uint      `gorm:"not null"`
	TrackDescription string    `gorm:"not null"`
	Status           Status    `gorm:"foreignKey:TrackStatusRefer"`
	ServiceId        uint      `gorm:"not null"`
	Service          Service   `gorm:"foreignKey:ServiceId"`
	TrackStaff       string    `gorm:"not null"`
	CreatedAt        string    `gorm:"not null"`
	UpdatedAt        string
}
