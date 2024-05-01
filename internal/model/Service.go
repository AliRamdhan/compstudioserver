package model

type Service struct {
	// gorm.Model
	ServiceID            uint            `gorm:"primaryKey"`
	ServiceCustonmerName string          `gorm:"not null"`
	ServiceLaptopName    string          `gorm:"not null"`
	ServiceLaptopVersion string          `gorm:"not null"`
	ServiceDate          string          `gorm:"not null"`
	ServiceEstTime       string          `gorm:"not null"`
	ServiceComplaint     string          `gorm:"not null"`
	IsCompleteService    string          `gorm:"not null"`
	CustomerUser         uint            `gorm:"not null"`
	User                 User            `gorm:"foreignKey:CustomerUser"`
	ServiceCategory      uint            `gorm:"not null"`
	CategoryService      CategoryService `gorm:"foreignKey:ServiceCategory"`
	CreatedAt            string          `gorm:"not null"`
	UpdatedAt            string
}

//Hardware Fix
//Software Fix
//Cleaning
//Cleaning & Software
//Cleaning & Hardware
//Complete
