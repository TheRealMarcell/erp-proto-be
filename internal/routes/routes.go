package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.RouterGroup) {
	server.POST("/items", createItem)     // terima barang
	server.PUT("/items", updateItem)      // retur barang
	server.PUT("items/:id", correctItem)  // koreksi
	server.PUT("items/rusak", brokenItem) // retur barang rusak
	server.PUT("items/price", updateItemPrice)

	server.POST("/transactions", createTransaction)
	server.PUT("/transactions/payment/:id", updatePayment)

	server.POST("/inventory", moveInventory) // pindahan

	server.POST("/verify-user", verifyUserByPassword)

	server.POST("/items/:id", insertItem) // THIS IS DEFUNCT
}
