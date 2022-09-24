package models

import "gorm.io/gorm"

type Orders struct {
	Id              uint   `json:"id" gorm:"primary_id"`
	Type            int    `json:"type"`
	ClientId        int    `json:"client_id"`
	Status          int    `json:"status"`
	LieuDeLivraison string `json:"lieu_de_livraison"`
	ModeDePayement  int    `json:"mode_de_payement"` //1.CIF,2.FOB,
	Creator         int    `json:"creator"`
	gorm.Model
}
type OrdersInfo struct {
	Id             uint    `json:"id" gorm:"primary_id"`
	OrderId        int     `json:"order_id"`
	ProduitId      int     `json:"produit_id"`
	Quality        int     `json:"quality"`
	Quantity       float32 `json:"quantity"`
	Echeance       string  `json:"echeance"`
	ServedQuantity float32 `json:"served_quantity"`
	gorm.Model
}
type FinancementOrder struct {
	Id      uint    `json:"id" gorm:"primary_id"`
	OrderId int     `json:"order_id"`
	Montant float32 `json:"montant"`
	gorm.Model
}

type ServedOrder struct {
	Id      uint `json:"id" gorm:"primary_id"`
	OrderId int  `json:"order_id"`
	Creator int  `json:"creator"`
	gorm.Model
}
type ServedOrderInfo struct {
	Id            int     `json:"id" gorm:"primary_id"`
	ServedOrderId int     `json:"served_id"`
	OrderId       int     `json:"order_id"`
	ProduitId     int     `json:"produit_id"`
	Quality       int     `json:"quality"`
	Quantity      float32 `json:"quantity"`
	Price         float32 `json:"price"`
	AchatInfoId   int     `json:"achat_info"`
	gorm.Model
}
