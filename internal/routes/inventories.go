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

	// qty, err := model.GetItemQty(inventory_move.ItemID, inventory_move.Source)

	if err != nil{
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not find item in storage", nil)
		return
	}

	// source_qty := *qty - inventory_move.Quantity

	var source_inventory_item model.StorageItem
	source_inventory_item.Location = inventory_move.Source
	source_inventory_item.Quantity = 200
	source_inventory_item.ItemID = inventory_move.ItemID
	source_inventory_item.Description = inventory_move.Description

	err = source_inventory_item.UpdateItem()
	if err !=nil{
		httpres.APIResponse(ctx, http.StatusInternalServerError, "failed to deduct item", nil)
		return
	}

	var destination_inventory_item model.StorageItem

	destination_inventory_item.Location = inventory_move.Destination
	destination_inventory_item.Quantity = inventory_move.Quantity
	destination_inventory_item.ItemID = inventory_move.ItemID
	destination_inventory_item.Description = inventory_move.Description

	destination_inventory_item.Save()

	httpres.APIResponse(ctx, http.StatusOK, "successly moved", nil)
}