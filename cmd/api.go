package main

import (
	db "erp-api/database"
	docs "erp-api/docs"
	"erp-api/internal/pkg/log"
	"net/http"
	"time"

	"erp-api/internal/pkg/util/configuration"
	"erp-api/internal/routes"

	inventoryHandler "erp-api/internal/modules/inventory/handlers"
	inventoryRepoQuery "erp-api/internal/modules/inventory/repositories/queries"

	inventoryUseCase "erp-api/internal/modules/inventory/usecases"

	saleHandler "erp-api/internal/modules/sale/handlers"
	saleRepoQuery "erp-api/internal/modules/sale/repositories/queries"

	saleUseCase "erp-api/internal/modules/sale/usecases"

	itemHandler "erp-api/internal/modules/item/handlers"

	itemRepoQuery "erp-api/internal/modules/item/repositories/queries"
	itemUseCase "erp-api/internal/modules/item/usecases"

	transactionHandler "erp-api/internal/modules/transaction/handlers"
	transactionRepoQuery "erp-api/internal/modules/transaction/repositories/queries"

	transactionUseCase "erp-api/internal/modules/transaction/usecases"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sandbox API - ERP Prototype Service
// @version 1.1
// @description This is a sandbox API for ERP Prototype Service used for development purposes

// @contact.name Marcellus Simanjuntak
// @contact.email marcellusgerson@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.AllowAll())
	return r

}

func main() {
	logger := configuration.Logger()

	logger_log := log.GetLogger()

	db.InitDB(*logger)

	defer db.DB.Close()

	docs.SwaggerInfo.Title = "Sandbox API - IDP BR Service"
	docs.SwaggerInfo.Description = "This is a sandbox API for IDP BR service used for development purposes"
	docs.SwaggerInfo.Version = "1.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/repository"
	docs.SwaggerInfo.Schemes = []string{"http"}

	server := setupRouter()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoutes := server.Group("/api")
	routes.RegisterRoutes(apiRoutes)

	inventoryQueryPostgresRepo := inventoryRepoQuery.NewQueryPostgresRepository(db.DB, logger_log)
	inventoryUseCase := inventoryUseCase.NewQueryUsecase(inventoryQueryPostgresRepo, logger_log)

	inventoryHandler.InitInventoryHttpHandler(server, inventoryUseCase, logger_log)

	saleQueryPostgresRepo := saleRepoQuery.NewQueryPostgresRepository(db.DB, logger_log)
	saleUseCase := saleUseCase.NewQueryUsecase(saleQueryPostgresRepo, logger_log)

	saleHandler.InitSaleHttpHandler(server, saleUseCase, logger_log)

	itemQueryPostgresRepo := itemRepoQuery.NewQueryPostgresRepository(db.DB, logger_log)
	itemUseCase := itemUseCase.NewQueryUsecase(itemQueryPostgresRepo, logger_log)

	itemHandler.InitItemHttpHandler(server, itemUseCase, logger_log)

	transactionQueryPostgresRepo := transactionRepoQuery.NewQueryPostgresRepository(db.DB, logger_log)
	transactionUseCase := transactionUseCase.NewQueryUsecase(transactionQueryPostgresRepo, logger_log)

	transactionHandler.InitTransactionHttpHandler(server, transactionUseCase, logger_log)

	httpServer := &http.Server{
		Addr:           ":8080",
		Handler:        server,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpServer.ListenAndServe()
}
