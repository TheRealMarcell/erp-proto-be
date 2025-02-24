package main

import (
	db "erp-api/database"
	docs "erp-api/docs"
	"erp-api/internal/routes"

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
	r.Use(cors.AllowAll(),)
	return r

}

func main() {
	db.InitDB()

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
	server.Run(":8080")
}