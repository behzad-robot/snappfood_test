package models

type VendorDelay struct {
	VendorID uint  `json:"vendorID"`
	Delay    int64 `json:"delay"`
}
