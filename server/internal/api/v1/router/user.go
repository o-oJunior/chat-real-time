package router

import (
	"server/internal/api/dependency"
	"server/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouters(router *gin.RouterGroup, database *mongo.Database) {
	rt := router.Group("/user")
	handler := dependency.InitializeUser(database)
	middleware := middleware.NewMiddlewareToken()
	rt.GET("/search", handler.GetUsers)
	rt.POST("/create", handler.CreateUser)
	rt.POST("/authentication", handler.Authentication)
	rt.GET("/validate/authentication", middleware.ValidateCookie, handler.GetUserToken)
	rt.GET("/logout", handler.Logout)
	rt.GET("/contacts", middleware.ValidateCookie, handler.GetContacts)
}
