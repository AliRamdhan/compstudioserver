package model

type Role struct {
	RoleID      uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   string `gorm:"not null"`
	UpdatedAt   string
}
