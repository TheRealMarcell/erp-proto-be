package handlers

import (
	"erp-api/internal/modules/sale"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaleHttpHandler struct {
	SaleUsecaseQuery sale.UsecaseQuery
	Validator        *validator.Validate
	Logger           log.Logger
}

func InitSaleHttpHandler(app *gin.Engine, auq sale.UsecaseQuery, log log.Logger) {
	handler := &SaleHttpHandler{
		SaleUsecaseQuery: auq,
		Logger:           log,
		Validator:        validator.New(),
	}

	route := app.Group("/api/sales")
	route.GET("", handler.GetListSales)
}

// GetListSales godoc
// @Summary      Get sales
// @Description  Get a list of sales
// @Tags         Sales
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/sales [get]
func (s SaleHttpHandler) GetListSales(ctx *gin.Context) {
	resp, err := s.SaleUsecaseQuery.GetSales(ctx)
	if err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not get list of sales", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "sale successfully fetched", resp)
}
