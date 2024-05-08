package models

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	OrderID     uint       `json:"orderID"`
	BikeID      uint       `json:"bikeID"`
	Status      TripStatus `json:"status"`
	DeliveredAt *time.Time `json:"deliveredAt"`
}
