package config

import (
	"fmt"
	"os"

	"github.com/Raha2071/heridionibd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DATABASE")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	// dsn := "root:@tcp(127.0.0.1:3306)/compusdb?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//mysqlite
	// db, err := gorm.Open(sqlite.Open("coffeedb.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	DB = db
	db.AutoMigrate(&models.Users{}, models.Achats{}, models.Banques{}, models.Branche{},
		models.Clients{}, models.Clients{}, models.Invoice{}, models.InvoiceInfo{}, models.Products{},
		models.Suppliers{}, models.Champs{}, models.Axes{}, models.ServedOrder{}, models.ServedOrderInfo{},
		models.Orders{}, models.OrdersInfo{}, models.AchatsInfos{}, models.FinancementOrder{},
	)
	db.AutoMigrate(&models.AchatsInfos{}, models.TraitementInfo{})
	db.AutoMigrate(&models.FinancementDepot{}, models.FinancementInfo{}, models.Depenses{}, models.Documents{})

}
