package grpc

import (
	"log"
	"net"

	"ecopoint/booking-service/internal/services"
	"google.golang.org/grpc"
)

type Server struct {
	bookingService *services.BookingService
	grpcServer     *grpc.Server
}

func NewServer(bookingService *services.BookingService) *Server {
	grpcServer := grpc.NewServer()
	
	return &Server{
		bookingService: bookingService,
		grpcServer:     grpcServer,
	}
}

func (s *Server) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening on %s", address)
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}