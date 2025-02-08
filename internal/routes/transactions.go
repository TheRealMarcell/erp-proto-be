package routes

import (
	"erp-api/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTransaction(ctx *gin.Context){
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)

	if err != nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sales request"})
		return
	}

	err = transaction.Save()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully saved transaction"})

}