package httpres

import "github.com/gin-gonic/gin"

type HTTPResponse struct {
	Meta MetaData			`json:"meta"`
	Data any					`json:"data,omitempty"`
}

type MetaData struct {
	Code int				`json:"code"`
	Status string		`json:"status"`
}

func APIResponse(ctx *gin.Context, code int, status string, data any){
	ctx.JSON(code, HTTPResponse{
		Meta: MetaData{
			Code: code,
			Status: status,
		},
		Data: data,
	})
}