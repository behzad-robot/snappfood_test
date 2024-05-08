package models

type TripStatus byte

const (
	TRIP_STATUS_NONE      TripStatus = 0
	TRIP_STATUS_ASSIGNED  TripStatus = 1
	TRIP_STATUS_AT_VENDOR TripStatus = 2
	TRIP_STATUS_PICKED    TripStatus = 3
	TRIP_STATUS_DELIVERED TripStatus = 4
)
