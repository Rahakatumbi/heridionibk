package models

import "gorm.io/gorm"

type AchatsInfos struct {
	Id           uint    `json:"id" gorm:"primary_key"`
	ProduitId    int     `json:"produit_id"`
	AchatId      int     `json:"achat_id"`
	Kgs          float32 `json:"kgs"`
	ChampId      int     `json:"champ_id"`
	Qualite      int     `json:"qualite"`
	UnitPrice    float32 `json:"price"`
	UsedQuantity float32 `json:"used_quantity"`
	gorm.Model
}
