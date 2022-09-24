package products

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Product(c *gin.Context) {
	var Product models.Products
	c.ShouldBindJSON(&Product)
	config.DB.Save(&Product)
	c.JSON(200, Product)
}
func Products(c *gin.Context) {
	var products []models.Products
	config.DB.Find(&products)
	c.JSON(200, products)
}
