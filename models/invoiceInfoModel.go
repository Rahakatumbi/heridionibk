package models

type InvoiceInfo struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	InvoiceId   int     `json:"inovoice_id"`
	OrderInfoId int     `json:"order_info_id"`
	OrderId     int     `json:"order_id"`
	Qualite     int     `json:"qualite"`
	Kgs         float32 `json:"kgs"`
	ItemId      int     `json:"item_id"`
	UnitPrice   float32 `json:"unit_price"`
	CreatedBy   int     `json:"creator"`
	// gorm.Model
}
