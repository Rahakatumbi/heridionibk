package controller

import (
	"strconv"

	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.Users
	//save university
	config.DB.Select("id").Last(&user)
	pass := "HR" + user.Code + ""
	code := pass + "D" + strconv.Itoa(int(user.Id))

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	cearetuser := models.Users{Password: password, Code: code}
	c.ShouldBindJSON(&cearetuser)
	config.DB.Save(&cearetuser)
	c.JSON(200, cearetuser)
}
func Users(c *gin.Context) {
	var users []models.Users
	config.DB.Find(&users)
	// for user,_ :=range users{

	// }
	c.JSON(200, users)
}
