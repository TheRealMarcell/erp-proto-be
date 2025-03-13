package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.RouterGroup) {
	server.PUT("items/price", updateItemPrice)

	server.POST("/transactions", createTransaction)
	server.PUT("/transactions/payment/:id", updatePayment)

	server.POST("/inventory", moveInventory) // pindahan

	server.POST("/verify-user", verifyUserByPassword)
}
