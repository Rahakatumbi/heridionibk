package controller

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/controller/find"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Branche(c *gin.Context) {
	var branche models.Branche
	c.ShouldBindJSON(&branche)
	config.DB.Save(&branche)
	find.FindUser(branche.ChefID, &branche.Chef)
	c.JSON(200, branche)
}
func Branches(c *gin.Context) {
	var barnches []models.Branche
	config.DB.Joins("branches").Find(&barnches)
	// for _,branche := range barnches {
	// 	var user models.Users
	// 	config.DB.Find(user,"chef_id=?",branche.ChefID)
	// }
	c.JSON(200, barnches)
}

type HelperData struct {
	Id         int                  `json:"id"`
	BrancheId  int                  `json:"branche_id"`
	SupplierId int                  `json:"supplier_id"`
	CreatedBy  int                  `json:"creator"`
	Data       []models.AchatsInfos `json:"data"`
}
type BrancheItem struct {
	Id         int                  `json:"id"`
	BrancheId  int                  `json:"branche_id"`
	SupplierId int                  `json:"supplier_id"`
	CreatedBy  int                  `json:"creator"`
	Branche    models.Branche       `json:"branche"`
	Data       []models.AchatsInfos `json:"data"`
}
type TraiteItem struct {
	Data []models.AchatsInfos `json:"data"`
}
type AchatM struct {
	Id           uint    `json:"id" gorm:"primary_key"`
	ProduitId    int     `json:"produit_id"`
	AchatId      int     `json:"achat_id"`
	Kgs          float32 `json:"kgs"`
	ChampId      int     `json:"champ_id"`
	Qualite      int     `json:"qualite"`
	UnitPrice    float32 `json:"price"`
	UsedQuantity float32 `json:"used_quantity"`
	Names        string  `json:"names"`
}

func ResponseAchat(Achat models.Achats, Stock []models.AchatsInfos) HelperData {
	return HelperData{Id: int(Achat.Id), SupplierId: Achat.SupplierId,
		CreatedBy: Achat.CreatedBy, Data: Stock}
}
func TraitementResponse(Data []models.AchatsInfos) TraiteItem {
	return TraiteItem{Data: Data}
}
func ResponseItem(Branche models.Branche, Achat models.Achats, Stock []models.AchatsInfos) BrancheItem {
	return BrancheItem{Branche: Branche, Id: int(Achat.Id), SupplierId: Achat.SupplierId,
		CreatedBy: Achat.CreatedBy, Data: Stock}
}
func TraiteItems(c *gin.Context) {
	var achats []models.Achats
	var helper []HelperData
	config.DB.Find(&achats, "branche_id=?", c.Param("id"))
	for _, achat := range achats {
		var data []models.AchatsInfos
		config.DB.Find(&data, "achat_id", achat.Id)
		responseOrder := ResponseAchat(achat, data)
		helper = append(helper, responseOrder)
	}
	c.JSON(200, helper)
}
func ResponseAchats(ac models.AchatsInfos, Name models.Products) AchatM {
	return AchatM{Id: ac.Id, AchatId: int(ac.AchatId), ChampId: ac.ChampId, ProduitId: ac.ProduitId,
		Qualite: ac.Qualite, UsedQuantity: ac.UsedQuantity, Kgs: ac.Kgs, UnitPrice: ac.UnitPrice, Names: Name.Names}
}
func DepotData(c *gin.Context) {
	var achats []models.Achats
	var info []AchatM
	config.DB.Select("achats.branche_id,achats.id as achatId,achats.supplier_id,achats_infos.id as id,achats_infos.achat_id,achats_infos.produit_id,achats_infos.kgs,achats_infos.unit_price as price,achats_infos.qualite,achats_infos.used_quantity,products.id as productId,products.names").
		Joins("inner join branches on branches.id=achats.branche_id").Joins("inner join achats_infos on achats_infos.achat_id=achats.id").
		Joins("inner join products on products.id = achats_infos.produit_id").
		Find(&achats, "achats.branche_id=?", c.Param("id")).Scan(&info)
	// for _, achat := range achats {
	// 	var data []models.AchatsInfos
	// 	config.DB.Find(&data, "achat_id", achat.Id)
	// 	for _, item := range data {
	// 		var product models.Products
	// 		config.DB.Find(&product, "id=?", item.ProduitId)
	// 		responseOrder := ResponseAchats(item, product)
	// 		info = append(info, responseOrder)
	// 	}
	// }
	if len(info) > 0 {
		c.JSON(200, info)
	} else {
		c.JSON(200, achats)
	}
}
