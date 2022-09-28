package models

import "gorm.io/gorm"

type Invoice struct {
	Id         uint    `json:"id" gorm:"primary_key"`
	ClientId   int     `json:"client_id"`
	Montant    float32 `json:"montant"`
	PaidAmount float32 `json:"paid_amount"`
	BanqueId   int     `json:"banque_id"`
	Bordereau  string  `json:"bordereau"`
	CreatedBy  int     `json:"creator"`
	OrderId    int     `json:"order_id"`
	gorm.Model
}
