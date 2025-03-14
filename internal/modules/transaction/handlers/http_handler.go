package handlers

import (
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/request"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionHttpHandler struct {
	TransactionUsecaseQuery   transaction.UsecaseQuery
	TransactionUsecaseCommand transaction.UsecaseCommand
	Validator                 *validator.Validate
	Logger                    log.Logger
}

func InitTransactionHttpHandler(app *gin.Engine, auq transaction.UsecaseQuery, auc transaction.UsecaseCommand, log log.Logger) {
	handler := &TransactionHttpHandler{
		TransactionUsecaseQuery:   auq,
		TransactionUsecaseCommand: auc,
		Logger:                    log,
		Validator:                 validator.New(),
	}

	route := app.Group("/api/transactions")
	route.GET("", handler.GetTransactions)
	route.GET("/discount_percent", handler.GetTransactionDiscount)

	route.POST("", handler.CreateTransaction)
	route.PUT("/payment/:id", handler.UpdatePayment)
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

func (t TransactionHttpHandler) CreateTransaction(ctx *gin.Context) {
	req := new(request.Transaction)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := t.Validator.Struct(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := t.TransactionUsecaseCommand.InsertTransaction(ctx, *req); err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not insert transaction", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully added transaction", nil)
}

func (t TransactionHttpHandler) UpdatePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	var payment_status struct {
		PaymentStatus string `json:"payment_status"`
	}

	if err := ctx.ShouldBindJSON(&payment_status); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request data", nil)
		return
	}

	if err := t.TransactionUsecaseCommand.UpdatePaymentStatus(ctx, id, payment_status.PaymentStatus); err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update payment status", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated payment", nil)
}
