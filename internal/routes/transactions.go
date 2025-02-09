package routes

import (
	"erp-api/internal/model"
	"erp-api/util/httpres"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllTransactions(ctx *gin.Context){
	transactions, err := model.GetTransactions()
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not get transaction", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", transactions)
}

func createTransaction(ctx *gin.Context){
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse sales request", nil)
		return
	}

	err = transaction.Save()
	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save transaction", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", "successfully saved transaction")
}