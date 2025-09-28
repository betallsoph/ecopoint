package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ecopoint/api-gateway/internal/middleware"
	"ecopoint/api-gateway/internal/services"
	httpClient "ecopoint/api-gateway/pkg/http"
)

var apiService *services.APIService

func SetAPIService(service *services.APIService) {
	apiService = service
}

// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	var req httpClient.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get token from header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	user, err := apiService.CreateUser(c.Request.Context(), &req, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
		"user":    user,
	})
}

// GetCurrentUser handles getting current user info
func GetCurrentUser(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Get token from header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	userDetails, err := apiService.GetCurrentUser(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    userDetails,
	})
}

// UpdateUser handles user profile updates
func UpdateUser(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// TODO: Implement user update
	// 1. Validate request body
	// 2. Call user service
	// 3. Return updated user
	
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser endpoint - TODO: Implement",
		"user_id": user.UID,
	})
}

// GetUser handles getting user by ID
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	
	// TODO: Implement get user by ID
	// 1. Call user service
	// 2. Return user details
	
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser endpoint - TODO: Implement",
		"user_id": userID,
	})
}