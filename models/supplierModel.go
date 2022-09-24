package models

import "gorm.io/gorm"

type Suppliers struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Names     string `json:"names"`
	Telephone int    `json:"phone"`
	Code      string `json:"code" gorm:"unique"`
	BrancheId int    `json:"branche_id"` // 2.branche
	Address   string `json:"address"`
	Sexe      int    `json:"sexe"`
	EtatCivil int    `json:"civil"`
	Bod       string `json:"bod"`
	Createdby int    `json:"creator"`
	gorm.Model
}
