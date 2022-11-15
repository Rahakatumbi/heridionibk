package api

import (
	"github.com/Raha2071/heridionibd/controller"
	"github.com/Raha2071/heridionibd/controller/achats"
	"github.com/Raha2071/heridionibd/controller/banque"
	"github.com/Raha2071/heridionibd/controller/clients"
	"github.com/Raha2071/heridionibd/controller/depensedocs"
	"github.com/Raha2071/heridionibd/controller/products"
	"github.com/Raha2071/heridionibd/controller/rapports"
	"github.com/Raha2071/heridionibd/controller/ventes"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.RouterGroup) {
	admin := app.Group("/")
	{
		admin.POST("/users", controller.Register)
		admin.GET("/users", controller.Users)
		admin.POST("/branche", controller.Branche)
		admin.GET("/branche", controller.Branches)
		admin.GET("/financement/:id", controller.FinancementByDepot)
		admin.GET("/financements", controller.AllFinancement)
		admin.POST("financement", controller.FinancementDepot)
		admin.GET("balance/:id", controller.Finance)
		admin.GET("/depotData/:id", controller.DepotData)
		admin.POST("/achat", achats.Achat)
		admin.GET("/achat", achats.Achats)
		admin.GET("/order", ventes.Orders)
		admin.POST("/order", ventes.Order)
		admin.POST("/facturation", ventes.Facturation)
		admin.GET("/order/:id", ventes.OrderInfos)
		admin.POST("/solveOrder", ventes.SolveOrder)
		admin.POST("/product", products.Product)
		admin.GET("/product", products.Products)
		admin.POST("/banque", banque.Banque)
		admin.GET("/banque", banque.Banques)
		admin.POST("/client", clients.Client)
		admin.GET("/client", clients.Clients)
		admin.POST("/axe", controller.Axe)
		admin.GET("/axe", controller.Axes)
		admin.POST("/champ", clients.AddChamp)
		admin.GET("/champ", clients.AllChamps)
		admin.GET("/champ/:id", clients.Champs)
		admin.GET("/supllier", clients.Suppliers)
		admin.POST("/supllier", clients.Supllier)
		admin.POST("/arapport", rapports.Approvisionnement)
		admin.POST("/epport", rapports.Facturation)
		admin.GET("/supllier/:id", clients.Search)
		admin.POST("/login", controller.Login)
		admin.POST("/depense", depensedocs.Depense)
		admin.GET("/depense", depensedocs.Depenses)
		admin.POST("/document", depensedocs.Document)
		admin.GET("/document", depensedocs.Documents)
		admin.GET("/doneorders", ventes.TraitOrders)
	}
}
