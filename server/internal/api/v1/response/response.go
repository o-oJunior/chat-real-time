package response

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"statusCode"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendError(ctx *gin.Context, code int, message string) {
	response := response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, response)
}

func SendSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	response := response{
		code,
		message,
		data,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, response)
}
