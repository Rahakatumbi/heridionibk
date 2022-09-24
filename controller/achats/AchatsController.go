package achats

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

//head-quater
type helperAchat struct {
	Id         int    `json:"id"`
	SupplierId int    `json:"supplier_id"`
	BrancheId  int    `json:"branche_id"`
	CreatedBy  int    `json:"creator"`
	Info       []Info `json:"data"`
}
type Info struct {
	Id        int     `json:"id"`
	ProduitId int     `json:"produit_id"`
	Kgs       float32 `json:"kgs"`
	UnitPrice float32 `json:"price"`
	Qualite   int     `json:"quality"`
	ChampID   int     `json:"field_id"`
}
type HelperData struct {
	Id         int                  `json:"id"`
	BrancheId  int                  `json:"branche_id"`
	SupplierId int                  `json:"supplier_id"`
	CreatedBy  int                  `json:"creator"`
	Data       []models.AchatsInfos `json:"data"`
}

func ResponseAchat(Achat models.Achats, Stock []models.AchatsInfos) HelperData {
	return HelperData{Id: int(Achat.Id), SupplierId: Achat.SupplierId,
		CreatedBy: Achat.CreatedBy, Data: Stock}
}
func Achat(c *gin.Context) {
	var helper helperAchat
	c.ShouldBindJSON(&helper)
	// var response HelperData
	achat := models.Achats{SupplierId: helper.SupplierId, CreatedBy: helper.CreatedBy, BrancheId: helper.BrancheId}
	config.DB.Create(&achat)
	for _, data := range helper.Info {
		// price, _ := strconv.ParseFloat(data.UnitPrice, 32)
		// kgs, _ := strconv.ParseFloat(data.Kgs, 32)
		achatID := achat.Id
		achatinfo := models.AchatsInfos{AchatId: int(achatID), ProduitId: data.ProduitId, Kgs: data.Kgs, UnitPrice: data.UnitPrice,
			ChampId: data.ChampID, Qualite: data.Qualite}
		config.DB.Create(&achatinfo).Scan(&data)
	}
	c.JSON(200, helper)
}

//all achats head quater
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
