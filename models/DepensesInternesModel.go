package models

import "gorm.io/gorm"

type Depenses struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	Montant     float32 `json:"montant"`
	Motif       string  `json:"motif"`
	Description string  `json:"description"`
	Createdby   int     `json:"createdby"`
	gorm.Model
}
type Documents struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Intitule    string `json:"intitule"`
	Object      string `json:"object"`
	Document    string `json:"document"`
	Description string `json:"description"`
	Createdby   int    `json:"createdby"`
	gorm.Model
}
