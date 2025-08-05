package main

import (
	"log"
	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/order-service/handlers"
	"ecopoint-backend/order-service/middleware"

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
	orderHandler := handlers.NewOrderHandler()

	// Routes
	api := r.Group("/api/v1")
	{
		orders := api.Group("/orders")
		orders.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("", orderHandler.GetOrders)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.PUT("/:id/accept", orderHandler.AcceptOrder)
			orders.PUT("/:id/complete", orderHandler.CompleteOrder)
			orders.GET("/available", orderHandler.GetAvailableOrders) // For collectors
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "order-service"})
	})

	port := "8083"
	log.Printf("Order Service starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}