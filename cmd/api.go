package main

import (
	db "erp-api/database"
	"erp-api/internal/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.AllowAll(),)
	return r

}

func main() {
	db.InitDB()

	server := setupRouter()
	
	apiRoutes := server.Group("/api")
	routes.RegisterRoutes(apiRoutes)
	server.Run(":8080")
}