package clients

import (
	"log"

	pb_auth "ecopoint-backend/proto"
	pb_order "ecopoint-backend/proto"
	pb_user "ecopoint-backend/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	AuthClient  pb_auth.AuthServiceClient
	UserClient  pb_user.UserServiceClient
	OrderClient pb_order.OrderServiceClient
}

func NewGRPCClients() (*GRPCClients, error) {
	// Connect to Auth Service
	authConn, err := grpc.Dial("localhost:9081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Connect to User Service
	userConn, err := grpc.Dial("localhost:9082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Connect to Order Service
	orderConn, err := grpc.Dial("localhost:9083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	clients := &GRPCClients{
		AuthClient:  pb_auth.NewAuthServiceClient(authConn),
		UserClient:  pb_user.NewUserServiceClient(userConn),
		OrderClient: pb_order.NewOrderServiceClient(orderConn),
	}

	log.Println("gRPC clients initialized successfully")
	return clients, nil
}