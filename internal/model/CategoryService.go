package model

type CategoryService struct {
	CatID         uint   `gorm:"primaryKey"`
	CatName       string `gorm:"not null"`
	CatEstTime    uint   `gorm:"not null"`
	CatRangePrice uint   `gorm:"not null"`
	CreatedAt     string `gorm:"not null"`
	UpdatedAt     string
}
