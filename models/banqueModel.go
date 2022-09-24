package models

import "gorm.io/gorm"

type Banques struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	Names         string `json:"names"`
	Rccm          string `json:"rccm"`
	Identifcation string `json:"identification"`
	Compte        string `json:"compte"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Telephone     int    `json:"phone"`
	Status        int    `json:"status"`
	CreatedBy     int    `json:"creator"`
	gorm.Model
}
