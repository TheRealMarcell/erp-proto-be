package handlers

import (
	"erp-api/internal/modules/item"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ItemHttpHandler struct {
	ItemUsecaseQuery item.UsecaseQuery
	Validator        *validator.Validate
	Logger           log.Logger
}

func InitItemHttpHandler(app *gin.Engine, auq item.UsecaseQuery, log log.Logger) {
	handler := &ItemHttpHandler{
		ItemUsecaseQuery: auq,
		Logger:           log,
		Validator:        validator.New(),
	}

	route := app.Group("/api/items")
	route.GET("", handler.GetItems)
}

func (i ItemHttpHandler) GetItems(ctx *gin.Context) {
	resp, err := i.ItemUsecaseQuery.GetItems(ctx)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not get items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "Items fetched successfully", resp)
}
