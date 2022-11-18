package models

import (
	"gorm.io/gorm"
)

type Prices struct {
	Id        uint    `json:"id" gorm:"primary_id"`
	Price     float32 `json:"price"`
	Status    string  `json:"status"`
	CreatedBy int     `json:"createdby"`
	gorm.Model
}
