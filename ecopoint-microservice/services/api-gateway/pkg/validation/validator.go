package validation

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// Validator provides validation methods
type Validator struct{}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateEmail validates an email address
func (v *Validator) ValidateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Message: "email is required"}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &ValidationError{Field: "email", Message: "invalid email format"}
	}

	return nil
}

// ValidatePhone validates a phone number
func (v *Validator) ValidatePhone(phone string) error {
	if phone == "" {
		return &ValidationError{Field: "phone", Message: "phone is required"}
	}

	// Simple phone validation (adjust regex as needed)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(phone) {
		return &ValidationError{Field: "phone", Message: "invalid phone format"}
	}

	return nil
}

// ValidateRequired validates that a field is not empty
func (v *Validator) ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s is required", fieldName)}
	}
	return nil
}

// ValidateLength validates string length
func (v *Validator) ValidateLength(value, fieldName string, min, max int) error {
	length := len(strings.TrimSpace(value))
	if length < min {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be at least %d characters", fieldName, min)}
	}
	if length > max {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be at most %d characters", fieldName, max)}
	}
	return nil
}

// ValidateRange validates numeric range
func (v *Validator) ValidateRange(value int, fieldName string, min, max int) error {
	if value < min {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be at least %d", fieldName, min)}
	}
	if value > max {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be at most %d", fieldName, max)}
	}
	return nil
}

// ValidateCoordinates validates latitude and longitude
func (v *Validator) ValidateCoordinates(lat, lng float64) error {
	if lat < -90 || lat > 90 {
		return &ValidationError{Field: "latitude", Message: "latitude must be between -90 and 90"}
	}
	if lng < -180 || lng > 180 {
		return &ValidationError{Field: "longitude", Message: "longitude must be between -180 and 180"}
	}
	return nil
}

// ValidateWasteType validates waste type
func (v *Validator) ValidateWasteType(wasteType string) error {
	validTypes := []string{"PLASTIC", "PAPER", "GLASS", "METAL", "ORGANIC", "ELECTRONIC", "OTHER"}
	for _, validType := range validTypes {
		if wasteType == validType {
			return nil
		}
	}
	return &ValidationError{Field: "waste_type", Message: "invalid waste type"}
}

// ValidateBookingStatus validates booking status
func (v *Validator) ValidateBookingStatus(status string) error {
	validStatuses := []string{"PENDING", "CONFIRMED", "ASSIGNED", "IN_PROGRESS", "COMPLETED", "CANCELLED", "FAILED"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}
	return &ValidationError{Field: "status", Message: "invalid booking status"}
}

// ValidateUserRole validates user role
func (v *Validator) ValidateUserRole(role string) error {
	validRoles := []string{"USER", "COLLECTOR", "ADMIN"}
	for _, validRole := range validRoles {
		if role == validRole {
			return nil
		}
	}
	return &ValidationError{Field: "role", Message: "invalid user role"}
}

// ValidatePagination validates pagination parameters
func (v *Validator) ValidatePagination(page, limit int) error {
	if page < 1 {
		return &ValidationError{Field: "page", Message: "page must be at least 1"}
	}
	if limit < 1 || limit > 100 {
		return &ValidationError{Field: "limit", Message: "limit must be between 1 and 100"}
	}
	return nil
}