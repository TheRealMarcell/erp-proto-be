package routes

import (
	"erp-api/internal/model"
	"erp-api/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func moveInventory(ctx *gin.Context){
	var inventory_request model.InventoryRequest

	err := ctx.ShouldBindJSON(&inventory_request)
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, inventory_item := range inventory_request.Items {
		inventory_item.Destination = inventory_request.Destination
		inventory_item.Source = inventory_request.Source

		fmt.Println(inventory_item)

		err = inventory_item.MoveInventory()
		if err != nil{
			httpres.APIResponse(ctx, http.StatusInternalServerError, "could not move item", nil)
			return
		}
	}

	httpres.APIResponse(ctx, http.StatusOK, "successly moved", nil)
}

func getInventory(ctx *gin.Context) {
	location := ctx.Param("location") // Get location from URL param

	fmt.Println("Fetching inventory for location:", location)

	inventory, err := model.GetInventory(location)
	if err != nil {
		fmt.Println("Error fetching inventory:", err)
		httpres.APIResponse(ctx, http.StatusBadRequest, fmt.Sprintf("Invalid location: %s", location), nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "Inventory fetched successfully", inventory)
}