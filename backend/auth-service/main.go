package main

import (
	"log"
	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/auth-service/handlers"
	"ecopoint-backend/auth-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	err := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.DisconnectMongoDB()

	// Setup Gin
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(cfg)

	// Routes
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/google", authHandler.GoogleSignIn)
			auth.POST("/phone/verify", authHandler.PhoneVerify)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.GET("/profile", middleware.AuthMiddleware(cfg.JWTSecret), authHandler.GetProfile)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "auth-service"})
	})

	port := cfg.Port
	if port == "" {
		port = "8081"
	}

	log.Printf("Auth Service starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}