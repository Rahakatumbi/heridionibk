package models

import "gorm.io/gorm"

type Clients struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Names     string `json:"names"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone int    `json:"phone"`
	Country   int    `json:"country"`
	CreatedBy int    `json:"creator"`
	gorm.Model
}
