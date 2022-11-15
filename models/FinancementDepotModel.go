package models

import "gorm.io/gorm"

type FinancementDepot struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	Motif       string  `json:"motif"`
	Montant     float32 `json:"montant"`
	DepotId     int     `json:"depot_id"`
	UsedAmount  float32 `json:"used_amount"`
	Status      int     `json:"status"`
	Description string  `json:"description"`
	Createdby   int     `json:"createdby"`
	gorm.Model
}
type FinancementInfo struct {
	Id            uint    `json:"id" gorm:"primary_key"`
	AchatId       int     `json:"achat_id"`
	FinancementId int     `json:"financement_id"`
	Montant       float32 `json:"montant"`
	DepotId       int     `json:"depot_id"`
	Createdby     int     `json:"createdby"`
	gorm.Model
}
