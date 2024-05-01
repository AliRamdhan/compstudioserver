package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	Username  string `json:"username" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	WANumber  uint   `json:"wanumber"`
	Address   uint   `json:"address"`
	Password  string `json:"password"`
	RoleUser  uint   `json:"roleuser" gorm:"not null"`
	Role      Role   `gorm:"foreignKey:RoleUser"`
	CreatedAt string `gorm:"not null"`
	UpdatedAt string
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
