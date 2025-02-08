package routes

import (
	"erp-api/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func getAllSales(ctx *gin.Context){
	sales, err := model.GetSales()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch sales"})
	}

	ctx.JSON(http.StatusOK, sales)
}

func createNewSale(ctx *gin.Context){
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)

	if err != nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sales request"})
		return
	}

	// fmt.Println(transaction.Sales)

	// for _, sale := range transaction.Sales{
	// 	err = sale.Save()

	// 	if err != nil{
	// 		fmt.Println(err)
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not save sale"})
	// 		return
	// 	}

	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully created sale",
	})

}
