package depensedocs

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func CreatePrice(c *gin.Context) {
	data := models.Prices{}
	c.ShouldBindJSON(&data)
	config.DB.Save(&data)
	c.JSON(200, data)
}
func GetPriceList(c *gin.Context) {
	var data []models.Prices
	config.DB.Find(&data)
	c.JSON(200, data)
}
func GetLastActivePrice(c *gin.Context) {
	var data models.Prices
	config.DB.Where("status", "Active").Last(&data).Find(&data)
	c.JSON(200, data)
}
