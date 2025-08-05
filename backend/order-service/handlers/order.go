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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderHandler struct {
	orderDB *mongo.Collection
	userDB  *mongo.Collection
}

func NewOrderHandler() *OrderHandler {
	db := database.GetDatabase("ecopoint")
	return &OrderHandler{
		orderDB: db.Collection("orders"),
		userDB:  db.Collection("users"),
	}
}

// CreateOrder creates a new waste collection order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	// Only customers can create orders
	if userType != string(models.UserTypeCustomer) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only customers can create orders"})
		return
	}

	var req models.OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customerID, _ := primitive.ObjectIDFromHex(userID)
	now := time.Now()

	// Calculate payment amount based on waste types and weight
	amount := calculatePayment(req.WasteTypes, req.EstimatedWeight)

	order := models.Order{
		ID:              primitive.NewObjectID(),
		CustomerID:      customerID,
		Status:          models.OrderStatusPending,
		WasteTypes:      req.WasteTypes,
		EstimatedWeight: req.EstimatedWeight,
		PickupAddress:   req.PickupAddress,
		ScheduledTime:   req.ScheduledTime,
		Notes:           req.Notes,
		Payment: models.Payment{
			Amount:   amount,
			Currency: "VND",
			Method:   "cash",
			IsPaid:   false,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := h.orderDB.InsertOne(ctx, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	order.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, order)
}

// GetOrders gets orders for the current user
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var filter bson.M

	if userType == string(models.UserTypeCustomer) {
		customerID, _ := primitive.ObjectIDFromHex(userID)
		filter = bson.M{"customer_id": customerID}
	} else if userType == string(models.UserTypeCollector) {
		collectorID, _ := primitive.ObjectIDFromHex(userID)
		filter = bson.M{"collector_id": collectorID}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid user type"})
		return
	}

	// Sort by created_at descending
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := h.orderDB.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrderByID gets a specific order by ID
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	err = h.orderDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check if user has permission to view this order
	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	if userType == string(models.UserTypeCustomer) && order.CustomerID != userObjectID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to view this order"})
		return
	} else if userType == string(models.UserTypeCollector) && 
		(order.CollectorID == nil || *order.CollectorID != userObjectID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to view this order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAvailableOrders gets orders available for collectors to pick up
func (h *OrderHandler) GetAvailableOrders(c *gin.Context) {
	userType := c.GetString("user_type")

	// Only collectors can view available orders
	if userType != string(models.UserTypeCollector) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only collectors can view available orders"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find orders that are pending and not assigned to any collector
	filter := bson.M{
		"status":       models.OrderStatusPending,
		"collector_id": bson.M{"$exists": false},
	}

	// Sort by scheduled time (earliest first)
	opts := options.Find().SetSort(bson.D{{Key: "scheduled_time", Value: 1}})
	cursor, err := h.orderDB.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available orders"})
		return
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// AcceptOrder allows a collector to accept an order
func (h *OrderHandler) AcceptOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	// Only collectors can accept orders
	if userType != string(models.UserTypeCollector) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only collectors can accept orders"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	collectorID, _ := primitive.ObjectIDFromHex(userID)

	// Update order to assign collector and change status
	filter := bson.M{
		"_id":          objectID,
		"status":       models.OrderStatusPending,
		"collector_id": bson.M{"$exists": false},
	}

	update := bson.M{
		"$set": bson.M{
			"collector_id": collectorID,
			"status":       models.OrderStatusAccepted,
			"updated_at":   time.Now(),
		},
	}

	result, err := h.orderDB.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept order"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Order not available or already accepted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order accepted successfully"})
}

// UpdateOrder updates order status or details
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	var req models.OrderUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Get the order first to check permissions
	var order models.Order
	err = h.orderDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check permissions
	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	if userType == string(models.UserTypeCustomer) && order.CustomerID != userObjectID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this order"})
		return
	} else if userType == string(models.UserTypeCollector) && 
		(order.CollectorID == nil || *order.CollectorID != userObjectID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this order"})
		return
	}

	// Build update document
	updateDoc := bson.M{
		"updated_at": time.Now(),
	}

	if req.Status != nil {
		updateDoc["status"] = *req.Status
	}
	if req.ActualWeight != nil {
		updateDoc["actual_weight"] = *req.ActualWeight
	}
	if req.Notes != nil {
		updateDoc["notes"] = *req.Notes
	}

	_, err = h.orderDB.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updateDoc})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// CompleteOrder marks an order as completed
func (h *OrderHandler) CompleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetString("user_id")
	userType := c.GetString("user_type")

	// Only collectors can complete orders
	if userType != string(models.UserTypeCollector) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only collectors can complete orders"})
		return
	}

	var req struct {
		ActualWeight float64 `json:"actual_weight" binding:"required,min=0.1"`
		Notes        string  `json:"notes,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	collectorID, _ := primitive.ObjectIDFromHex(userID)
	now := time.Now()

	// Update order to completed status
	filter := bson.M{
		"_id":          objectID,
		"collector_id": collectorID,
		"status":       bson.M{"$in": []models.OrderStatus{models.OrderStatusAccepted, models.OrderStatusInProgress}},
	}

	update := bson.M{
		"$set": bson.M{
			"status":         models.OrderStatusCompleted,
			"actual_weight":  req.ActualWeight,
			"completed_time": now,
			"updated_at":     now,
		},
	}

	if req.Notes != "" {
		update["$set"].(bson.M)["notes"] = req.Notes
	}

	result, err := h.orderDB.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete order"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or cannot be completed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order completed successfully"})
}

// calculatePayment calculates payment amount based on waste types and weight
func calculatePayment(wasteTypes []models.WasteType, weight float64) float64 {
	// Simple pricing calculation
	// In reality, this would be more complex based on market rates
	basePrice := 10000.0 // 10,000 VND per kg base price

	multiplier := 1.0
	for _, wasteType := range wasteTypes {
		switch wasteType {
		case models.WasteTypeMetal:
			multiplier += 0.5 // Metal is more valuable
		case models.WasteTypeElectronic:
			multiplier += 0.8 // Electronics are highly valuable
		case models.WasteTypePlastic:
			multiplier += 0.2
		case models.WasteTypePaper:
			multiplier += 0.1
		case models.WasteTypeGlass:
			multiplier += 0.3
		}
	}

	return basePrice * weight * multiplier
}