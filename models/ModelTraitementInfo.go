package models

import "gorm.io/gorm"

type TraitementInfo struct {
	Id                uint    `json:"id" gorm:"primary_key"`
	ApprovisionInfoId int     `json:"achat_info_id"`
	BrancheId         int     `json:"branche_id"`
	OrderId           int     `json:"order_id"`
	OrderInfoId       int     `json:"order_info_id"`
	ProductId         int     `json:"produit_id"`
	Quality           int     `json:"quality"`
	Quantity          float32 `json:"quantity"`
	CreatedBy         int     `json:"createdBy"`
	gorm.Model
}
