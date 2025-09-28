package models

import (
	"time"
)

// BookingStatus represents the status of a booking
type BookingStatus string

const (
	StatusPending     BookingStatus = "PENDING"
	StatusConfirmed   BookingStatus = "CONFIRMED"
	StatusAssigned    BookingStatus = "ASSIGNED"
	StatusInProgress  BookingStatus = "IN_PROGRESS"
	StatusCompleted   BookingStatus = "COMPLETED"
	StatusCancelled   BookingStatus = "CANCELLED"
	StatusFailed      BookingStatus = "FAILED"
)

// WasteType represents the type of waste
type WasteType string

const (
	WasteTypePlastic   WasteType = "PLASTIC"
	WasteTypePaper     WasteType = "PAPER"
	WasteTypeGlass     WasteType = "GLASS"
	WasteTypeMetal     WasteType = "METAL"
	WasteTypeOrganic   WasteType = "ORGANIC"
	WasteTypeElectronic WasteType = "ELECTRONIC"
	WasteTypeOther     WasteType = "OTHER"
)

// Address represents a physical address
type Address struct {
	Street    string  `json:"street" db:"street"`
	City      string  `json:"city" db:"city"`
	State     string  `json:"state" db:"state"`
	ZipCode   string  `json:"zip_code" db:"zip_code"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

// Booking represents a waste collection booking
type Booking struct {
	ID             string       `json:"id" db:"id"`
	UserID         string       `json:"user_id" db:"user_id"`
	CollectorID    *string      `json:"collector_id" db:"collector_id"`
	WasteType      WasteType    `json:"waste_type" db:"waste_type"`
	Quantity       int          `json:"quantity" db:"quantity"`
	Description    string       `json:"description" db:"description"`
	PickupAddress  Address      `json:"pickup_address" db:"pickup_address"`
	ScheduledDate  time.Time    `json:"scheduled_date" db:"scheduled_date"`
	ScheduledTime  string       `json:"scheduled_time" db:"scheduled_time"`
	Status         BookingStatus `json:"status" db:"status"`
	EstimatedPrice float64      `json:"estimated_price" db:"estimated_price"`
	FinalPrice     *float64     `json:"final_price" db:"final_price"`
	Notes          string       `json:"notes" db:"notes"`
	Images         []string     `json:"images" db:"images"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
}

// CreateBookingRequest represents the request to create a booking
type CreateBookingRequest struct {
	UserID         string    `json:"user_id" validate:"required"`
	WasteType      WasteType `json:"waste_type" validate:"required"`
	Quantity       int       `json:"quantity" validate:"required,min=1"`
	Description    string    `json:"description" validate:"required"`
	PickupAddress  Address   `json:"pickup_address" validate:"required"`
	ScheduledDate  string    `json:"scheduled_date" validate:"required"`
	ScheduledTime  string    `json:"scheduled_time" validate:"required"`
	Notes          string    `json:"notes"`
	Images         []string  `json:"images"`
}

// UpdateBookingRequest represents the request to update a booking
type UpdateBookingRequest struct {
	ID             string     `json:"id" validate:"required"`
	WasteType      *WasteType `json:"waste_type,omitempty"`
	Quantity       *int       `json:"quantity,omitempty"`
	Description    *string    `json:"description,omitempty"`
	PickupAddress  *Address   `json:"pickup_address,omitempty"`
	ScheduledDate  *string    `json:"scheduled_date,omitempty"`
	ScheduledTime  *string    `json:"scheduled_time,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
	Images         []string   `json:"images,omitempty"`
}

// AssignCollectorRequest represents the request to assign a collector
type AssignCollectorRequest struct {
	BookingID   string `json:"booking_id" validate:"required"`
	CollectorID string `json:"collector_id" validate:"required"`
}

// UpdateStatusRequest represents the request to update booking status
type UpdateStatusRequest struct {
	BookingID string         `json:"booking_id" validate:"required"`
	Status    BookingStatus  `json:"status" validate:"required"`
	Notes     string         `json:"notes"`
}

// CancelBookingRequest represents the request to cancel a booking
type CancelBookingRequest struct {
	BookingID string `json:"booking_id" validate:"required"`
	Reason    string `json:"reason"`
}

// ListBookingsRequest represents the request to list bookings
type ListBookingsRequest struct {
	UserID      *string        `json:"user_id,omitempty"`
	CollectorID *string        `json:"collector_id,omitempty"`
	Status      *BookingStatus `json:"status,omitempty"`
	Page        int            `json:"page"`
	Limit       int            `json:"limit"`
}

// ListBookingsResponse represents the response for listing bookings
type ListBookingsResponse struct {
	Bookings []Booking `json:"bookings"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
}