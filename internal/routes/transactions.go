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

func updatePayment(ctx *gin.Context){
	id := ctx.Param("id")
	
	var payment_status struct{
		PaymentStatus string `json:"payment_status"`
	}

	err := ctx.ShouldBindJSON(&payment_status)

	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request data", nil)
		return
	}

	err = model.UpdatePaymentStatus(id, payment_status.PaymentStatus)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save payment status", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", "successfully updated transaction payment status")

}