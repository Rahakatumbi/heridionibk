package controller

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
)

func Axe(c *gin.Context) {
	var axe models.Axes
	c.ShouldBindJSON(&axe)
	config.DB.Save(&axe)
	c.JSON(200, axe)
}
func Axes(c *gin.Context) {
	var axes []models.Axes
	config.DB.Find(&axes)
	c.JSON(200, axes)
}
