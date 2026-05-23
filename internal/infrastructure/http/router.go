package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"trade-organization/internal/infrastructure/http/handlers"
	"trade-organization/internal/infrastructure/http/middleware"
)

func SetupRouter(
	authHandler *handlers.AuthHandler,
	tradeHandler *handlers.TradeHandler,
	reportHandler *handlers.ReportHandler,
	supplyHandler *handlers.SupplyHandler,
	storeHandler *handlers.StoreHandler,
) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	api := router.Group("/api/v1")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)

		protected := api.Group("/")
		protected.Use(middleware.RequireAuth())
		{
			seller := protected.Group("/trade")
			seller.Use(middleware.RequireRole("role_seller", "role_director"))
			{
				seller.POST("/sales", tradeHandler.CreateSale)
				seller.POST("/transfer", tradeHandler.TransferProduct)
			}

			manager := protected.Group("/supply")
			manager.Use(middleware.RequireRole("role_purchase_manager", "role_director"))
			{
				manager.POST("/requests", supplyHandler.CreateRequest)
				manager.POST("/orders/generate", supplyHandler.GenerateOrder) // <-- ИСПРАВЛЕНО ЗДЕСЬ
				manager.POST("/orders/receive", supplyHandler.ReceiveOrder)

				manager.GET("/order-details", supplyHandler.GetOrderDetails)
			}

			director := protected.Group("/reports")
			director.Use(middleware.RequireRole("role_director"))
			{
				director.GET("/inventory", reportHandler.GetInventory)
				director.GET("/profitability", reportHandler.GetProfitability)
				director.GET("/turnover", reportHandler.GetTurnover)
				director.GET("/efficiency", reportHandler.GetStoreEfficiency)
				director.GET("/supplier-deliveries", reportHandler.GetSupplierDeliveries)
				director.GET("/product-customers", reportHandler.GetProductCustomers)
			}

			director.POST("/stores", storeHandler.CreateStore)
		}
	}

	return router
}
