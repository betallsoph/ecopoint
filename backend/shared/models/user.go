package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType string

const (
	UserTypeCustomer  UserType = "customer"
	UserTypeCollector UserType = "collector"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirebaseUID string             `bson:"firebase_uid" json:"firebase_uid"`
	Email       string             `bson:"email" json:"email"`
	DisplayName string             `bson:"display_name" json:"display_name"`
	PhotoURL    string             `bson:"photo_url,omitempty" json:"photo_url,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	UserType    UserType           `bson:"user_type" json:"user_type"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Customer struct {
	User    `bson:",inline"`
	Address *Address `bson:"address,omitempty" json:"address,omitempty"`
}

type Collector struct {
	User     `bson:",inline"`
	Rating   float64  `bson:"rating" json:"rating"`
	IsOnline bool     `bson:"is_online" json:"is_online"`
	Vehicle  *Vehicle `bson:"vehicle,omitempty" json:"vehicle,omitempty"`
	Location *Location `bson:"location,omitempty" json:"location,omitempty"`
}

type Address struct {
	Street   string  `bson:"street" json:"street"`
	District string  `bson:"district" json:"district"`
	City     string  `bson:"city" json:"city"`
	Lat      float64 `bson:"lat" json:"lat"`
	Lng      float64 `bson:"lng" json:"lng"`
}

type Vehicle struct {
	Type         string `bson:"type" json:"type"`
	LicensePlate string `bson:"license_plate" json:"license_plate"`
}

type Location struct {
	Lat       float64   `bson:"lat" json:"lat"`
	Lng       float64   `bson:"lng" json:"lng"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}