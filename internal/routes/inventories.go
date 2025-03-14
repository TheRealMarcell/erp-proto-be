package routes

import (
	"erp-api/internal/model"
	"erp-api/internal/pkg/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func moveInventory(ctx *gin.Context) {
	var inventory_request model.InventoryRequest

	err := ctx.ShouldBindJSON(&inventory_request)
	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, inventory_item := range inventory_request.Items {
		inventory_item.Destination = inventory_request.Destination
		inventory_item.Source = inventory_request.Source

		err = inventory_item.MoveInventory()
		if err != nil {
			httpres.APIResponse(ctx, http.StatusInternalServerError, "could not move item", nil)
			return
		}
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully moved", nil)
}
