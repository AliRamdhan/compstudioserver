package model

type Status struct {
	StatusID          uint   `gorm:"primaryKey"`
	StatusName        string `gorm:"not null"`
	StatusDescription string `gorm:"not null"`
	CreatedAt         string `gorm:"not null"`
	UpdatedAt         string
}
