package controller

import (
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
	})
}

func sendSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"data":       data,
	})
}
