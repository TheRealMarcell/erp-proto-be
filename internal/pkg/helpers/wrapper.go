package helpers

import (
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Result common output
type Result struct {
	Data     interface{}
	MetaData interface{}
	Error    error
}

type response struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type Meta struct {
	Method        string    `json:"method"`
	Url           string    `json:"url"`
	Code          string    `json:"code"`
	ContentLength int64     `json:"content_length"`
	Date          time.Time `json:"date"`
	Ip            string    `json:"ip"`
}

type MetaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*errors.ErrorString)
	if ok {
		return errString.Code()
	}

	// default http status code
	return http.StatusInternalServerError
}

// RespSuccess sends a success response
func RespSuccess(c *gin.Context, log log.Logger, data interface{}, message string) {
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.ClientIP() // Use client IP if X-Forwarded-For is not present
	}

	meta := Meta{
		Date:          time.Now(),
		Url:           c.FullPath(),
		Method:        c.Request.Method,
		Code:          fmt.Sprintf("%v", http.StatusOK),
		ContentLength: int64(c.Request.ContentLength),
		Ip:            ip,
	}

	log.Info(c, "audit-log", fmt.Sprintf("%+v", meta))

	c.JSON(http.StatusOK, response{
		Meta: MetaResponse{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

// RespError sends an error response
func RespError(c *gin.Context, log log.Logger, err error) {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.FullPath(),
		Method:        c.Request.Method,
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		ContentLength: int64(c.Request.ContentLength),
	}

	log.Info(c, "audit-log", fmt.Sprintf("%+v", meta))

	// Return the JSON response directly in the method
	c.JSON(getErrorStatusCode(err), response{
		Meta: MetaResponse{
			Code:    getErrorStatusCode(err),
			Message: err.Error(),
		},
		Data: nil,
	})
}
// RespErrorWithData sends an error response with data
func RespErrorWithData(c *gin.Context, log log.Logger, data interface{}, err error) {
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.ClientIP() // Use client IP if X-Forwarded-For is not present
	}

	meta := Meta{
		Date:          time.Now(),
		Url:           c.FullPath(),
		Method:        c.Request.Method,
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		Ip:            ip,
		ContentLength: int64(c.Request.ContentLength),
	}

	log.Info(c, "audit-log", fmt.Sprintf("%+v", meta))

	c.JSON(getErrorStatusCode(err), response{
		Meta: MetaResponse{
			Code:    getErrorStatusCode(err),
			Message: err.Error(),
		},
		Data: data,
	})
}

// RespCustomError sends a custom error response with a custom error code
func RespCustomError(c *gin.Context, log log.Logger, err error) {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.FullPath(),
		Method:        c.Request.Method,
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		ContentLength: int64(c.Request.ContentLength),
	}

	log.Info(c, "audit-log", fmt.Sprintf("%+v", meta))

	fmt.Println("Safe")

	errString, ok := err.(*errors.ErrorString)
	metaErrorCode := http.StatusInternalServerError
	if ok {
		if errString.HttpCode() != 0 {
			metaErrorCode = errString.HttpCode()
		} else {
			metaErrorCode = errString.Code()
		}
	}
	fmt.Println("safe")
	c.JSON(metaErrorCode, response{
		Meta: MetaResponse{
			Code:    getErrorStatusCode(err),
			Message: err.Error(),
		},
		Data: nil,
	})
}
