package ventes

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type FactureInfo struct {
	OrderInfo int     `json:"id"`
	OrderId   int     `json:"order_id"`
	ItemId    int     `json:"product_id"`
	Kgs       float32 `json:"quantity"`
	Price     float32 `json:"price"`
	Qualite   int     `json:"qualite"`
}
type Facture struct {
	ClientId   int           `json:"client_id"`
	Montant    float32       `json:"montant"`
	PaidAmount float32       `json:"paid_amount"`
	BanqueId   int           `json:"banque_id"`
	Bordereau  string        `json:"bordereau"`
	CreatedBy  int           `json:"creator"`
	OrderId    int           `json:"order_id"`
	Data       []FactureInfo `json:"items"`
}

func Facturation(c *gin.Context) {
	var facture Facture
	c.ShouldBindJSON(&facture)
	if (facture.Montant - facture.PaidAmount) == 0 {
		order := models.Orders{Status: 3, Id: uint(facture.OrderId)}
		config.DB.Updates(&order)
	} else if (facture.Montant - facture.PaidAmount) >= 1 {
		order := models.Orders{Status: 2, Id: uint(facture.OrderId)}
		config.DB.Updates(&order)
	}
	fac := models.Invoice{Montant: facture.Montant, PaidAmount: facture.PaidAmount, BanqueId: facture.BanqueId, ClientId: facture.ClientId,
		Bordereau: facture.Bordereau, CreatedBy: facture.CreatedBy, OrderId: facture.OrderId}
	config.DB.Create(&fac)
	for _, data := range facture.Data {
		info := models.InvoiceInfo{InvoiceId: int(fac.Id), OrderInfoId: data.OrderInfo, Qualite: data.Qualite, OrderId: data.OrderId, ItemId: data.ItemId, Kgs: data.Kgs, UnitPrice: data.Price, CreatedBy: fac.CreatedBy}
		config.DB.Create(&info)

	}
	c.JSON(200, facture)
}
