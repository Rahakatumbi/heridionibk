package models

import "gorm.io/gorm"

type Achats struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	Financement float32 `json:"financement"`
	SupplierId  int     `json:"fornisseur_id"`
	CreatedBy   int     `json:"creator"`
	BrancheId   int     `json:"branche_id"`
	gorm.Model
}
