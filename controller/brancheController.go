package controller

import (
	"time"

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
	if len(info) > 0 {
		c.JSON(200, info)
	} else {
		c.JSON(200, achats)
	}
}
func FinancementDepot(c *gin.Context) {
	var fin models.FinancementDepot
	c.ShouldBindJSON(&fin)
	config.DB.Save(&models.FinancementDepot{Id: fin.Id, Montant: fin.Montant, Createdby: fin.Createdby, Status: 1,
		Motif: fin.Motif, Description: fin.Description, DepotId: fin.DepotId, UsedAmount: 0}).Scan(&fin)
	c.JSON(200, fin)
}
func AllFinancement(c *gin.Context) {
	var fin []models.FinancementDepot
	var helps []ResponseFinancement
	config.DB.Select("financement_depots.id as id,financement_depots.motif,financement_depots.description,financement_depots.montant,financement_depots.status,financement_depots.created_at as date,financement_depots.used_amount,financement_depots.depot_id,financement_depots.createdby,branches.id as brancheid,branches.axe_id,branches.names as depot,users.id as userId,users.names as names,users.code,financement_depots.updated_at,axes.id as axeid,axes.names as axe").
		Joins("inner join branches on branches.id = financement_depots.depot_id").
		Joins("inner join users on users.id = financement_depots.createdby").
		Joins("inner join axes on axes.id = branches.axe_id").Find(&fin).Scan(&helps)
	if len(fin) > 0 {
		c.JSON(200, helps)
	} else {
		c.JSON(200, fin)
	}
}

type CheckFinancement struct {
	Id         int     `json:"id"`
	Montant    float32 `json:"montant"`
	UsedAmount float32 `json:"used_amount"`
	DepotId    int     `json:"depot_id"`
}

func FinancementByDepot(c *gin.Context) {
	var fin []models.FinancementDepot
	var helps []ResponseFinancement

	// var check CheckFinancement
	// config.DB.Select("used_amount as used,montant as montant,depot_id,id").Where("montant - used_amount >=? AND depot_id=?", 1, 1).First(&models.FinancementDepot{}).Scan(&check)

	config.DB.Select("financement_depots.id as id,financement_depots.montant,financement_depots.motif,financement_depots.description,financement_depots.status,financement_depots.created_at as date,financement_depots.used_amount,financement_depots.depot_id,financement_depots.createdby,branches.id as brancheid,branches.axe_id,branches.names as depot,users.id as useId,users.names as names,users.code,financement_depots.updated_at,axes.id as axeid,axes.names as axes").
		Joins("inner join branches on branches.id = financement_depots.depot_id").
		Joins("inner join users on users.id = financement_depots.createdby").
		Joins("inner join axes on axes.id = branches.axe_id").
		Where("depot_id", c.Param("id")).Find(&fin).Scan(&helps)
	if len(fin) > 0 {
		c.JSON(200, helps)
	} else {
		c.JSON(200, fin)
	}
}
func Finance(c *gin.Context) {
	var check []CheckFinancement
	var fin []models.FinancementDepot
	config.DB.Select("montant,used_amount,depot_id,id,depot_id,id").Where("montant - used_amount >=? AND depot_id=?", 1, c.Param("id")).Find(&fin).Scan(&check)
	c.JSON(200, check)
}

type ResponseFinancement struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Montant     float32   `json:"montant"`
	DepotId     int       `json:"depot_id"`
	Motif       string    `json:"motif"`
	Description string    `json:"description"`
	UsedAmount  float32   `json:"used_amount"`
	Status      int       `json:"status"`
	Createdby   int       `json:"createdby"`
	EnvoyerPar  string    `json:"names"`
	Depot       string    `json:"depot"`
	Date        time.Time `json:"date"`
	Axe         string    `json:"axe"`
	Code        string    `json:"code"`
}
