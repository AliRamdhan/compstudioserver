package model

type Messages struct {
	MessageId      uint    `gorm:"primaryKey"`
	MessageContent string  `gorm:"not null"`
	MessageISRead  bool    `gorm:"not null"`
	MessageUser    uint    `gorm:"not null"`
	User           User    `gorm:"foreignKey:MessageUser"`
	MessageService uint    `gorm:"not null"`
	Service        Service `gorm:"foreignKey:MessageService"`
	CreatedAt      string  `gorm:"not null"`
	UpdatedAt      string
}
