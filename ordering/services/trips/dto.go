package trips

import "snappfood/ordering/internal/models"

type CreateTripCommand struct {
	OrderID uint `json:"orderID"`
	BikeID  uint `json:"bikeID"`
}
type ChangeTripStatusCommand struct {
	OrderID uint              `json:"orderID"`
	Status  models.TripStatus `json:"status"`
}
