package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"ecopoint-backend/graphql-service/auth"
	"ecopoint-backend/graphql-service/utils"

	"github.com/gin-gonic/gin"
)

// AuthContextKey is the key for storing auth info in context
type AuthContextKey string

const (
	UserContextKey AuthContextKey = "user"
	UIDContextKey  AuthContextKey = "uid"
)

// AuthenticatedUser contains authenticated user information
type AuthenticatedUser struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Verified bool   `json:"verified"`
}

// AuthMiddleware validates Firebase ID tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer TOKEN"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		
		// Try JWT access token first
		claims, jwtErr := utils.ValidateAccessToken(token)
		if jwtErr == nil {
			// JWT access token - get user info from token claims
			user := &AuthenticatedUser{
				UID:      claims.UserID,
				Email:    "", // Not stored in JWT
				Name:     "", // Not stored in JWT  
				Picture:  "", // Not stored in JWT
				Verified: true, // JWT tokens are considered verified
			}
			
			// Store in contexts
			c.Set(string(UserContextKey), user)
			c.Set(string(UIDContextKey), claims.UserID)
			ctx := context.WithValue(c.Request.Context(), UserContextKey, user)
			ctx = context.WithValue(ctx, UIDContextKey, claims.UserID)
			c.Request = c.Request.WithContext(ctx)
			
			c.Next()
			return
		}
		
		// Try Firebase ID token as fallback
		userInfo, err := auth.VerifyIDToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Firebase ID token - get user info from Firebase
		user := &AuthenticatedUser{
			UID:      userInfo.UID,
			Email:    userInfo.Email,
			Name:     userInfo.Name,
			Picture:  userInfo.Picture,
			Verified: userInfo.Verified,
		}

		// Store in Gin context
		c.Set(string(UserContextKey), user)
		c.Set(string(UIDContextKey), userInfo.UID)

		// Store in request context for GraphQL resolvers
		ctx := context.WithValue(c.Request.Context(), UserContextKey, user)
		ctx = context.WithValue(ctx, UIDContextKey, userInfo.UID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// GetUserFromContext extracts authenticated user from context
func GetUserFromContext(ctx context.Context) (*AuthenticatedUser, error) {
	user, ok := ctx.Value(UserContextKey).(*AuthenticatedUser)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

// GetUIDFromContext extracts user UID from context
func GetUIDFromContext(ctx context.Context) (string, error) {
	uid, ok := ctx.Value(UIDContextKey).(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in context")
	}
	return uid, nil
}

// RequireAuth ensures user is authenticated (for GraphQL resolvers)
func RequireAuth(ctx context.Context) (*AuthenticatedUser, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("authentication required: %v", err)
	}
	return user, nil
}