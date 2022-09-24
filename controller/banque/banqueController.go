package banque

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Banque(c *gin.Context) {
	var banque models.Banques
	c.ShouldBindJSON(&banque)
	config.DB.Save(&banque)
	c.JSON(200, banque)
}
func Banques(c *gin.Context) {
	var banque []models.Banques
	config.DB.Find(&banque)
	c.JSON(200, banque)
}
