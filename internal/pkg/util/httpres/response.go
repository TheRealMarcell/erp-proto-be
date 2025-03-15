package httpres

import (
	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Meta MetaData `json:"meta"`
	Data any      `json:"data"`
}

type HTTPError struct {
	Meta  MetaData `json:"meta"`
	Error string   `json:"error"`
}

type MetaData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type DataObject struct {
}

func APIResponse(ctx *gin.Context, code int, status string, data any) {
	ctx.JSON(code, HTTPResponse{
		Meta: MetaData{
			Code:   code,
			Status: status,
		},
		Data: data,
	})
}

func APIErrorResponse(ctx *gin.Context, code int, status string, error_msg error) {
	ctx.JSON(code, HTTPError{
		Meta: MetaData{
			Code:   code,
			Status: status,
		},
		Error: error_msg.Error(),
	})
}
