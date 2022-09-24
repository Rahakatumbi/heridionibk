package find

import (
	"errors"

	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
)

func FindUser(id int, user *models.Users) error {
	config.DB.Find(&user, id)
	if user.ID == 0 {
		return errors.New("Introuvable")
	}
	return nil
}
