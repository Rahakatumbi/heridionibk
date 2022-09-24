package models

import "gorm.io/gorm"

type Products struct {
	Id        uint   `json:"id" gorm:"primary_id"`
	Names     string `json:"names"`
	CreatedBy int    `json:"creator"`
	gorm.Model
}
