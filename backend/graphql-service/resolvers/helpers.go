package resolvers

import (
	"ecopoint-backend/shared/models"
	pb_user "ecopoint-backend/proto"
	pb_order "ecopoint-backend/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// convertProtoUserToModel converts protobuf User to models.User
func convertProtoUserToModel(pbUser *pb_user.User) *models.User {
	if pbUser == nil {
		return nil
	}

	objectID, _ := primitive.ObjectIDFromHex(pbUser.Id)
	
	user := &models.User{
		ID:          objectID,
		FirebaseUID: pbUser.FirebaseUid,
		Email:       pbUser.Email,
		DisplayName: pbUser.DisplayName,
		PhotoURL:    pbUser.PhotoUrl,
		PhoneNumber: pbUser.PhoneNumber,
		UserType:    models.UserType(pbUser.UserType.String()),
		IsActive:    pbUser.IsActive,
		CreatedAt:   pbUser.CreatedAt.AsTime(),
		UpdatedAt:   pbUser.UpdatedAt.AsTime(),
	}

	// Note: Address and Vehicle are not part of base User model in current setup
	// They would be handled separately or as part of extended user types

	return user
}

// convertProtoOrderToModel converts protobuf Order to models.Order
func convertProtoOrderToModel(pbOrder *pb_order.Order) *models.Order {
	if pbOrder == nil {
		return nil
	}

	customerID, _ := primitive.ObjectIDFromHex(pbOrder.CustomerId)
	objectID, _ := primitive.ObjectIDFromHex(pbOrder.Id)
	
	order := &models.Order{
		ID:              objectID,
		CustomerID:      customerID,
		Status:          models.OrderStatus(pbOrder.Status.String()),
		EstimatedWeight: pbOrder.EstimatedWeight,
		PickupAddress: models.Address{
			Street:   pbOrder.PickupAddress.Street,
			District: pbOrder.PickupAddress.District,
			City:     pbOrder.PickupAddress.City,
			Lat:      pbOrder.PickupAddress.Lat,
			Lng:      pbOrder.PickupAddress.Lng,
		},
		ScheduledTime: pbOrder.ScheduledTime.AsTime(),
		Notes:         pbOrder.Notes,
		CreatedAt:     pbOrder.CreatedAt.AsTime(),
		UpdatedAt:     pbOrder.UpdatedAt.AsTime(),
	}

	// Convert collector ID if present
	if pbOrder.CollectorId != "" {
		collectorID, _ := primitive.ObjectIDFromHex(pbOrder.CollectorId)
		order.CollectorID = &collectorID
	}

	// Convert actual weight if present
	if pbOrder.ActualWeight > 0 {
		order.ActualWeight = &pbOrder.ActualWeight
	}

	// Convert completed time if present
	if pbOrder.CompletedTime != nil {
		completedTime := pbOrder.CompletedTime.AsTime()
		order.CompletedTime = &completedTime
	}

	// Convert waste types
	for _, wasteType := range pbOrder.WasteTypes {
		order.WasteTypes = append(order.WasteTypes, models.WasteType(wasteType.String()))
	}

	// Convert payment
	if pbOrder.Payment != nil {
		order.Payment = models.Payment{
			Amount:   pbOrder.Payment.Amount,
			Currency: pbOrder.Payment.Currency,
			Method:   pbOrder.Payment.Method,
			IsPaid:   pbOrder.Payment.IsPaid,
		}
		if pbOrder.Payment.PaymentTime != nil {
			paymentTime := pbOrder.Payment.PaymentTime.AsTime()
			order.Payment.PaymentTime = &paymentTime
		}
	}

	return order
}

// convertWasteTypesToProto converts GraphQL waste types to protobuf
func convertWasteTypesToProto(wasteTypes []models.WasteType) []pb_order.WasteType {
	var pbWasteTypes []pb_order.WasteType
	for _, wt := range wasteTypes {
		wtStr := string(wt)
		switch wtStr {
		case "PAPER":
			pbWasteTypes = append(pbWasteTypes, pb_order.WasteType_WASTE_TYPE_PAPER)
		case "PLASTIC":
			pbWasteTypes = append(pbWasteTypes, pb_order.WasteType_WASTE_TYPE_PLASTIC)
		case "METAL":
			pbWasteTypes = append(pbWasteTypes, pb_order.WasteType_WASTE_TYPE_METAL)
		case "GLASS":
			pbWasteTypes = append(pbWasteTypes, pb_order.WasteType_WASTE_TYPE_GLASS)
		case "ELECTRONIC":
			pbWasteTypes = append(pbWasteTypes, pb_order.WasteType_WASTE_TYPE_ELECTRONIC)
		}
	}
	return pbWasteTypes
}