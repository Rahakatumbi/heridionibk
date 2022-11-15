package ventes

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Id              uint   `json:"id" gorm:"primary_id"`
	ClientId        int    `json:"client_id"`
	Type            int    `json:"type"`
	Status          int    `json:"status"`
	LieuDeLivraison string `json:"lieu_de_livraison"`
	ModeDePayement  int    `json:"mode_de_payement"` //1.CIF,2.FOB,
	CreatedAt       string `json:"createdAt" format:"02-january-2006:15:04"`
	Names           string `json:"names"`
	//Infos           []models.OrdersInfo `json:"info"`
}
type Data struct {
	Id              uint    `json:"id" gorm:"primary_id"`
	ClientId        int     `json:"client_id"`
	Type            int     `json:"type"`
	Status          int     `json:"status"`
	Montant         float32 `json:"financement"`
	LieuDeLivraison string  `json:"lieu_de_livraison"`
	ModeDePayement  int     `json:"mode_de_payement"` //1.CIF,2.FOB,
	Info            []Info  `json:"info"`
}
type Infos struct {
	Id        uint    `json:"id" gorm:"primary_id"`
	OrderId   int     `json:"order_id"`
	ProduitId int     `json:"product_id"`
	Names     string  `json:"produit"`
	Quality   int     `json:"quality"`
	Quantity  float32 `json:"quantity"`
	Echeance  string  `json:"echeance"`
}
type Info struct {
	Id        uint    `json:"id" gorm:"primary_id"`
	OrderId   int     `json:"order_id"`
	ProduitId int     `json:"product_id"`
	Names     string  `json:"produit"`
	Quality   int     `json:"quality"`
	Quantity  float32 `json:"quantity"`
	Echeance  string  `json:"echeance"`
}

func Order(c *gin.Context) {
	var post Data
	c.ShouldBindJSON(&post)
	insert := models.Orders{ClientId: post.ClientId, Status: 1, LieuDeLivraison: post.LieuDeLivraison, ModeDePayement: post.ModeDePayement, Type: post.Type}
	config.DB.Create(&insert)
	for _, data := range post.Info {
		info := models.OrdersInfo{OrderId: int(insert.Id), ProduitId: data.ProduitId, Quality: data.Quality, Quantity: data.Quantity, Echeance: data.Echeance}
		config.DB.Create(&info)
	}
	//check if there is financement
	if insert.Type == 1 {
		financement := models.FinancementOrder{OrderId: int(insert.Id), Montant: post.Montant}
		config.DB.Create(&financement)
	}
	c.JSON(200, post)
}
func Orders(c *gin.Context) {
	var orders []models.Orders
	var helper []ResponseData
	//config.DB.Find(&orders
	config.DB.Model(&orders).Select("orders.id as id,orders.type,orders.created_at,orders.mode_de_payement,orders.status,orders.client_id,clients.id as clientId,clients.names").
		Joins("inner join clients on clients.id=orders.client_id").
		Where("orders.status != 3 AND orders.status !=2").Scan(&helper)
	if len(helper) > 0 {
		c.JSON(200, helper)
	} else {
		c.JSON(200, orders)
	}
}
func TraitOrders(c *gin.Context) {
	var orders []models.Orders
	var helper []ResponseData
	//config.DB.Find(&orders
	config.DB.Model(&orders).Select("orders.id as id,orders.type,orders.created_at,orders.mode_de_payement,orders.status,orders.client_id,clients.id as clientId,clients.names").
		Joins("inner join clients on clients.id=orders.client_id").
		Where("orders.status != 1").Scan(&helper)
	if len(helper) > 0 {
		c.JSON(200, helper)
	} else {
		c.JSON(200, orders)
	}
}

type Inforesponse struct {
	Id             int     `json:"id"`
	ProduitId      int     `json:"product_id"`
	OrderID        int     `json:"order_id"`
	Names          string  `json:"names"`
	Quality        float32 `json:"quality"`
	Quantity       float32 `json:"quantity"`
	ServedQuantity float32 `json:"served_quantity"`
}

func ResponseInfo(info models.OrdersInfo, pro models.Products) Inforesponse {
	return Inforesponse{Id: int(info.Id), ProduitId: info.ProduitId, Names: pro.Names, Quality: float32(info.Quality), Quantity: info.Quantity,
		ServedQuantity: info.ServedQuantity, OrderID: info.OrderId}
}
func OrderInfos(c *gin.Context) {
	var info []models.OrdersInfo
	var help []Inforesponse
	config.DB.Select("orders_infos.id as id,order_id,orders_infos.produit_id,orders_infos.quality,orders_infos.quantity,orders_infos.echeance,orders_infos.served_quantity,products.id as productId,products.names").Joins("inner join products on products.id=orders_infos.produit_id").
		Find(&info, "orders_infos.order_id=?", c.Param("id")).Scan(&help)
	// for _, data := range info {
	// 	var pr models.Products
	// 	config.DB.Select("names,id").Where("id", data.ProduitId).Find(&pr)
	// 	res := ResponseInfo(data, pr)
	// 	help = append(help, res)
	// }
	if len(help) > 0 {
		c.JSON(200, help)
	} else {
		c.JSON(200, info)
	}
}

type Serve struct {
	Id      int         `json:"id"`
	OrderId int         `json:"order_id"`
	Creator int         `json:"creator"`
	Data    []ServeInfo `json:"info"`
}
type ServeInfo struct {
	Id            int     `json:"id" gorm:"primary_id"`
	ServedOrderId int     `json:"served_id"`
	OrderId       int     `json:"order_id"`
	ProduitId     int     `json:"produit_id"`
	Quality       int     `json:"quality"`
	Quantity      float32 `json:"quantity"`
	Price         float32 `json:"price"`
	AchatInfoId   int     `json:"achat_info_id"`
}

func ServeOrder(c *gin.Context) {
	var helper Serve
	c.ShouldBindJSON(&helper)
	order := models.ServedOrder{OrderId: helper.OrderId, Creator: helper.Creator}
	config.DB.Create(&order).Scan(&helper)
	for _, data := range helper.Data {
		item := models.ServedOrderInfo{ServedOrderId: data.Id, OrderId: data.OrderId, AchatInfoId: data.AchatInfoId,
			Quality: data.Quality, Quantity: data.Quantity, Price: data.Price, ProduitId: data.ProduitId}
		config.DB.Create(&item)
	}
	c.JSON(200, helper)
}

type Orderitem struct {
	Id        int          `json:"id"`
	CreatedBy int          `json:"createdBy"`
	Items     []OrderItems `json:"items"`
}
type OrderItems struct {
	Achat_info_id int     `json:"achat_info_id"`
	Branche_id    int     `json:"branche_id"`
	Order_id      int     `json:"order_id"`
	Order_info_id int     `json:"order_info_id"`
	Produit_id    int     `json:"produit_id"`
	Quality       int     `json:"quality"`
	Quantity      float32 `json:"quantity"`
}

func SolveOrder(c *gin.Context) {
	var items Orderitem
	c.ShouldBindJSON(&items)
	for _, data := range items.Items {
		var md models.OrdersInfo
		var ach models.AchatsInfos
		config.DB.Select("sum(served_quantity) as served_quantity, id").Where("id", data.Order_info_id).Find(&md)
		config.DB.Select("sum(used_quantity) as used_quantity, id").Where("id", data.Achat_info_id).Find(&ach)
		usedqty := ach.UsedQuantity + data.Quantity
		result := md.ServedQuantity + data.Quantity
		nfo := models.OrdersInfo{Id: uint(data.Order_info_id), ServedQuantity: result}
		achat := models.AchatsInfos{Id: uint(data.Order_info_id), UsedQuantity: usedqty}
		config.DB.Updates(&nfo)
		config.DB.Updates(&achat)
		traitement := models.TraitementInfo{OrderId: data.Order_id, ApprovisionInfoId: data.Achat_info_id, Quantity: data.Quantity,
			Quality: data.Quality, ProductId: data.Produit_id, OrderInfoId: data.Order_info_id, BrancheId: data.Branche_id, CreatedBy: items.CreatedBy}
		config.DB.Create(&traitement)
	}
	c.JSON(200, items)
}
