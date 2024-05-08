package analytics

type CreateDelayResultCommand struct {
	OrderID   uint
	VendorID  uint
	DelayTime int64
}
