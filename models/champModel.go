package models

import "gorm.io/gorm"

type Champs struct {
	Id         uint    `json:"id" gorm:"primary_key"`
	SupplierId int     `json:"supplier_id"`
	AxeId      int     `json:"axe_id"`
	Plante     float32 `json:"plante"`
	Size       float32 `json:"size"`
	Adelivrer  float32 `json:"adelivrer"`
	Address    string  `json:"address"`
	Createdby  int     `json:"createdby"`
	gorm.Model
}
