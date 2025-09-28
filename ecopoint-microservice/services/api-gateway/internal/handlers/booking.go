package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ecopoint/api-gateway/internal/middleware"
	"ecopoint/api-gateway/pkg/grpc"
)

// CreateBooking handles booking creation
func CreateBooking(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var req grpc.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set user ID from context
	req.UserID = user.UID

	resp, err := apiService.CreateBooking(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": resp.Success,
		"message": resp.Message,
		"booking": resp.Booking,
	})
}

// GetBookings handles getting user bookings
func GetBookings(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Parse query parameters
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	resp, err := apiService.ListBookings(c.Request.Context(), &user.UID, nil, statusPtr, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  resp.Success,
		"message":  resp.Message,
		"bookings": resp.Bookings,
		"total":    resp.Total,
		"page":     resp.Page,
		"limit":    resp.Limit,
	})
}

// GetBooking handles getting booking by ID
func GetBooking(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	bookingID := c.Param("id")
	
	// TODO: Implement get booking by ID
	// 1. Call booking service via gRPC
	// 2. Return booking details
	
	c.JSON(http.StatusOK, gin.H{
		"message":    "GetBooking endpoint - TODO: Implement",
		"user_id":    user.UID,
		"booking_id": bookingID,
	})
}

// UpdateBooking handles booking updates
func UpdateBooking(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	bookingID := c.Param("id")
	
	// TODO: Implement booking update
	// 1. Validate request body
	// 2. Call booking service via gRPC
	// 3. Return updated booking
	
	c.JSON(http.StatusOK, gin.H{
		"message":    "UpdateBooking endpoint - TODO: Implement",
		"user_id":    user.UID,
		"booking_id": bookingID,
	})
}

// CancelBooking handles booking cancellation
func CancelBooking(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	bookingID := c.Param("id")
	
	// TODO: Implement booking cancellation
	// 1. Call booking service via gRPC
	// 2. Return success response
	
	c.JSON(http.StatusOK, gin.H{
		"message":    "CancelBooking endpoint - TODO: Implement",
		"user_id":    user.UID,
		"booking_id": bookingID,
	})
}

// AssignCollector handles collector assignment
func AssignCollector(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	bookingID := c.Param("id")
	
	// TODO: Implement collector assignment
	// 1. Validate request body (collector_id)
	// 2. Call booking service via gRPC
	// 3. Return updated booking
	
	c.JSON(http.StatusOK, gin.H{
		"message":    "AssignCollector endpoint - TODO: Implement",
		"user_id":    user.UID,
		"booking_id": bookingID,
	})
}

// UpdateBookingStatus handles booking status updates
func UpdateBookingStatus(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	bookingID := c.Param("id")
	
	// TODO: Implement booking status update
	// 1. Validate request body (status, notes)
	// 2. Call booking service via gRPC
	// 3. Return updated booking
	
	c.JSON(http.StatusOK, gin.H{
		"message":    "UpdateBookingStatus endpoint - TODO: Implement",
		"user_id":    user.UID,
		"booking_id": bookingID,
	})
}