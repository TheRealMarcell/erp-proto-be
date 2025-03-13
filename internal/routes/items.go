package routes

import (
	"encoding/json"
	"erp-api/internal/model"
	"erp-api/internal/pkg/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func correctItem(ctx *gin.Context) {
	id := ctx.Param("id")

	var item model.StorageItem

	err := json.NewDecoder(ctx.Request.Body).Decode(&item)
	item.ItemID = id

	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	err = item.UpdateItem("set")
	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)
}

func brokenItem(ctx *gin.Context) {
	// update total on sale id, insert to inventory_rusak

	var broken_item model.BrokenItemRequest

	err := ctx.ShouldBindJSON(&broken_item)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, broken_item := range broken_item.Items {
		err = model.SaveBrokenItem(broken_item)
		if err != nil {
			httpres.APIResponse(ctx, http.StatusBadRequest, "could not update database", nil)
			return
		}
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated broken item", nil)
}

func updateItemPrice(ctx *gin.Context) {
	var item_request model.ItemPriceRequest

	err := ctx.ShouldBindJSON(&item_request)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	err = model.UpdatePrice(item_request)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update price", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated price", nil)

}
