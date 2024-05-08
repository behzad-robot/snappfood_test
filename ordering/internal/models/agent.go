package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Name string `json:"name"`
}
