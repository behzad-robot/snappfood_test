package orders

type CreateOrderCommand struct {
	UserID       uint  `json:"userID"`
	VendorID     uint  `json:"vendorID"`
	DeliveryTime int64 `json:"deliveryTime"`
}
