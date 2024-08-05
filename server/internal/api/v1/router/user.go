package router

import (
	"server/internal/api/v1/handler"
	"server/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouters(v1 *gin.RouterGroup, handler handler.UserHandler) {
	newMiddlewareToken := middleware.NewMiddlewareToken()
	v1.POST("/create", handler.CreateUser)
	v1.POST("/authentication", handler.Authentication)
	v1.GET("/validate/authentication", newMiddlewareToken.ValidateCookie)
}
