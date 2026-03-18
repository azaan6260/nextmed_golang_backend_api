package main

import (
	"log"
	"nextmed-backend/config"
	"nextmed-backend/models"
	"nextmed-backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// 2. Initialize Database Connection
	config.InitDatabase()

	// 3. Auto-Migrate Models
	log.Println("Running database migrations...")
	err = config.DB.AutoMigrate(
		&models.User{},
		&models.PatientProfile{},
		&models.Facility{},
		&models.FacilityDoctor{},
		&models.Appointment{},
		&models.Post{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// 4. Initialize Gin Router
	r := gin.Default()

	// 5. CORS Middleware
	r.Use(cors.Default())

	// 6. Health Check Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Nextmed 2.0 API is running",
		})
	})

	// 7. Setup Routes
	routes.SetupAuthRoutes(r)
	routes.SetupBookingRoutes(r)
	routes.SetupFeedRoutes(r)

	// 7. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
