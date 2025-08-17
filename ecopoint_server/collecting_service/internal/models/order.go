package models

import (
    "errors"
    "time"
)

type OrderStatus string

const (
    StatusCreated   OrderStatus = "created"
    StatusAccepted  OrderStatus = "accepted"
    StatusOnWay     OrderStatus = "on_way"
    StatusComplete  OrderStatus = "complete"
    StatusCancelled OrderStatus = "cancelled"
)

type Address struct {
    FullText string  `bson:"full_text"`
    Lat      float64 `bson:"lat"`
    Lng      float64 `bson:"lng"`
}

type CustomerSnapshot struct {
    DisplayName string `bson:"display_name"`
    Phone       string `bson:"phone"`
}

type WasteItem struct {
    Type   string  `bson:"type"`
    Weight float64 `bson:"weight"`
}

type Order struct {
    ID                  string            `bson:"id"`
    CustomerID          string            `bson:"customer_id"`
    Status              OrderStatus       `bson:"status"`
    AcceptedBy          *string           `bson:"accepted_by,omitempty"`
    PickAddressSnapshot Address           `bson:"pick_address_snapshot"`
    CustomerSnapshot    CustomerSnapshot  `bson:"customer_snapshot"`
    Items               []WasteItem       `bson:"items"`
    TotalWeight         float64           `bson:"total_weight"`
    EstimatedPrice      float64           `bson:"estimated_price"`
    DistanceKm          float64           `bson:"distance_km"`
    EtaMinutes          int               `bson:"eta_minutes"`
    Note                string            `bson:"note"`
    CreatedAt           time.Time         `bson:"created_at"`
    UpdatedAt           time.Time         `bson:"updated_at"`
    AcceptedAt          *time.Time        `bson:"accepted_at,omitempty"`
    CompletedAt         *time.Time        `bson:"completed_at,omitempty"`
    CancelReason        string            `bson:"cancel_reason,omitempty"`
    CancelSide          CancelBy          `bson:"cancel_side,omitempty"`
    Version             int64             `bson:"version"`
}

var (
    ErrInvalidStatusTransition = errors.New("invalid status transition")
)

// CanTransition validates allowed transitions
func (o *Order) CanTransition(next OrderStatus) bool {
    switch o.Status {
    case StatusCreated:
        return next == StatusAccepted || next == StatusCancelled
    case StatusAccepted:
        return next == StatusOnWay || next == StatusCancelled
    case StatusOnWay:
        return next == StatusComplete || next == StatusCancelled
    case StatusComplete, StatusCancelled:
        return false
    default:
        return false
    }
}

type CancelBy string

const (
    CancelByCustomer CancelBy = "customer"
    CancelByCollector CancelBy = "collector"
    CancelBySystem    CancelBy = "system"
)


