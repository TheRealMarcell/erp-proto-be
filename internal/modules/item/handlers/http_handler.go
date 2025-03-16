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
	route.PUT("/:id", handler.correctItem)
	route.PUT("/rusak", handler.brokenItem)
	route.PUT("/price", handler.updateItemPrice)
}

// GetItems godoc
// @Summary      Get items
// @Description  Get a list of items stored in the repository
// @Tags         Items
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items [get]
func (i ItemHttpHandler) getItems(ctx *gin.Context) {
	resp, err := i.ItemUsecaseQuery.GetItems(ctx)
	if err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not get items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "items fetched successfully", resp)
}

// CreateItem godoc
// @Summary      Create new item (terima barang)
// @Description  Create and save a new item, along with an initial quantity that is stored in the inventory (gudang)
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        request body request.SubmitItem true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items [post]
func (i ItemHttpHandler) createItem(ctx *gin.Context) {
	req := new(request.SubmitItem)
	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.SaveItem(ctx, *req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not save items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully added items", nil)
}

// UpdateItem godoc
// @Summary      Update an item (retur barang)
// @Description  Return an item into the inventory, update sale retur quantity
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        request body request.UpdateItem true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items [put]
func (i ItemHttpHandler) updateItem(ctx *gin.Context) {
	req := new(request.UpdateItem)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.UpdateItem(ctx, *req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not update items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated items", nil)
}

// CorrectItem godoc
// @Summary      Correct an item (koreksi barang)
// @Description  Correct an existing item's attributes
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        request body request.CorrectItem true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items/:id [put]
func (i ItemHttpHandler) correctItem(ctx *gin.Context) {
	id := ctx.Param("id")

	req := new(request.CorrectItem)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.CorrectItem(ctx, *req, id); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not correct item", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)
}

// BrokenItem godoc
// @Summary      Update an item as broken (retur barang rusak)
// @Description  Return an item and add into inventory rusak
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        request body request.BrokenItem true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items/rusak [put]
func (i ItemHttpHandler) brokenItem(ctx *gin.Context) {
	req := new(request.UpdateItem)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.BrokenItem(ctx, *req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not update broken items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "broken items added", nil)

}

// UpdateItemPrice godoc
// @Summary      Update item price
// @Description  Update the price of an item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        request body request.ItemPrice true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/items/price [put]
func (i ItemHttpHandler) updateItemPrice(ctx *gin.Context) {
	req := new(request.ItemPrice)

	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusBadRequest, "could not parse request", err)
		return
	}

	if err := i.Validator.Struct(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "validator error", err)
		return
	}

	if err := i.ItemUsecaseCommand.UpdateItemPrice(ctx, *req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not update item price", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "item price updated", nil)

}
