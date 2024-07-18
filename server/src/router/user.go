package router

import (
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.RouterGroup) {
	router.POST("/user/create", middleware.NewUserMiddleware().CreateUser)
}
