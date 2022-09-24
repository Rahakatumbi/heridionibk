package clients

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Client(c *gin.Context) {
	var client models.Clients
	c.ShouldBindJSON(&client)
	config.DB.Save(&client)
	c.JSON(200, client)
}
func Clients(c *gin.Context) {
	var clients []models.Clients
	config.DB.Find(&clients)
	c.JSON(200, clients)
}
