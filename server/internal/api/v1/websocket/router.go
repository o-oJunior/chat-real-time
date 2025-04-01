package websocket

import (
	"server/internal/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func Router(rt *gin.Engine) {
	middleware := middleware.NewMiddlewareToken()
	rt.GET("/ws", middleware.ValidateCookie, websocketHandler)
}
