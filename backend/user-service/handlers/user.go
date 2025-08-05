package handlers

import (
	"context"
	"net/http"
	"time"

	"ecopoint-backend/shared/database"
	"ecopoint-backend/shared/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userDB  *mongo.Collection
	orderDB *mongo.Collection
}

type UpdateProfileRequest struct {
	DisplayName *string         `json:"display_name,omitempty"`
	PhoneNumber *string         `json:"phone_number,omitempty"`
	Address     *models.Address `json:"address,omitempty"`
	Vehicle     *models.Vehicle `json:"vehicle,omitempty"`
}

type UserStats struct {
	TotalOrders     int64   `json:"total_orders"`
	CompletedOrders int64   `json:"completed_orders"`
	Rating          float64 `json:"rating,omitempty"`
	TotalEarnings   float64 `json:"total_earnings,omitempty"`
	IsOnline        bool    `json:"is_online,omitempty"`
}

func NewUserHandler() *UserHandler {
	db := database.GetDatabase("ecopoint")
	return &UserHandler{
		userDB:  db.Collection("users"),
		orderDB: db.Collection("orders"),
	}
}

// GetProfile returns the current user's profile
func (h *UserHandler) GetProfile(c *gin.Context) {
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

// UpdateProfile updates the current user's profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(userID)

	// Build update document
	updateDoc := bson.M{
		"updated_at": time.Now(),
	}

	if req.DisplayName != nil {
		updateDoc["display_name"] = *req.DisplayName
	}
	if req.PhoneNumber != nil {
		updateDoc["phone_number"] = *req.PhoneNumber
	}
	if req.Address != nil {
		updateDoc["address"] = *req.Address
	}
	if req.Vehicle != nil {
		updateDoc["vehicle"] = *req.Vehicle
	}

	_, err := h.userDB.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updateDoc})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// GetUserStats returns statistics for the current user
func (h *UserHandler) GetUserStats(c *gin.Context) {
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(userID)
	
	var stats UserStats

	if userType == string(models.UserTypeCustomer) {
		// Customer stats
		totalCount, _ := h.orderDB.CountDocuments(ctx, bson.M{"customer_id": objectID})
		completedCount, _ := h.orderDB.CountDocuments(ctx, bson.M{
			"customer_id": objectID,
			"status":      models.OrderStatusCompleted,
		})
		
		stats.TotalOrders = totalCount
		stats.CompletedOrders = completedCount
		
	} else if userType == string(models.UserTypeCollector) {
		// Collector stats
		totalCount, _ := h.orderDB.CountDocuments(ctx, bson.M{"collector_id": objectID})
		completedCount, _ := h.orderDB.CountDocuments(ctx, bson.M{
			"collector_id": objectID,
			"status":       models.OrderStatusCompleted,
		})
		
		// Calculate total earnings
		pipeline := []bson.M{
			{"$match": bson.M{
				"collector_id": objectID,
				"status":       models.OrderStatusCompleted,
			}},
			{"$group": bson.M{
				"_id":   nil,
				"total": bson.M{"$sum": "$payment.amount"},
			}},
		}
		
		cursor, err := h.orderDB.Aggregate(ctx, pipeline)
		if err == nil {
			var result []bson.M
			cursor.All(ctx, &result)
			if len(result) > 0 {
				if total, ok := result[0]["total"].(float64); ok {
					stats.TotalEarnings = total
				}
			}
		}
		
		// Get collector info for rating and online status
		var collector models.Collector
		err = h.userDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&collector)
		if err == nil {
			stats.Rating = collector.Rating
			stats.IsOnline = collector.IsOnline
		}
		
		stats.TotalOrders = totalCount
		stats.CompletedOrders = completedCount
	}

	c.JSON(http.StatusOK, stats)
}