package achats

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

// head-quater
type helperAchat struct {
	Id         int    `json:"id"`
	SupplierId int    `json:"supplier_id"`
	BrancheId  int    `json:"branche_id"`
	CreatedBy  int    `json:"creator"`
	Info       []Info `json:"data"`
}
type Info struct {
	Id        int     `json:"id"`
	Price     int     `json:"prix_id"`
	ProduitId int     `json:"produit_id"`
	Kgs       float32 `json:"kgs"`
	UnitPrice float32 `json:"price"`
	Qualite   int     `json:"quality"`
	ChampID   int     `json:"field_id"`
}
type HelperData struct {
	Id          int                  `json:"id"`
	Financement float32              `json:"financement"`
	SupplierId  int                  `json:"fornisseur_id"`
	CreatedBy   int                  `json:"creator"`
	BrancheId   int                  `json:"branche_id"`
	Data        []models.AchatsInfos `json:"data"`
}

func ResponseAchat(Achat models.Achats, Stock []models.AchatsInfos) HelperData {
	return HelperData{Id: int(Achat.Id), SupplierId: Achat.SupplierId, Financement: Achat.Financement,
		CreatedBy: Achat.CreatedBy, Data: Stock}
}

type CheckFinancement struct {
	Id         int     `json:"id"`
	Montant    float32 `json:"montant"`
	UsedAmount float32 `json:"used_amount"`
	DepotId    int     `json:"depot_id"`
}

func Achat(c *gin.Context) {
	var helper helperAchat
	c.ShouldBindJSON(&helper)
	//check the financement available

	var check CheckFinancement
	config.DB.Select("used_amount as used,montant as montant,depot_id,id").Where("montant - used_amount >=? AND depot_id=?", 1, helper.BrancheId).First(&models.FinancementDepot{}).Scan(&check)
	var res float32 = 0
	for _, data := range helper.Info {
		res += float32(data.UnitPrice) * float32(data.Kgs)
	}
	var b float32 = check.Montant - check.UsedAmount
	if b > 0 {
		if b-res <= 0 {
			var c float32 = res - b
			config.DB.Updates(&models.FinancementDepot{Id: uint(check.Id), UsedAmount: check.UsedAmount + c})
			config.DB.Select("used_amount as used,montant as montant,depot_id,id").Where("montant - used_amount >=? AND depot_id=?", 1, helper.BrancheId).First(&models.FinancementDepot{}).Scan(&check)
			config.DB.Updates(&models.FinancementDepot{Id: uint(check.Id), UsedAmount: check.UsedAmount + b})
		} else {
			config.DB.Select("used_amount as used,montant as montant,depot_id,id").Where("montant - used_amount >=? AND depot_id=?", 1, helper.BrancheId).First(&models.FinancementDepot{}).Scan(&check)
			config.DB.Updates(&models.FinancementDepot{Id: uint(check.Id), UsedAmount: check.UsedAmount + b})
		}
	}
	// var response HelperData
	achat := models.Achats{SupplierId: helper.SupplierId, CreatedBy: helper.CreatedBy, BrancheId: helper.BrancheId, Financement: res}
	config.DB.Create(&achat)
	for _, data := range helper.Info {
		achatID := achat.Id
		achatinfo := models.AchatsInfos{AchatId: int(achatID), ProduitId: data.ProduitId, Kgs: data.Kgs, UnitPrice: data.UnitPrice,
			ChampId: data.ChampID, Qualite: data.Qualite, PriceId: data.Price}
		config.DB.Create(&achatinfo).Scan(&data)
	}
	//save info financement
	config.DB.Create(&models.FinancementInfo{AchatId: int(achat.Id), Montant: check.Montant, DepotId: check.DepotId, Createdby: achat.CreatedBy})
	c.JSON(200, helper)

}

type ResponseAchatData struct {
	Id          uint    `json:"id" gorm:"primary_key"`
	SupplierId  int     `json:"fornisseur_id"`
	Names       string  `json:"supplier"`
	Quantity    float32 `json:"quantity"`
	ProduitId   int     `json:"produit_id"`
	AchatInfoId int     `json:"achat_info_id"`
	Qualite     int     `json:"qualite"`
	Price       float32 `json:"price"`
	Produit     string  `json:"produit"`
}

func AchatBrancheById(c *gin.Context) {
	var achats []models.Achats
	var res []ResponseAchatData
	// config.DB.Find(&achats, "branche_id =?", c.Param("id")).Scan(&res)
	config.DB.Select("achats.id as id, achats.supplier_id,achats_infos.id as achat_info_id,achats.branche_id,achats_infos.kgs as quantity,achats_infos.produit_id,achats_infos.qualite,achats_infos.unit_price as price,suppliers.id as sup,suppliers.names as names, suppliers.code,products.id as pro,products.names as produit").
		Joins("inner join suppliers on suppliers.id =achats.supplier_id").
		Joins("inner join achats_infos on achats_infos.achat_id = achats.id").
		Joins("inner join products on products.id = achats_infos.produit_id").
		Where("achats.branche_id =?", c.Param("id")).
		Find(&achats).Scan(&res)
	c.JSON(200, res)
}

// all achats head quater
func Achats(c *gin.Context) {
	var achats []models.Achats
	var helper []HelperData
	config.DB.Find(&achats)
	for _, achat := range achats {
		var data []models.AchatsInfos
		config.DB.Find(&data, "achat_id", achat.Id)
		responseOrder := ResponseAchat(achat, data)
		helper = append(helper, responseOrder)
	}
	c.JSON(200, helper)
}
