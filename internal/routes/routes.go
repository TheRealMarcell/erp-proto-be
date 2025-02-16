package routes

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.RouterGroup){
	server.POST("/verify-user", verifyUserByPassword)

	server.GET("/sales", getAllSales)

	server.GET("/transactions", getAllTransactions)
	server.POST("/transactions", createTransaction)

	server.GET("/items", getAllItems)
	server.POST("/items", createItem)

	server.POST("/items/:id", insertItem)
	server.PUT("/items/:id", updateItem) // retur barang

	server.POST("/inventory/:id", moveItem)
	server.GET("/inventory/:location", getInventory)
}