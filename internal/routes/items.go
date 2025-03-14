package routes

import (
	"erp-api/internal/model"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
