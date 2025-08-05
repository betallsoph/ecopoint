package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusAccepted   OrderStatus = "accepted"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type WasteType string

const (
	WasteTypePaper      WasteType = "paper"
	WasteTypePlastic    WasteType = "plastic"
	WasteTypeMetal      WasteType = "metal"
	WasteTypeGlass      WasteType = "glass"
	WasteTypeElectronic WasteType = "electronic"
)

type Order struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID       primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	CollectorID      *primitive.ObjectID `bson:"collector_id,omitempty" json:"collector_id,omitempty"`
	Status           OrderStatus        `bson:"status" json:"status"`
	WasteTypes       []WasteType        `bson:"waste_types" json:"waste_types"`
	EstimatedWeight  float64            `bson:"estimated_weight" json:"estimated_weight"` // in kg
	ActualWeight     *float64           `bson:"actual_weight,omitempty" json:"actual_weight,omitempty"`
	PickupAddress    Address            `bson:"pickup_address" json:"pickup_address"`
	ScheduledTime    time.Time          `bson:"scheduled_time" json:"scheduled_time"`
	CompletedTime    *time.Time         `bson:"completed_time,omitempty" json:"completed_time,omitempty"`
	Notes            string             `bson:"notes,omitempty" json:"notes,omitempty"`
	Payment          Payment            `bson:"payment" json:"payment"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

type Payment struct {
	Amount      float64 `bson:"amount" json:"amount"`
	Currency    string  `bson:"currency" json:"currency"`
	Method      string  `bson:"method" json:"method"`
	IsPaid      bool    `bson:"is_paid" json:"is_paid"`
	PaymentTime *time.Time `bson:"payment_time,omitempty" json:"payment_time,omitempty"`
}

type OrderCreateRequest struct {
	WasteTypes      []WasteType `json:"waste_types" binding:"required"`
	EstimatedWeight float64     `json:"estimated_weight" binding:"required,min=0.1"`
	PickupAddress   Address     `json:"pickup_address" binding:"required"`
	ScheduledTime   time.Time   `json:"scheduled_time" binding:"required"`
	Notes           string      `json:"notes,omitempty"`
}

type OrderUpdateRequest struct {
	Status       *OrderStatus `json:"status,omitempty"`
	ActualWeight *float64     `json:"actual_weight,omitempty"`
	Notes        *string      `json:"notes,omitempty"`
}