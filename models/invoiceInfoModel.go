package models

type InvoiceInfo struct {
	Id        uint    `json:"id" gorm:"primary_key"`
	InvoiceId int     `json:"inovoice_id"`
	Quantite  float32 `json:"quantite"`
	Kgs       float32 `json:"kgs"`
	ItemId    int     `json:"item_id"`
	UnitPrice float32 `json:"unit_price"`
	CreatedBy int     `json:"creator"`
	// gorm.Model
}
