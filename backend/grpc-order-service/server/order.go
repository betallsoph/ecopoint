package server

import (
	"context"
	"time"

	"ecopoint-backend/shared/config"
	"ecopoint-backend/shared/database"
	"ecopoint-backend/shared/models"
	pb "ecopoint-backend/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	config  *config.Config
	orderDB *mongo.Collection
}

func NewOrderServer(cfg *config.Config) *OrderServer {
	db := database.GetDatabase("ecopoint")
	return &OrderServer{
		config:  cfg,
		orderDB: db.Collection("orders"),
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	customerID, err := primitive.ObjectIDFromHex(req.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid customer ID: %v", err)
	}

	now := time.Now()

	// Convert protobuf waste types to models
	wasteTypes := make([]models.WasteType, len(req.WasteTypes))
	for i, wt := range req.WasteTypes {
		switch wt {
		case pb.WasteType_WASTE_TYPE_PAPER:
			wasteTypes[i] = models.WasteTypePaper
		case pb.WasteType_WASTE_TYPE_PLASTIC:
			wasteTypes[i] = models.WasteTypePlastic
		case pb.WasteType_WASTE_TYPE_METAL:
			wasteTypes[i] = models.WasteTypeMetal
		case pb.WasteType_WASTE_TYPE_GLASS:
			wasteTypes[i] = models.WasteTypeGlass
		case pb.WasteType_WASTE_TYPE_ELECTRONIC:
			wasteTypes[i] = models.WasteTypeElectronic
		}
	}

	// Calculate payment amount
	amount := calculatePayment(wasteTypes, req.EstimatedWeight)

	order := models.Order{
		ID:              primitive.NewObjectID(),
		CustomerID:      customerID,
		Status:          models.OrderStatusPending,
		WasteTypes:      wasteTypes,
		EstimatedWeight: req.EstimatedWeight,
		PickupAddress: models.Address{
			Street:   req.PickupAddress.Street,
			District: req.PickupAddress.District,
			City:     req.PickupAddress.City,
			Lat:      req.PickupAddress.Lat,
			Lng:      req.PickupAddress.Lng,
		},
		ScheduledTime: req.ScheduledTime.AsTime(),
		Notes:         req.Notes,
		Payment: models.Payment{
			Amount:   amount,
			Currency: "VND",
			Method:   "cash",
			IsPaid:   false,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := s.orderDB.InsertOne(ctx, order)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create order: %v", err)
	}

	order.ID = result.InsertedID.(primitive.ObjectID)

	// Convert to protobuf
	pbOrder := convertOrderToProto(&order)

	return &pb.CreateOrderResponse{
		Order: pbOrder,
	}, nil
}

func (s *OrderServer) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	var filter bson.M
	if req.UserType == "customer" {
		filter = bson.M{"customer_id": userID}
	} else if req.UserType == "collector" {
		filter = bson.M{"collector_id": userID}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user type")
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.orderDB.Find(ctx, filter, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch orders: %v", err)
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode orders: %v", err)
	}

	// Convert to protobuf
	pbOrders := make([]*pb.Order, len(orders))
	for i, order := range orders {
		pbOrders[i] = convertOrderToProto(&order)
	}

	return &pb.GetUserOrdersResponse{
		Orders: pbOrders,
	}, nil
}

func (s *OrderServer) GetAvailableOrders(ctx context.Context, req *pb.GetAvailableOrdersRequest) (*pb.GetAvailableOrdersResponse, error) {
	filter := bson.M{
		"status":       models.OrderStatusPending,
		"collector_id": bson.M{"$exists": false},
	}

	opts := options.Find().SetSort(bson.D{{Key: "scheduled_time", Value: 1}})
	cursor, err := s.orderDB.Find(ctx, filter, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch available orders: %v", err)
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode orders: %v", err)
	}

	// Convert to protobuf
	pbOrders := make([]*pb.Order, len(orders))
	for i, order := range orders {
		pbOrders[i] = convertOrderToProto(&order)
	}

	return &pb.GetAvailableOrdersResponse{
		Orders: pbOrders,
	}, nil
}

func (s *OrderServer) AcceptOrder(ctx context.Context, req *pb.AcceptOrderRequest) (*pb.AcceptOrderResponse, error) {
	orderID, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order ID: %v", err)
	}

	collectorID, err := primitive.ObjectIDFromHex(req.CollectorId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collector ID: %v", err)
	}

	filter := bson.M{
		"_id":          orderID,
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

	result, err := s.orderDB.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to accept order: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Order not available or already accepted")
	}

	return &pb.AcceptOrderResponse{
		Success: true,
		Message: "Order accepted successfully",
	}, nil
}

func (s *OrderServer) CompleteOrder(ctx context.Context, req *pb.CompleteOrderRequest) (*pb.CompleteOrderResponse, error) {
	orderID, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order ID: %v", err)
	}

	collectorID, err := primitive.ObjectIDFromHex(req.CollectorId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid collector ID: %v", err)
	}

	now := time.Now()
	filter := bson.M{
		"_id":          orderID,
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

	result, err := s.orderDB.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to complete order: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Order not found or cannot be completed")
	}

	return &pb.CompleteOrderResponse{
		Success: true,
		Message: "Order completed successfully",
	}, nil
}

// Helper functions
func convertOrderToProto(order *models.Order) *pb.Order {
	pbOrder := &pb.Order{
		Id:              order.ID.Hex(),
		CustomerId:      order.CustomerID.Hex(),
		Status:          convertOrderStatusToProto(order.Status),
		EstimatedWeight: order.EstimatedWeight,
		PickupAddress: &pb.Address{
			Street:   order.PickupAddress.Street,
			District: order.PickupAddress.District,
			City:     order.PickupAddress.City,
			Lat:      order.PickupAddress.Lat,
			Lng:      order.PickupAddress.Lng,
		},
		ScheduledTime: timestamppb.New(order.ScheduledTime),
		Notes:         order.Notes,
		Payment: &pb.Payment{
			Amount:   order.Payment.Amount,
			Currency: order.Payment.Currency,
			Method:   order.Payment.Method,
			IsPaid:   order.Payment.IsPaid,
		},
		CreatedAt: timestamppb.New(order.CreatedAt),
		UpdatedAt: timestamppb.New(order.UpdatedAt),
	}

	// Convert collector ID if present
	if order.CollectorID != nil {
		pbOrder.CollectorId = order.CollectorID.Hex()
	}

	// Convert actual weight if present
	if order.ActualWeight != nil {
		pbOrder.ActualWeight = *order.ActualWeight
	}

	// Convert completed time if present
	if order.CompletedTime != nil {
		pbOrder.CompletedTime = timestamppb.New(*order.CompletedTime)
	}

	// Convert waste types
	for _, wt := range order.WasteTypes {
		pbOrder.WasteTypes = append(pbOrder.WasteTypes, convertWasteTypeToProto(wt))
	}

	// Convert payment time if present
	if order.Payment.PaymentTime != nil {
		pbOrder.Payment.PaymentTime = timestamppb.New(*order.Payment.PaymentTime)
	}

	return pbOrder
}

func convertOrderStatusToProto(status models.OrderStatus) pb.OrderStatus {
	switch status {
	case models.OrderStatusPending:
		return pb.OrderStatus_ORDER_STATUS_PENDING
	case models.OrderStatusAccepted:
		return pb.OrderStatus_ORDER_STATUS_ACCEPTED
	case models.OrderStatusInProgress:
		return pb.OrderStatus_ORDER_STATUS_IN_PROGRESS
	case models.OrderStatusCompleted:
		return pb.OrderStatus_ORDER_STATUS_COMPLETED
	case models.OrderStatusCancelled:
		return pb.OrderStatus_ORDER_STATUS_CANCELLED
	default:
		return pb.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}

func convertWasteTypeToProto(wt models.WasteType) pb.WasteType {
	switch wt {
	case models.WasteTypePaper:
		return pb.WasteType_WASTE_TYPE_PAPER
	case models.WasteTypePlastic:
		return pb.WasteType_WASTE_TYPE_PLASTIC
	case models.WasteTypeMetal:
		return pb.WasteType_WASTE_TYPE_METAL
	case models.WasteTypeGlass:
		return pb.WasteType_WASTE_TYPE_GLASS
	case models.WasteTypeElectronic:
		return pb.WasteType_WASTE_TYPE_ELECTRONIC
	default:
		return pb.WasteType_WASTE_TYPE_UNSPECIFIED
	}
}

func calculatePayment(wasteTypes []models.WasteType, weight float64) float64 {
	basePrice := 10000.0 // 10,000 VND per kg base price
	multiplier := 1.0

	for _, wasteType := range wasteTypes {
		switch wasteType {
		case models.WasteTypeMetal:
			multiplier += 0.5
		case models.WasteTypeElectronic:
			multiplier += 0.8
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