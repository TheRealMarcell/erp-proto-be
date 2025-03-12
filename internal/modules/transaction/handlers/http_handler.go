package handlers

import (
	"erp-api/internal/modules/transaction"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionHttpHandler struct {
	TransactionUsecaseQuery transaction.UsecaseQuery
	Validator               *validator.Validate
	Logger                  log.Logger
}

func InitTransactionHttpHandler(app *gin.Engine, auq transaction.UsecaseQuery, log log.Logger) {
	handler := &TransactionHttpHandler{
		TransactionUsecaseQuery: auq,
		Logger:                  log,
		Validator:               validator.New(),
	}

	route := app.Group("/api/transactions")
	route.GET("", handler.GetTransactions)
	route.GET("/discount_percent", handler.GetTransactionDiscount)
}

func (t TransactionHttpHandler) GetTransactions(ctx *gin.Context) {
	resp, err := t.TransactionUsecaseQuery.GetTransactions(ctx)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not get transaction", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", resp)
}

func (t TransactionHttpHandler) GetTransactionDiscount(ctx *gin.Context) {
	discount_percent, err := t.TransactionUsecaseQuery.GetDiscountPercentages(ctx)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not fetch discount percent", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", discount_percent)

}
