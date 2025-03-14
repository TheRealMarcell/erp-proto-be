package routes

import (
	"erp-api/internal/model"
	"erp-api/internal/pkg/util/httpres"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func updatePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	var payment_status struct {
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
