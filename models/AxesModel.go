package models

import "gorm.io/gorm"

type Axes struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Names     string `json:"names"`
	Address   string `json:"address"`
	Createdby int    `json:"createdby"`
	gorm.Model
}
