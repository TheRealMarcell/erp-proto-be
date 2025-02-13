package routes

import (
	"encoding/json"
	"erp-api/internal/model"
	"erp-api/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func moveItem(ctx *gin.Context){
	id := ctx.Param("id")

	var inventory_move model.InventoryRequest

	err := json.NewDecoder(ctx.Request.Body).Decode(&inventory_move)

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}
	inventory_move.ItemID = id

	err = inventory_move.MoveItem()
	if err != nil{
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not move item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successly moved", nil)
}