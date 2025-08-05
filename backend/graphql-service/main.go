package main

import (
	"log"
	"net/http"
	"os"

	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/graphql-service/generated"
	"ecopoint-backend/graphql-service/middleware"
	"ecopoint-backend/graphql-service/resolvers"
	"ecopoint-backend/graphql-service/clients"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

const defaultPort = "4000"

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	err := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.DisconnectMongoDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
		AllowCredentials: true,
	}))

	// Initialize gRPC clients
	grpcClients, err := clients.NewGRPCClients()
	if err != nil {
		log.Fatal("Failed to initialize gRPC clients:", err)
	}

	// Initialize GraphQL resolver
	resolver := &resolvers.Resolver{
		Config:      cfg,
		GRPCClients: grpcClients,
	}

	// Create GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// GraphQL endpoint with authentication middleware
	r.POST("/query", middleware.AuthMiddleware(), func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL Playground (development only)
	if gin.Mode() != gin.ReleaseMode {
		r.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
		})
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "graphql-service"})
	})

	log.Printf("GraphQL Server starting on port %s", port)
	log.Printf("GraphQL Playground: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}