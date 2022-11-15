package depensedocs

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

type DepensesHelper struct {
	Id          uint    `json:"id"`
	Montant     float32 `json:"montant"`
	Motif       string  `json:"motif"`
	Description string  `json:"description"`
	Createdby   int     `json:"*"`
	CreatedAt   string  `json:"created_at"`
	User        string  `json:"names"`
}

func Depense(c *gin.Context) {
	var dps models.Depenses
	c.ShouldBindJSON(&dps)
	config.DB.Save(&dps)
	c.JSON(200, dps)
}
func Depenses(c *gin.Context) {
	var dp []DepensesHelper
	var depenses []models.Depenses
	config.DB.Select("depenses.id,depenses.montant,depenses.motif,depenses.description,depenses.createdby,depenses.created_at,users.id,users.names,users.code").
		Joins("inner join users on users.id = depenses.createdby").Find(&depenses).Scan(&dp)
	if len(dp) > 0 {
		c.JSON(200, dp)
	} else {
		c.JSON(200, depenses)
	}
}
