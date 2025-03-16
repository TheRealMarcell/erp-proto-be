package handlers

import (
	"erp-api/internal/modules/history"
	"erp-api/internal/pkg/log"
	"erp-api/internal/pkg/util/httpres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HistoryHttpHandler struct {
	HistoryUsecaseQuery   history.UsecaseQuery
	HistoryUsecaseCommand history.UsecaseCommand
	Validator             *validator.Validate
	Logger                log.Logger
}

func InitHistoryHttpHandler(app *gin.Engine, auq history.UsecaseQuery, auc history.UsecaseCommand, log log.Logger) {
	handler := &HistoryHttpHandler{
		HistoryUsecaseQuery:   auq,
		HistoryUsecaseCommand: auc,
		Validator:             validator.New(),
		Logger:                log,
	}

	route := app.Group("/api/history")
	route.GET("", handler.getHistory)
}

// GetHistory godoc
// @Summary      Get history
// @Description  Get a history of past item moves from different inventories
// @Tags         History
// @Accept       json
// @Produce      json
// @Success      200 {object} httpres.HTTPResponse
// @Failure      400 {object} httpres.HTTPError
// @Failure      500 {object} httpres.HTTPError
// @Router       /api/history [get]
func (h HistoryHttpHandler) getHistory(ctx *gin.Context) {
	resp, err := h.HistoryUsecaseQuery.GetHistory(ctx)

	if err != nil {
		httpres.APIErrorResponse(ctx, http.StatusInternalServerError, "could not fetch history", err)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully fetched history", resp)
}
