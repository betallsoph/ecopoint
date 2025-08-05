package main

import (
	"log"
	"net"

	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/grpc-auth-service/server"
	pb "ecopoint-backend/proto"

	"google.golang.org/grpc"
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

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register Auth service
	authServer := server.NewAuthServer(cfg)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	// Listen on port
	port := "9081"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("gRPC Auth Service listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}