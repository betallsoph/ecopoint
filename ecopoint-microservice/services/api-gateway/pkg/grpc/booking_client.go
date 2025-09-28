package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingClient struct {
	conn   *grpc.ClientConn
	client BookingServiceClient
}

type BookingServiceClient interface {
	CreateBooking(ctx context.Context, req *CreateBookingRequest) (*CreateBookingResponse, error)
	GetBooking(ctx context.Context, req *GetBookingRequest) (*GetBookingResponse, error)
	ListBookings(ctx context.Context, req *ListBookingsRequest) (*ListBookingsResponse, error)
	UpdateBooking(ctx context.Context, req *UpdateBookingRequest) (*UpdateBookingResponse, error)
	CancelBooking(ctx context.Context, req *CancelBookingRequest) (*CancelBookingResponse, error)
	AssignCollector(ctx context.Context, req *AssignCollectorRequest) (*AssignCollectorResponse, error)
	UpdateBookingStatus(ctx context.Context, req *UpdateBookingStatusRequest) (*UpdateBookingStatusResponse, error)
}

// gRPC Request/Response types (matching booking service)
type CreateBookingRequest struct {
	UserID         string    `json:"user_id"`
	WasteType      string    `json:"waste_type"`
	Quantity       int       `json:"quantity"`
	Description    string    `json:"description"`
	PickupAddress  Address   `json:"pickup_address"`
	ScheduledDate  string    `json:"scheduled_date"`
	ScheduledTime  string    `json:"scheduled_time"`
	Notes          string    `json:"notes"`
	Images         []string  `json:"images"`
}

type Address struct {
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateBookingResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Booking Booking `json:"booking"`
}

type Booking struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CollectorID    *string   `json:"collector_id"`
	WasteType      string    `json:"waste_type"`
	Quantity       int       `json:"quantity"`
	Description    string    `json:"description"`
	PickupAddress  Address   `json:"pickup_address"`
	ScheduledDate  string    `json:"scheduled_date"`
	ScheduledTime  string    `json:"scheduled_time"`
	Status         string    `json:"status"`
	EstimatedPrice *float64  `json:"estimated_price"`
	FinalPrice     *float64  `json:"final_price"`
	Notes          string    `json:"notes"`
	Images         []string  `json:"images"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetBookingRequest struct {
	BookingID string `json:"booking_id"`
}

type GetBookingResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Booking Booking `json:"booking"`
}

type ListBookingsRequest struct {
	UserID      *string `json:"user_id,omitempty"`
	CollectorID *string `json:"collector_id,omitempty"`
	Status      *string `json:"status,omitempty"`
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
}

type ListBookingsResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Bookings []Booking `json:"bookings"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
}

type UpdateBookingRequest struct {
	BookingID     string    `json:"booking_id"`
	WasteType     *string   `json:"waste_type,omitempty"`
	Quantity      *int      `json:"quantity,omitempty"`
	Description   *string   `json:"description,omitempty"`
	PickupAddress *Address  `json:"pickup_address,omitempty"`
	ScheduledDate *string   `json:"scheduled_date,omitempty"`
	ScheduledTime *string   `json:"scheduled_time,omitempty"`
	Notes         *string   `json:"notes,omitempty"`
	Images        []string  `json:"images,omitempty"`
}

type UpdateBookingResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Booking Booking `json:"booking"`
}

type CancelBookingRequest struct {
	BookingID string `json:"booking_id"`
	Reason    string `json:"reason"`
}

type CancelBookingResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AssignCollectorRequest struct {
	BookingID   string `json:"booking_id"`
	CollectorID string `json:"collector_id"`
}

type AssignCollectorResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Booking Booking `json:"booking"`
}

type UpdateBookingStatusRequest struct {
	BookingID string `json:"booking_id"`
	Status    string `json:"status"`
	Notes     string `json:"notes"`
}

type UpdateBookingStatusResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Booking Booking `json:"booking"`
}

func NewBookingClient(address string) (*BookingClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to booking service: %w", err)
	}

	return &BookingClient{
		conn:   conn,
		client: nil, // Will be implemented when gRPC service is ready
	}, nil
}

func (c *BookingClient) Close() error {
	return c.conn.Close()
}

// Mock implementations for now (will be replaced with actual gRPC calls)
func (c *BookingClient) CreateBooking(ctx context.Context, req *CreateBookingRequest) (*CreateBookingResponse, error) {
	log.Printf("Creating booking for user %s", req.UserID)
	
	// TODO: Implement actual gRPC call
	// For now, return mock response
	return &CreateBookingResponse{
		Success: true,
		Message: "Booking created successfully",
		Booking: Booking{
			ID:            "mock-booking-id",
			UserID:        req.UserID,
			WasteType:     req.WasteType,
			Quantity:      req.Quantity,
			Description:   req.Description,
			PickupAddress: req.PickupAddress,
			Status:        "PENDING",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}, nil
}

func (c *BookingClient) GetBooking(ctx context.Context, req *GetBookingRequest) (*GetBookingResponse, error) {
	log.Printf("Getting booking %s", req.BookingID)
	
	// TODO: Implement actual gRPC call
	return &GetBookingResponse{
		Success: true,
		Message: "Booking retrieved successfully",
		Booking: Booking{
			ID:     req.BookingID,
			Status: "PENDING",
		},
	}, nil
}

func (c *BookingClient) ListBookings(ctx context.Context, req *ListBookingsRequest) (*ListBookingsResponse, error) {
	log.Printf("Listing bookings for user %v", req.UserID)
	
	// TODO: Implement actual gRPC call
	return &ListBookingsResponse{
		Success:  true,
		Message:  "Bookings retrieved successfully",
		Bookings: []Booking{},
		Total:    0,
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (c *BookingClient) UpdateBooking(ctx context.Context, req *UpdateBookingRequest) (*UpdateBookingResponse, error) {
	log.Printf("Updating booking %s", req.BookingID)
	
	// TODO: Implement actual gRPC call
	return &UpdateBookingResponse{
		Success: true,
		Message: "Booking updated successfully",
		Booking: Booking{
			ID:     req.BookingID,
			Status: "PENDING",
		},
	}, nil
}

func (c *BookingClient) CancelBooking(ctx context.Context, req *CancelBookingRequest) (*CancelBookingResponse, error) {
	log.Printf("Cancelling booking %s", req.BookingID)
	
	// TODO: Implement actual gRPC call
	return &CancelBookingResponse{
		Success: true,
		Message: "Booking cancelled successfully",
	}, nil
}

func (c *BookingClient) AssignCollector(ctx context.Context, req *AssignCollectorRequest) (*AssignCollectorResponse, error) {
	log.Printf("Assigning collector %s to booking %s", req.CollectorID, req.BookingID)
	
	// TODO: Implement actual gRPC call
	return &AssignCollectorResponse{
		Success: true,
		Message: "Collector assigned successfully",
		Booking: Booking{
			ID:          req.BookingID,
			CollectorID: &req.CollectorID,
			Status:      "ASSIGNED",
		},
	}, nil
}

func (c *BookingClient) UpdateBookingStatus(ctx context.Context, req *UpdateBookingStatusRequest) (*UpdateBookingStatusResponse, error) {
	log.Printf("Updating booking %s status to %s", req.BookingID, req.Status)
	
	// TODO: Implement actual gRPC call
	return &UpdateBookingStatusResponse{
		Success: true,
		Message: "Booking status updated successfully",
		Booking: Booking{
			ID:     req.BookingID,
			Status: req.Status,
		},
	}, nil
}