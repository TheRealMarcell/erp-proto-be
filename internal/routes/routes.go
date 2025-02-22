package routes

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.RouterGroup){
	server.POST("/verify-user", verifyUserByPassword)

	server.GET("/sales", getAllSales)

	server.GET("/transactions", getAllTransactions)
	server.POST("/transactions", createTransaction)
  server.PUT("/transactions/payment/:id", updatePayment)
  server.GET("/transactions/discount_percent/:id", getTransactionDiscount)

	server.GET("/items", getAllItems)

	server.POST("/items", createItem) // terima barang
  server.PUT("/items", updateItem) // retur barang
  server.PUT("items/:id", correctItem) // koreksi
  server.PUT("items/rusak", brokenItem) // retur barang rusak

  server.PUT("items/price", updateItemPrice)

  server.POST("/inventory", moveInventory)  // pindahan
	server.GET("/inventory/:location", getInventory)

	server.POST("/items/:id", insertItem) // THIS IS DEFUNCT
}