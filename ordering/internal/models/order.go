package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID              uint  `json:"userID"`
	VendorID            uint  `json:"vendorID"`
	InitialDeliveryTime int64 `json:"initialDeliveryTime"`
	DeliveryTime        int64 `json:"deliveryTime"`
}

func (order *Order) GetDeliveredAt() time.Time {
	return order.CreatedAt.Add(time.Duration(order.DeliveryTime))
}
func (order *Order) GetInitialDeliveredAt() time.Time {
	return order.CreatedAt.Add(time.Duration(order.InitialDeliveryTime))
}
func (order *Order) HasDelay() bool {
	return order.GetInitialDeliveredAt().Before(time.Now())
}
func (order *Order) GetDelay() time.Duration {
	return time.Since(order.GetInitialDeliveredAt())
}
