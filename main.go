package main

import (
	"github.com/Raha2071/heridionibd/api"
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/infrastructure"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	app := gin.Default()
	app.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		// MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
		// Authorization: false,

	}))
	router := app.Group("/api/v1")
	api.Setup(router)
	infrastructure.LoadEnv()
	// infrastructure.NewDatabase() //new database connection
	config.SetupDB()

	app.Run(":5004")
	// app.Run(":5004")
	// app.Run("192.168.43.174:5004")
}
