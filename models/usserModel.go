package models

import "gorm.io/gorm"

type Users struct {
	Id uint `json:"id" gorm:"primary_key"`
	gorm.Model
	Names      string `json:"names"`
	Email      string `json:"email" gorm:"unique"`
	Password   []byte `json:"-"`
	Code       string `json:"code"`
	Phone      int    `json:"phone"`
	Address    string `json:"address"`
	Gender     int    `json:"gender"`
	Etat_civil int    `json:"etat_civil"`
	Status     int    `json:"status"`
	Role       int    `json:"role"` //1.Admin,2.chef de depot, 3.Gerant,4.Agronomme,5.Financier
	CreatedBy  int    `json:"createdBy"`
}
