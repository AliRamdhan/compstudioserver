package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")

	// Construct the DSN string using environment variables
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s", dbUsername, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(&model.User{}, &model.User{}, &model.Track{}, &model.Status{}, &model.Service{}, &model.CategoryService{}, &model.Product{}, &model.Messages{})
	// return DB.AutoMigrate(&model.Client{}, &model.Profile{})
}
func SeedData() ([]model.User, []model.Role) {
	var roles = []model.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "visitor", Description: "Unauthenticated customer role"},
	}
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), 14)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	var user = []model.User{
		{
			Username:  os.Getenv("ADMIN_USERNAME"),
			Email:     os.Getenv("ADMIN_EMAIL"),
			Password:  string(hashedPassword),
			RoleUser:  1,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	if err := DB.Save(&roles).Error; err != nil {
		log.Fatalf("Error saving roles: %v", err)
	}
	if err := DB.Save(&user).Error; err != nil {
		log.Fatalf("Error saving users: %v", err)
	}

	return user, roles
}
