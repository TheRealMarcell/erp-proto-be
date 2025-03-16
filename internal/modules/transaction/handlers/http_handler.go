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

// GetTransactions godoc
// @Summary      Get Transactions
// @Description  Get a list of transactions
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/transactions [get]
func (t TransactionHttpHandler) GetTransactions(ctx *gin.Context) {
	resp, err := t.TransactionUsecaseQuery.GetTransactions(ctx)
	if err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not get transaction", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", resp)
}

// GetTransactionDiscount godoc
// @Summary      Get Transaction Discounts
// @Description  Get a list of transactions with their corresponding discount amount
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/transactions/discount_percent [get]
func (t TransactionHttpHandler) GetTransactionDiscount(ctx *gin.Context) {
	discount_percent, err := t.TransactionUsecaseQuery.GetDiscountPercentages(ctx)
	if err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not fetch discount percent", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", discount_percent)

}

// CreateTransaction godoc
// @Summary      Create Transaction
// @Description  Save transaction to the database
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body request.Transaction true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/transactions [post]
func (t TransactionHttpHandler) CreateTransaction(ctx *gin.Context) {
	req := new(request.Transaction)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := t.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := t.TransactionUsecaseCommand.InsertTransaction(ctx, *req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not insert transaction", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully added transaction", nil)
}

// CreateTransaction godoc
// @Summary      Update Payment
// @Description  Update the payment method associated with a transaction
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body string true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/transactions/payment/:id [post]
func (t TransactionHttpHandler) UpdatePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	var payment_status struct {
		PaymentStatus string `json:"payment_status"`
	}

	if err := ctx.ShouldBindJSON(&payment_status); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request data", err)
		return
	}

	if err := t.TransactionUsecaseCommand.UpdatePaymentStatus(ctx, id, payment_status.PaymentStatus); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not update payment status", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated payment", nil)
}
