package router

import (
	"server/internal/api/injector"
	"server/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ContactRouters(router *gin.RouterGroup, database *mongo.Database) {
	rt := router.Group("/contact")
	handler := injector.InitializeContact(database)
	middleware := middleware.NewMiddlewareToken()
	rt.PUT("/update/:status", middleware.ValidateCookie, handler.UpdateStatusContact)
}
