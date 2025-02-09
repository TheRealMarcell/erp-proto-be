package routes

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.RouterGroup){
	server.POST("/verify-user", verifyUserByPassword)

	server.GET("/sales", getAllSales)
	server.POST("/sales", createNewSale)

	server.GET("/transactions", getAllTransactions)
	server.POST("/transactions", createTransaction)

}