package model

type Product struct {
	ProductID    uint   `gorm:"primaryKey"`
	ProductName  string `gorm:"unique;not null"`
	ProductPrice uint   `gorm:"not null"`
	ProductLink  string `gorm:"not null"`
	ProductImage string `gorm:"not null"`
	CreatedAt    string `gorm:"not null"`
	UpdatedAt    string
}
