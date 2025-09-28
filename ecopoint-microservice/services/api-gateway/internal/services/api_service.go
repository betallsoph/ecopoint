package services

import (
	"context"
	"fmt"
	"log"

	"ecopoint/api-gateway/pkg/grpc"
	"ecopoint/api-gateway/pkg/http"
	"ecopoint/api-gateway/pkg/utils"
	"ecopoint/api-gateway/pkg/validation"
)

type APIService struct {
	userClient    *http.UserClient
	bookingClient *grpc.BookingClient
	validator     *validation.Validator
	retryConfig   *utils.RetryConfig
}

func NewAPIService(userServiceURL, bookingServiceURL string) (*APIService, error) {
	// Initialize User Service HTTP client
	userClient := http.NewUserClient(userServiceURL)

	// Initialize Booking Service gRPC client
	bookingClient, err := grpc.NewBookingClient(bookingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking client: %w", err)
	}

	// Initialize validator
	validator := validation.NewValidator()

	// Initialize retry config
	retryConfig := utils.DefaultRetryConfig()

	return &APIService{
		userClient:    userClient,
		bookingClient: bookingClient,
		validator:     validator,
		retryConfig:   retryConfig,
	}, nil
}

func (s *APIService) Close() error {
	if s.bookingClient != nil {
		return s.bookingClient.Close()
	}
	return nil
}

// User Service methods
func (s *APIService) CreateUser(ctx context.Context, req *http.CreateUserRequest, token string) (*http.User, error) {
	log.Printf("Creating user: %s", req.Email)
	
	// Validate request
	if err := s.validator.ValidateEmail(req.Email); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRequired(req.FirstName, "first_name"); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRequired(req.LastName, "last_name"); err != nil {
		return nil, err
	}
	if req.Phone != "" {
		if err := s.validator.ValidatePhone(req.Phone); err != nil {
			return nil, err
		}
	}
	if err := s.validator.ValidateUserRole(req.Role); err != nil {
		return nil, err
	}

	// Retry with exponential backoff
	return utils.RetryWithResult(ctx, s.retryConfig, func() (*http.User, error) {
		return s.userClient.CreateUser(ctx, req, token)
	})
}

func (s *APIService) GetCurrentUser(ctx context.Context, token string) (*http.User, error) {
	log.Printf("Getting current user")
	return s.userClient.GetCurrentUser(ctx, token)
}

func (s *APIService) UpdateUser(ctx context.Context, req *http.UpdateUserRequest, token string) (*http.User, error) {
	log.Printf("Updating user")
	return s.userClient.UpdateUser(ctx, req, token)
}

func (s *APIService) GetUser(ctx context.Context, userID, token string) (*http.User, error) {
	log.Printf("Getting user: %s", userID)
	return s.userClient.GetUser(ctx, userID, token)
}

// Address methods
func (s *APIService) GetAddresses(ctx context.Context, token string) ([]http.Address, error) {
	log.Printf("Getting user addresses")
	return s.userClient.GetAddresses(ctx, token)
}

func (s *APIService) CreateAddress(ctx context.Context, req *http.CreateAddressRequest, token string) (*http.Address, error) {
	log.Printf("Creating address")
	return s.userClient.CreateAddress(ctx, req, token)
}

func (s *APIService) UpdateAddress(ctx context.Context, addressID string, req *http.UpdateAddressRequest, token string) (*http.Address, error) {
	log.Printf("Updating address: %s", addressID)
	return s.userClient.UpdateAddress(ctx, addressID, req, token)
}

func (s *APIService) DeleteAddress(ctx context.Context, addressID, token string) error {
	log.Printf("Deleting address: %s", addressID)
	return s.userClient.DeleteAddress(ctx, addressID, token)
}

// Booking Service methods
func (s *APIService) CreateBooking(ctx context.Context, req *grpc.CreateBookingRequest) (*grpc.CreateBookingResponse, error) {
	log.Printf("Creating booking for user: %s", req.UserID)
	
	// Validate request
	if err := s.validator.ValidateRequired(req.UserID, "user_id"); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateWasteType(req.WasteType); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRange(req.Quantity, "quantity", 1, 1000); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRequired(req.Description, "description"); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateCoordinates(req.PickupAddress.Latitude, req.PickupAddress.Longitude); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRequired(req.ScheduledDate, "scheduled_date"); err != nil {
		return nil, err
	}
	if err := s.validator.ValidateRequired(req.ScheduledTime, "scheduled_time"); err != nil {
		return nil, err
	}

	// Retry with exponential backoff
	return utils.RetryWithResult(ctx, s.retryConfig, func() (*grpc.CreateBookingResponse, error) {
		return s.bookingClient.CreateBooking(ctx, req)
	})
}

func (s *APIService) GetBooking(ctx context.Context, bookingID string) (*grpc.GetBookingResponse, error) {
	log.Printf("Getting booking: %s", bookingID)
	req := &grpc.GetBookingRequest{
		BookingID: bookingID,
	}
	return s.bookingClient.GetBooking(ctx, req)
}

func (s *APIService) ListBookings(ctx context.Context, userID *string, collectorID *string, status *string, page, limit int) (*grpc.ListBookingsResponse, error) {
	log.Printf("Listing bookings for user: %v", userID)
	req := &grpc.ListBookingsRequest{
		UserID:      userID,
		CollectorID: collectorID,
		Status:      status,
		Page:        page,
		Limit:       limit,
	}
	return s.bookingClient.ListBookings(ctx, req)
}

func (s *APIService) UpdateBooking(ctx context.Context, bookingID string, req *grpc.UpdateBookingRequest) (*grpc.UpdateBookingResponse, error) {
	log.Printf("Updating booking: %s", bookingID)
	req.BookingID = bookingID
	return s.bookingClient.UpdateBooking(ctx, req)
}

func (s *APIService) CancelBooking(ctx context.Context, bookingID, reason string) (*grpc.CancelBookingResponse, error) {
	log.Printf("Cancelling booking: %s", bookingID)
	req := &grpc.CancelBookingRequest{
		BookingID: bookingID,
		Reason:    reason,
	}
	return s.bookingClient.CancelBooking(ctx, req)
}

func (s *APIService) AssignCollector(ctx context.Context, bookingID, collectorID string) (*grpc.AssignCollectorResponse, error) {
	log.Printf("Assigning collector %s to booking %s", collectorID, bookingID)
	req := &grpc.AssignCollectorRequest{
		BookingID:   bookingID,
		CollectorID: collectorID,
	}
	return s.bookingClient.AssignCollector(ctx, req)
}

func (s *APIService) UpdateBookingStatus(ctx context.Context, bookingID, status, notes string) (*grpc.UpdateBookingStatusResponse, error) {
	log.Printf("Updating booking %s status to %s", bookingID, status)
	req := &grpc.UpdateBookingStatusRequest{
		BookingID: bookingID,
		Status:    status,
		Notes:     notes,
	}
	return s.bookingClient.UpdateBookingStatus(ctx, req)
}