package handlers

import (
	"context"
	"net/http"
	"time"

	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/shared/models"
	"ecopoint-backend/auth-service/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	config *config.Config
	userDB *mongo.Collection
}

type GoogleSignInRequest struct {
	IDToken  string           `json:"id_token" binding:"required"`
	UserType models.UserType  `json:"user_type" binding:"required"`
}

type PhoneVerifyRequest struct {
	PhoneNumber string          `json:"phone_number" binding:"required"`
	UserType    models.UserType `json:"user_type" binding:"required"`
	DisplayName string          `json:"display_name" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	db := database.GetDatabase("ecopoint")
	return &AuthHandler{
		config: cfg,
		userDB: db.Collection("users"),
	}
}

// GoogleSignIn handles Google authentication
func (h *AuthHandler) GoogleSignIn(c *gin.Context) {
	var req GoogleSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify Google ID token with Firebase
	userInfo, err := utils.VerifyGoogleIDToken(req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
		return
	}

	// Check if user exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser models.User
	err = h.userDB.FindOne(ctx, bson.M{"firebase_uid": userInfo.UID}).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		// Create new user
		now := time.Now()
		newUser := models.User{
			ID:          primitive.NewObjectID(),
			FirebaseUID: userInfo.UID,
			Email:       userInfo.Email,
			DisplayName: userInfo.Name,
			PhotoURL:    userInfo.Picture,
			UserType:    req.UserType,
			IsActive:    true,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		_, err = h.userDB.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		existingUser = newUser
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	} else {
		// Update last login
		update := bson.M{
			"$set": bson.M{
				"updated_at": time.Now(),
			},
		}
		_, err = h.userDB.UpdateOne(ctx, bson.M{"_id": existingUser.ID}, update)
		if err != nil {
			// Log error but don't fail the request
		}
	}

	// Generate JWT tokens
	accessToken, err := utils.GenerateAccessToken(existingUser.ID.Hex(), existingUser.UserType, h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(existingUser.ID.Hex(), h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	response := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         existingUser,
	}

	c.JSON(http.StatusOK, response)
}

// PhoneVerify handles phone number verification
func (h *AuthHandler) PhoneVerify(c *gin.Context) {
	var req PhoneVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For MVP, we'll create user directly with phone number
	// In production, you'd verify the phone number with Firebase Auth first

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user exists by phone number
	var existingUser models.User
	err := h.userDB.FindOne(ctx, bson.M{"phone_number": req.PhoneNumber}).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		// Create new user
		now := time.Now()
		newUser := models.User{
			ID:          primitive.NewObjectID(),
			PhoneNumber: req.PhoneNumber,
			DisplayName: req.DisplayName,
			UserType:    req.UserType,
			IsActive:    true,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		_, err = h.userDB.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		existingUser = newUser
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Generate JWT tokens
	accessToken, err := utils.GenerateAccessToken(existingUser.ID.Hex(), existingUser.UserType, h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(existingUser.ID.Hex(), h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	response := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         existingUser,
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := utils.ValidateRefreshToken(req.RefreshToken, h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Get user to generate new access token
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex(claims.UserID)
	var user models.User
	err = h.userDB.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Generate new access token
	accessToken, err := utils.GenerateAccessToken(user.ID.Hex(), user.UserType, h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

// GetProfile returns the current user's profile
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(userID)
	var user models.User
	err := h.userDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}