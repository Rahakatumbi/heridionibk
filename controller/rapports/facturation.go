package rapports

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type FacturationResponse struct {
	Id           int     `json:"achat_id"`
	ClientId     int     `json:"fornisseur_id"`
	CreatedBy    int     `json:"creator"`
	BrancheId    int     `json:"branche_id"`
	ItemId       int     `json:"produit_id"`
	Kgs          float32 `json:"kgs"`
	ChampId      int     `json:"champ_id"`
	Qualite      int     `json:"qualite"`
	UnitPrice    float32 `json:"price"`
	UsedQuantity float32 `json:"used_quantity"`
	Supplier     string  `json:"client"`
	User         string  `json:"user"`
	Branche      string  `json:"branche"`
	AxeId        int     `json:"axe_id"`
	Product      string  `json:"product"`
	CreatedAt    string  `json:"createdAt"`
}
type RequestExpedition struct {
	Type     int    `json:"type"`
	From     string `json:"from"`
	Until    string `json:"until"`
	ClientId int    `json:"client_id"`
}

func ExpeditonResponse(achat models.Invoice, Achatinfo models.InvoiceInfo, branche models.Branche, fermier models.Clients, user models.Users,
	produit models.Products) FacturationResponse {
	return FacturationResponse{Id: int(achat.Id), ClientId: achat.ClientId, CreatedBy: achat.CreatedBy,
		Supplier: fermier.Names, User: user.Names, Branche: branche.Names, AxeId: branche.AxeId, Product: produit.Names,
		Kgs: Achatinfo.Kgs, Qualite: Achatinfo.Qualite, UnitPrice: Achatinfo.UnitPrice, UsedQuantity: Achatinfo.Kgs,
		ItemId: Achatinfo.ItemId, CreatedAt: achat.CreatedAt.Format("02-January-2006:15:04")}
}
func Facturation(c *gin.Context) {
	var response []FacturationResponse
	var getpost RequestExpedition
	c.ShouldBindJSON(&getpost)
	var achats []models.Invoice
	if getpost.Type == 0 {
		config.DB.Where("created_at >=? AND created_at <=?", getpost.From, getpost.Until).Find(&achats)
		for _, data := range achats {
			var achatinfo models.InvoiceInfo
			var branche models.Branche
			var user models.Users
			var fermier models.Clients
			var produit models.Products
			config.DB.Find(&achatinfo, "invoice_id=?", data.Id)
			// config.DB.Find(&branche, "id=?", data.BrancheId)
			config.DB.Find(&user, "id", data.CreatedBy)
			config.DB.Find(&fermier, "id", data.ClientId)
			config.DB.Find(&produit, "id", achatinfo.ItemId)
			setresponse := ExpeditonResponse(data, achatinfo, branche, fermier, user, produit)
			response = append(response, setresponse)
		}
	} else {
		config.DB.Find(&achats).Where("created_at >=? AND created_at <=? AND client_id=?", getpost.From, getpost.Until, getpost.ClientId)
		for _, data := range achats {
			var achatinfo models.InvoiceInfo
			var branche models.Branche
			var user models.Users
			var fermier models.Clients
			var produit models.Products
			config.DB.Find(&achatinfo, "invoice_id=?", data.Id)
			// config.DB.Find(&branche, "id=?", data.BrancheId)
			config.DB.Find(&user, "id", data.CreatedBy)
			config.DB.Find(&fermier, "id", data.ClientId)
			config.DB.Find(&produit, "id", achatinfo.ItemId)
			setresponse := ExpeditonResponse(data, achatinfo, branche, fermier, user, produit)
			response = append(response, setresponse)
		}
	}
	c.JSON(200, response)
}
