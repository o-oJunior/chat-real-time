package websocket

import (
	"net/http"
	"server/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var logger = config.NewLogger("websocket")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(ctx *gin.Context) {
	logger.Info("Conectando ao websocket...")
	userID := ctx.Query("userId")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId é obrigatório"})
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error("Erro ao conectar WebSocket: %v", err)
		return
	}
	defer conn.Close()
	AddClient(userID, conn)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			RemoveClient(userID)
			logger.Warn("Usuário desconectado: %v", userID)
			break
		}
	}
}
