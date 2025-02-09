package routes

import (
	"erp-api/internal/model"
	"erp-api/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func getAllSales(ctx *gin.Context){
	sales, err := model.GetSales()
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not fetch sales", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", sales)
}

// func createNewSale(ctx *gin.Context){
// 	var transaction model.Transaction

// 	err := ctx.ShouldBindJSON(&transaction)

// 	if err != nil{
// 		fmt.Println(err)
// 		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse sales request", nil)
// 		return
// 	}

// 	httpres.APIResponse(ctx, http.StatusOK, "successfully created sale", nil)
// }
