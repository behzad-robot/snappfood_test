package models

import "gorm.io/gorm"

type DelayReport struct {
	gorm.Model
	OrderID            uint       `json:"orderID"`
	TripStatusAtReport TripStatus `json:"tripStatusAtReport"`
	AgentID            int        `json:"agentID"` //-1 => No Agent required //0 waiting for an agent
	IsRepliedByAgent   bool       `json:"isRepliedByAgent"`
	AgentMessage       string     `json:"agentMessage"`
}
