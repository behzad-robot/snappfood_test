package models

import "gorm.io/gorm"

type DelayResult struct {
	gorm.Model
	OrderID   uint  `json:"orderID"`
	VendorID  uint  `json:"vendorID"`
	DelayTime int64 `json:"delayTime"`
}
type DelayResultByVendor struct {
	VendorID   uint  `json:"vendorID"`
	TotalDelay int64 `json:"totalDelay"`
}
