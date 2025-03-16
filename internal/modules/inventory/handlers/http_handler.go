package handlers

import (
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/inventory/models/request"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InventoryHttpHandler struct {
	InventoryUsecaseQuery   inventory.UsecaseQuery
	InventoryUsecaseCommand inventory.UsecaseCommand
	Validator               *validator.Validate
	Logger                  log.Logger
}

func InitInventoryHttpHandler(app *gin.Engine, auq inventory.UsecaseQuery, auc inventory.UsecaseCommand, log log.Logger) {
	handler := &InventoryHttpHandler{
		InventoryUsecaseQuery:   auq,
		InventoryUsecaseCommand: auc,
		Logger:                  log,
		Validator:               validator.New(),
	}

	route := app.Group("/api/inventory")
	route.GET("/:location", handler.GetInventory)
	route.POST("", handler.MoveInventory)
}

// MoveInventory godoc
// @Summary      Move item from one inventory to another (pindahan barang)
// @Description  Given an list of items in a particular inventory, move that set of items into a different inventory and update the quantities
// @Tags         Inventory
// @Accept       json
// @Produce      json
// @Param        request body request.MoveInventory true "Request payload"
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/inventory/:location [get]
func (i InventoryHttpHandler) MoveInventory(ctx *gin.Context) {
	req := new(request.MoveInventory)
	if err := ctx.ShouldBindJSON(req); err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not parse request", err)
		return
	}

	// move inventory
	if err := i.InventoryUsecaseCommand.MoveInventory(ctx, *req); err != nil {
		fmt.Println(err)
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not move items", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "sucessfully moved inventory", nil)
}

// GetInventory godoc
// @Summary      Get a list of items in inventory
// @Description  Given an input location, return the items that exist in the inventory along with its quantity
// @Tags         Inventory
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/inventory/:location [get]
func (i InventoryHttpHandler) GetInventory(ctx *gin.Context) {
	location := ctx.Param("location") // Get location from URL param

	resp, err := i.InventoryUsecaseQuery.GetInventory(ctx, location)
	if err != nil {
		msg := fmt.Sprintf("could not get items in inventory %v", location)
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, msg, err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "Inventory fetched successfully", resp)
}
