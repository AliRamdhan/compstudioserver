package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AliRamdhan/compstudioserver/api"
	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	// Perform migrations
	// if err := config.AutoMigrate(); err != nil {
	// 	// Handle error
	// 	log.Fatalf("Error applying migration: %v", err)
	// }
	// log.Println("Migration Applied Successfully")
	// users, roles := config.SeedData()
	// log.Printf("Seeded %d users and %d roles into the database", len(users), len(roles))
	r := gin.Default()

	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// PORT := os.Getenv("PORT")
	// Initialize the product service
	authenticationService := services.NewAuthService()
	roleService := services.NewRoleService()
	trackStatusService := services.NewTrackStatusService()
	categoryService := services.NewCategoryService()
	compService := services.NewServiceComp()
	trackService := services.NewTrackService()
	productService := services.NewProductServices()
	messageService := services.NewMessageService()

	//Setup middleware
	r.Use(enableCORS())
	r.Use(jsonContentTypeMiddleware())

	// Setup routes
	api.ServiceAuth(r, authenticationService)
	api.ServiceRole(r, roleService)
	api.ServiceTrackStatus(r, trackStatusService)
	api.ServiceCategory(r, categoryService)
	api.ServiceTrackComp(r, trackService)
	api.Servicecomp(r, compService)
	api.ServiceProducts(r, productService)
	api.MessageService(r, messageService)

	r.GET("/", func(c *gin.Context) {
		// message := fmt.Sprintf("Hello World %s")
		c.String(http.StatusOK, "Hello World")
	})
	// port := fmt.Sprintf(PORT)
	r.Run()
}

func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Add Authorization header

		// Check if the request is for CORS preflight
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		c.Next()
	}
}

func jsonContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set JSON Content-Type
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
