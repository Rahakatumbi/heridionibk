package config

import (
	"github.com/Raha2071/heridionibd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	// DB_USER := os.Getenv("DB_USER")
	// DB_PASS := os.Getenv("DB_PASSWORD")
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_DBNAME := os.Getenv("DB_DATABASE")

	// URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?
	dsn := "root:admin@tcp(localhost:3306)/coffeedb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

}
