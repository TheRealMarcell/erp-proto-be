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
	inventoryCommandQuery "erp-api/internal/modules/inventory/repositories/commands"
	inventoryRepoQuery "erp-api/internal/modules/inventory/repositories/queries"

	inventoryUseCase "erp-api/internal/modules/inventory/usecases"

	saleHandler "erp-api/internal/modules/sale/handlers"
	saleRepoCommand "erp-api/internal/modules/sale/repositories/commands"
	saleRepoQuery "erp-api/internal/modules/sale/repositories/queries"

	saleUseCase "erp-api/internal/modules/sale/usecases"

	itemHandler "erp-api/internal/modules/item/handlers"

	itemRepoCommand "erp-api/internal/modules/item/repositories/commands"
	itemRepoQuery "erp-api/internal/modules/item/repositories/queries"
	itemUseCase "erp-api/internal/modules/item/usecases"

	transactionHandler "erp-api/internal/modules/transaction/handlers"
	transactionRepoCommand "erp-api/internal/modules/transaction/repositories/commands"
	transactionRepoQuery "erp-api/internal/modules/transaction/repositories/queries"

	transactionUseCase "erp-api/internal/modules/transaction/usecases"

	historyHandler "erp-api/internal/modules/history/handlers"
	historyRepoCommand "erp-api/internal/modules/history/repositories/command"
	historyRepoQuery "erp-api/internal/modules/history/repositories/queries"

	historyUseCase "erp-api/internal/modules/history/usecases"

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
	db_logger := configuration.Logger()
	logger := log.GetLogger()

	db.InitDB(*db_logger)

	defer db.DB.Close()

	docs.SwaggerInfo.Title = "Sandbox API - IDP BR Service"
	docs.SwaggerInfo.Description = "This is a sandbox API for IDP BR service used for development purposes"
	docs.SwaggerInfo.Version = "1.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/repository"
	docs.SwaggerInfo.Schemes = []string{"http"}

	server := setupRouter()

	// initialise swagger documentation url
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoutes := server.Group("/api")
	routes.RegisterRoutes(apiRoutes)

	historyQueryPostgresRepo := historyRepoQuery.NewQueryPostgresRepository(db.DB, logger)
	historyCommandPostgresRepo := historyRepoCommand.NewCommandPostgresRepository(db.DB, logger)
	historyUseCaseCommand := historyUseCase.NewCommandUsecase(historyCommandPostgresRepo, logger)
	historyUseCaseQuery := historyUseCase.NewQueryUsecase(historyQueryPostgresRepo, logger)

	historyHandler.InitHistoryHttpHandler(server, historyUseCaseQuery, historyUseCaseCommand, logger)

	inventoryQueryPostgresRepo := inventoryRepoQuery.NewQueryPostgresRepository(db.DB, logger)
	inventoryCommandPostgresRepo := inventoryCommandQuery.NewCommandPostgresRepository(db.DB, logger)
	inventoryUseCaseQuery := inventoryUseCase.NewQueryUsecase(inventoryQueryPostgresRepo, logger)
	inventoryUseCaseCommand := inventoryUseCase.NewCommandUsecase(inventoryCommandPostgresRepo, historyCommandPostgresRepo, logger)

	inventoryHandler.InitInventoryHttpHandler(server, inventoryUseCaseQuery, inventoryUseCaseCommand, logger)

	saleQueryPostgresRepo := saleRepoQuery.NewQueryPostgresRepository(db.DB, logger)
	saleCommandPostgresRepo := saleRepoCommand.NewCommandPostgresRepository(db.DB, logger)
	saleUseCase := saleUseCase.NewQueryUsecase(saleQueryPostgresRepo, logger)

	saleHandler.InitSaleHttpHandler(server, saleUseCase, logger)

	itemQueryPostgresRepo := itemRepoQuery.NewQueryPostgresRepository(db.DB, logger)
	itemCommandPostgresRepo := itemRepoCommand.NewCommandPostgresRepository(db.DB, logger)
	itemUseCaseQuery := itemUseCase.NewQueryUsecase(itemQueryPostgresRepo, logger)
	itemUseCaseCommand := itemUseCase.NewCommandUsecase(itemCommandPostgresRepo, inventoryCommandPostgresRepo, saleCommandPostgresRepo, logger)

	itemHandler.InitItemHttpHandler(server, itemUseCaseQuery, itemUseCaseCommand, logger)

	transactionQueryPostgresRepo := transactionRepoQuery.NewQueryPostgresRepository(db.DB, logger)
	transactionCommandPostgresRepo := transactionRepoCommand.NewCommandPostgresRepository(db.DB, logger)
	transactionUseCaseQuery := transactionUseCase.NewQueryUsecase(transactionQueryPostgresRepo, logger)
	transactionUseCaseCommand := transactionUseCase.NewCommandUsecase(transactionCommandPostgresRepo, saleCommandPostgresRepo, inventoryCommandPostgresRepo, logger)

	transactionHandler.InitTransactionHttpHandler(server, transactionUseCaseQuery, transactionUseCaseCommand, logger)

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
