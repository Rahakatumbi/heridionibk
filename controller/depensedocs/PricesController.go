package depensedocs

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type ResponsePrices struct {
	Id        uint    `json:"id" gorm:"primary_id"`
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
	Status    string  `json:"status"`
	CreatedBy int     `json:"createdby"`
	Names     string  `json:"names"`
}

func CreatePrice(c *gin.Context) {
	data := models.Prices{}
	c.ShouldBindJSON(&data)
	config.DB.Save(&data)
	c.JSON(200, data)
}
func GetPriceList(c *gin.Context) {
	var products []ResponsePrices
	var data []models.Prices
	config.DB.Select("prices.id,prices.price,prices.status,prices.created_by,prices.product_id as product_id,products.id as pro,products.names").
		Joins("inner join products on products.id = prices.product_id").Find(&data).Scan(&products)
	if len(products) > 0 {
		c.JSON(200, products)
	} else {
		c.JSON(200, data)
	}
}
func GetLastActivePrice(c *gin.Context) {
	var products []ResponsePrices
	var data []models.Prices
	config.DB.Select("prices.id,prices.price,prices.status,prices.created_by,prices.product_id as product_id,products.id as pro,products.names").
		Joins("inner join products on products.id = prices.product_id").Where("status", "Active").Find(&data).Scan(&products)
	if len(products) > 0 {
		c.JSON(200, products)
	} else {
		c.JSON(200, data)
	}
}
