package main

import (
	"log"
	"os"

	"ecopoint/api-gateway/internal/handlers"
	"ecopoint/api-gateway/internal/middleware"
	"ecopoint/api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// Health check endpoint
	r.GET("/health", handlers.HealthCheck)

	// API routes
	api := r.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/", handlers.CreateUser)
			users.GET("/me", middleware.FirebaseAuth(), handlers.GetCurrentUser)
			users.PUT("/me", middleware.FirebaseAuth(), handlers.UpdateUser)
			users.GET("/:id", middleware.FirebaseAuth(), handlers.GetUser)
		}

		// Address routes
		addresses := api.Group("/addresses")
		addresses.Use(middleware.FirebaseAuth())
		{
			addresses.POST("/", handlers.CreateAddress)
			addresses.GET("/", handlers.GetUserAddresses)
			addresses.PUT("/:id", handlers.UpdateAddress)
			addresses.DELETE("/:id", handlers.DeleteAddress)
		}

		// Booking routes
		bookings := api.Group("/bookings")
		bookings.Use(middleware.FirebaseAuth())
		{
			bookings.POST("/", handlers.CreateBooking)
			bookings.GET("/", handlers.GetBookings)
			bookings.GET("/:id", handlers.GetBooking)
			bookings.PUT("/:id", handlers.UpdateBooking)
			bookings.DELETE("/:id", handlers.CancelBooking)
			bookings.POST("/:id/assign", handlers.AssignCollector)
			bookings.PUT("/:id/status", handlers.UpdateBookingStatus)
		}

		// Collector routes
		collectors := api.Group("/collectors")
		collectors.Use(middleware.FirebaseAuth())
		{
			collectors.GET("/nearby", handlers.GetNearbyCollectors)
			collectors.GET("/:id/location", handlers.GetCollectorLocation)
			collectors.PUT("/:id/location", handlers.UpdateCollectorLocation)
			collectors.GET("/:id/bookings", handlers.GetCollectorBookings)
		}

		// Chat routes
		chat := api.Group("/chat")
		chat.Use(middleware.FirebaseAuth())
		{
			chat.GET("/conversations", handlers.GetConversations)
			chat.POST("/conversations", handlers.CreateConversation)
			chat.GET("/conversations/:id/messages", handlers.GetMessages)
			chat.POST("/conversations/:id/messages", handlers.SendMessage)
		}

		// Real-time tracking routes
		tracking := api.Group("/tracking")
		tracking.Use(middleware.FirebaseAuth())
		{
			tracking.GET("/:bookingId/events", handlers.GetTrackingEvents)
			tracking.POST("/:bookingId/location", handlers.UpdateLocation)
		}
	}

	// GraphQL endpoint (proxy to user service)
	r.POST("/graphql", middleware.FirebaseAuth(), handlers.GraphQLProxy)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 API Gateway starting on port %s", port)
	log.Printf("🌍 Environment: %s", cfg.Environment)
	
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}