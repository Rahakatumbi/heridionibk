package clients

import (
	"strconv"

	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Supllier(c *gin.Context) {
	var sub models.Suppliers
	config.DB.Select("id").Last(&sub)
	code := "FMHD" + strconv.Itoa(int(sub.Id))
	supplier := models.Suppliers{Code: code}
	c.ShouldBindJSON(&supplier)
	config.DB.Create(&supplier)
	c.JSON(200, supplier)
}
func Suppliers(c *gin.Context) {
	var suplliers []models.Suppliers
	config.DB.Find(&suplliers)
	c.JSON(200, suplliers)
}
func Search(c *gin.Context) {
	var help HelpSearch
	var supllier models.Suppliers
	config.DB.Select("names,code,address,telephone,id").Where("names LIKE ?", c.Param("id")).Or("code LIKE ?", c.Param("id")).First(&supllier).Scan(&help.Supplier)
	var axe models.Axes
	var champ []models.Champs
	config.DB.Find(&champ, "supplier_id=?", supllier.Id).Scan(&help.Champs)
	for _, data := range champ {
		config.DB.Find(&axe, data.AxeId).Scan(&help.Axe)
	}
	c.JSON(200, help)
}
func AddChamp(c *gin.Context) {
	var champs models.Champs
	c.ShouldBind(&champs)
	config.DB.Save(&champs)
	c.JSON(200, champs)
}

type HelpChamp struct {
	Id         uint             `json:"id" gorm:"primary_key"`
	SupplierId int              `json:"supplier_id"`
	AxeId      int              `json:"axe_id"`
	Plante     float32          `json:"plante"`
	Size       float32          `json:"size"`
	Adelivrer  float32          `json:"adelivrer"`
	Address    string           `json:"address"`
	Createdby  int              `json:"createdby"`
	Supplier   models.Suppliers `json:"supplier"`
	Axe        models.Axes      `json:"axe"`
	User       models.Users     `json:"user"`
}
type HelpSearch struct {
	Supplier models.Suppliers `json:"supplier"`
	Axe      models.Axes      `json:"axe"`
	Champs   []models.Champs  `json:"champs"`
}

type ChampHelp struct {
	Champs []models.Champs `json:"champs"`
	Axes   []models.Axes   `json:"axe"`
}

func ResponseChamp(Champ models.Champs, Axe models.Axes, Supplier models.Suppliers, User models.Users) HelpChamp {
	return HelpChamp{Id: Champ.Id, SupplierId: Champ.SupplierId, Createdby: Champ.Createdby, Address: Champ.Address, AxeId: Champ.AxeId,
		Axe: Axe, Supplier: Supplier, User: User, Adelivrer: Champ.Adelivrer, Size: Champ.Size, Plante: Champ.Plante}
}
func Champs(c *gin.Context) {
	var champs []models.Champs
	var helper []HelpChamp
	config.DB.Where("supplier_id=?", c.Param("id")).Find(&champs)
	for _, data := range champs {
		var supllier models.Suppliers
		var axe models.Axes
		var User models.Users
		config.DB.Find(&supllier, data.SupplierId)
		config.DB.Find(&axe, data.AxeId)
		config.DB.Find(&User, data.Createdby)
		responsedata := ResponseChamp(data, axe, supllier, User)
		helper = append(helper, responsedata)
	}
	if len(helper) > 0 {
		c.JSON(200, helper)
	} else {
		c.JSON(200, champs)
	}
}
func AllChamps(c *gin.Context) {
	var champs []models.Champs
	var helper []HelpChamp
	config.DB.Find(&champs)
	for _, data := range champs {
		var supllier models.Suppliers
		var axe models.Axes
		var User models.Users
		config.DB.Find(&supllier, data.SupplierId)
		config.DB.Find(&axe, data.AxeId)
		config.DB.Find(&User, data.Createdby)
		responsedata := ResponseChamp(data, axe, supllier, User)
		helper = append(helper, responsedata)
	}
	if len(helper) > 0 {
		c.JSON(200, helper)
	} else {
		c.JSON(200, champs)
	}
}
