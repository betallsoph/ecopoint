package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FirebaseUser struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	Role  string `json:"role,omitempty"`
}

// FirebaseAuth middleware validates Firebase tokens
func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// TODO: Implement Firebase token verification
		// For now, we'll extract user info from a mock token
		// In production, you would verify the token with Firebase Admin SDK
		user, err := verifyFirebaseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Add user to context
		c.Set("user", user)
		c.Next()
	}
}

// verifyFirebaseToken verifies a Firebase token and returns user info
// TODO: Implement actual Firebase token verification
func verifyFirebaseToken(token string) (*FirebaseUser, error) {
	// This is a mock implementation
	// In production, use Firebase Admin SDK to verify the token
	
	// For development, we'll accept a simple format: "uid:email:role"
	if token == "mock-token" {
		return &FirebaseUser{
			UID:   "mock-uid-123",
			Email: "test@example.com",
			Role:  "USER",
		}, nil
	}

	// TODO: Implement real Firebase token verification
	// 1. Get Firebase project ID from config
	// 2. Initialize Firebase Admin SDK
	// 3. Verify the token
	// 4. Extract user claims
	// 5. Return user info

	return nil, nil
}

// GetUserFromContext extracts user from Gin context
func GetUserFromContext(c *gin.Context) *FirebaseUser {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}
	return user.(*FirebaseUser)
}

// RequireRole middleware checks if user has required role
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetUserFromContext(c)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		hasRole := false
		for _, role := range roles {
			if user.Role == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}