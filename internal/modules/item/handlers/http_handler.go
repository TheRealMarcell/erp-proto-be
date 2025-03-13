package handlers

import (
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ItemHttpHandler struct {
	ItemUsecaseQuery   item.UsecaseQuery
	ItemUsecaseCommand item.UsecaseCommand
	Validator          *validator.Validate
	Logger             log.Logger
}

func InitItemHttpHandler(app *gin.Engine, auq item.UsecaseQuery, auc item.UsecaseCommand, log log.Logger) {
	handler := &ItemHttpHandler{
		ItemUsecaseQuery:   auq,
		ItemUsecaseCommand: auc,
		Logger:             log,
		Validator:          validator.New(),
	}

	route := app.Group("/api/items")
	route.GET("", handler.getItems)

	route.POST("", handler.createItem)

	route.PUT("", handler.updateItem)
	// route.PUT("/price", handler.updateItemPrice)
	// route.PUT("/:id", handler.correctItem)
	// route.PUT("/rusak", handler.brokenItem)
}

func (i ItemHttpHandler) getItems(ctx *gin.Context) {
	resp, err := i.ItemUsecaseQuery.GetItems(ctx)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not get items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "items fetched successfully", resp)
}

func (i ItemHttpHandler) createItem(ctx *gin.Context) {
	req := new(request.SubmitItem)
	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.SaveItem(ctx, *req); err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully added items", nil)
}

func (i ItemHttpHandler) updateItem(ctx *gin.Context) {
	req := new(request.UpdateItem)
	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", err)
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIResponse(ctx, http.StatusBadRequest, "validator error", err)
	}

	if err := i.ItemUsecaseCommand.UpdateItem(ctx, *req); err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated items", nil)
}

func (i ItemHttpHandler) correctItem(ctx *gin.Context) {}

func (i ItemHttpHandler) brokenItem(ctx *gin.Context) {}

func (i ItemHttpHandler) updateItemPrice(ctx *gin.Context) {}
