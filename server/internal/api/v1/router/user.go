package router

import (
	"server/internal/api/v1/handler"

	"github.com/gin-gonic/gin"
)

func UserRouters(v1 *gin.RouterGroup, handler handler.UserHandler) {
	v1.POST("/create", handler.CreateUser)
	v1.POST("/authentication", handler.Authentication)
}
