package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Loginform struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var SecretKey = "Uwike"

func ClaimToken(userId int) (string, error) {
	var err error
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
func Login(c *gin.Context) {
	var loghelper Loginform
	c.ShouldBindJSON(&loghelper)
	user := models.Users{}
	config.DB.Where("code", loghelper.Username).Last(&user)
	if user.Id == 0 {
		c.JSON(401, gin.H{
			"message": "Code d'Utilisateur ne Correspond pas.",
		})
	} else {
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(loghelper.Password)); err != nil {
			c.JSON(401, gin.H{
				"message": "Le mot de passe est Incorrect.",
			})
		} else {
			token, err := ClaimToken(int(user.Id))
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				// return
			}
			if user.Role == 2 {
				var branche models.Branche
				config.DB.Find(&branche, "chef_id=?", user.Id)
				c.JSON(http.StatusOK, gin.H{
					"user":    user,
					"token":   token,
					"branche": branche,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"user":  user,
					"token": token})
			}
		}
	}
}
