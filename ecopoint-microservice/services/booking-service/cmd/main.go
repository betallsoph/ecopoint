package main

import (
	"log"
	"os"

	"ecopoint/booking-service/internal/handlers"
	"ecopoint/booking-service/internal/services"
	"ecopoint/booking-service/pkg/grpc"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	// Initialize services
	bookingService := services.NewBookingService()
	
	// Initialize gRPC server
	grpcServer := grpc.NewServer(bookingService)
	
	// Start gRPC server in goroutine
	go func() {
		log.Printf("🚀 Booking Service gRPC server starting on port %s", port)
		if err := grpcServer.Start(":" + port); err != nil {
			log.Fatal("Failed to start gRPC server:", err)
		}
	}()

	// Start HTTP server for health checks
	router := gin.Default()
	
	// Health check endpoint
	router.GET("/health", handlers.HealthCheck)
	
	// Start HTTP server
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}
	
	log.Printf("🏥 Health check server starting on port %s", httpPort)
	if err := router.Run(":" + httpPort); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}