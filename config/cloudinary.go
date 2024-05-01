package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// cldSecret := "dU7ueb02HlTpGcqWi3LKLnhH9ic"
	// cldName := "dxv8lvmch"
	// cldKey := "398233461444433"

	cldSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cldName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cldKey := os.Getenv("CLOUDINARY_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
