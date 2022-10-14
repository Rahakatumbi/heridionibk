package rapports

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type ApprovisionnemetResponse struct {
	Id           int     `json:"achat_id"`
	SupplierId   int     `json:"fornisseur_id"`
	CreatedBy    int     `json:"creator"`
	BrancheId    int     `json:"branche_id"`
	ProduitId    int     `json:"produit_id"`
	Kgs          float32 `json:"kgs"`
	ChampId      int     `json:"champ_id"`
	Qualite      int     `json:"qualite"`
	UnitPrice    float32 `json:"price"`
	UsedQuantity float32 `json:"used_quantity"`
	Supplier     string  `json:"supplier"`
	User         string  `json:"user"`
	Branche      string  `json:"branche"`
	AxeId        int     `json:"axe_id"`
	Product      string  `json:"product"`
	CreatedAt    string  `json:"createdAt"`
}
type RequestRapp struct {
	Type    int    `json:"type"`
	From    string `json:"from"`
	Until   string `json:"until"`
	DepotId int    `json:"depot_id"`
}

func CreateResponse(achat models.Achats, Achatinfo models.AchatsInfos, branche models.Branche, fermier models.Suppliers, user models.Users,
	produit models.Products) ApprovisionnemetResponse {
	return ApprovisionnemetResponse{Id: int(achat.Id), SupplierId: achat.SupplierId, CreatedBy: achat.CreatedBy, BrancheId: achat.BrancheId,
		Supplier: fermier.Names, User: user.Names, Branche: branche.Names, AxeId: branche.AxeId, Product: produit.Names,
		Kgs: Achatinfo.Kgs, ChampId: Achatinfo.ChampId, Qualite: Achatinfo.Qualite, UnitPrice: Achatinfo.UnitPrice, UsedQuantity: Achatinfo.UsedQuantity,
		ProduitId: Achatinfo.ProduitId, CreatedAt: achat.CreatedAt.Format("02-January-2006:15:04")}
}
func Approvisionnement(c *gin.Context) {
	var response []ApprovisionnemetResponse
	var getpost RequestRapp
	c.ShouldBindJSON(&getpost)
	var achats []models.Achats
	if getpost.Type == 0 {
		//report general
		config.DB.Select("achats.id,achats.supplier_id,achats.created_by,achats.created_at").Where("achats.created_at >=? AND achats.created_at <=?", getpost.From, getpost.Until).Find(&achats)
		for _, data := range achats {
			var achatinfo models.AchatsInfos
			var branche models.Branche
			var user models.Users
			var fermier models.Suppliers
			var produit models.Products
			config.DB.Find(&achatinfo, "achat_id=?", data.Id)
			config.DB.Find(&branche, "id=?", data.BrancheId)
			config.DB.Find(&user, "id", data.CreatedBy)
			config.DB.Find(&fermier, "id", data.SupplierId)
			config.DB.Find(&produit, "id", achatinfo.ProduitId)
			setresponse := CreateResponse(data, achatinfo, branche, fermier, user, produit)
			response = append(response, setresponse)
		}
	} else {
		//selon branche
		config.DB.Find(&achats).Where("created_at >=? AND created_at <=? AND branche_id=?", getpost.From, getpost.Until, getpost.DepotId)
		for _, data := range achats {
			var achatinfo models.AchatsInfos
			var branche models.Branche
			var user models.Users
			var fermier models.Suppliers
			var produit models.Products
			config.DB.Find(&achatinfo, "achat_id=?", data.Id)
			config.DB.Find(&branche, "id=?", data.BrancheId)
			config.DB.Find(&user, "id", data.CreatedBy)
			config.DB.Find(&fermier, "id", data.SupplierId)
			config.DB.Find(&produit, "id", achatinfo.ProduitId)
			setresponse := CreateResponse(data, achatinfo, branche, fermier, user, produit)
			response = append(response, setresponse)
		}
	}
	if len(response) > 0 {
		c.JSON(200, response)
	} else {
		c.JSON(200, achats)
	}
}
