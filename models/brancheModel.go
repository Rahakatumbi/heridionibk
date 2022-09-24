package models

import "gorm.io/gorm"

type Branche struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Names     string `json:"names"`
	AxeId     int    `json:"axe_id"`
	Address   string `json:"address"`
	Status    int    `json:"status"`
	Phone     int    `json:"phone"`
	Createdby int    `json:"creator"`
	ChefID    int    `json:"chef_id"`
	Chef      Users  `json:"chef"`
	Axe       Axes   `json:"axe"`
	gorm.Model
}
