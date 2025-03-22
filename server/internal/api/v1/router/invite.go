package router

import (
	"server/internal/api/injector"
	"server/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InviteRouters(router *gin.RouterGroup, database *mongo.Database) {
	rt := router.Group("/invite")
	handler := injector.InitializeInvite(database)
	middleware := middleware.NewMiddlewareToken()
	rt.POST("/send", middleware.ValidateCookie, handler.InsertInvite)
	rt.PUT("/update/:status", middleware.ValidateCookie, handler.UpdateStatusInvite)
}
