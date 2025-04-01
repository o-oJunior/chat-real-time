package websocket

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	clients map[string]*websocket.Conn
	mutex   sync.Mutex
}

var wsManager = &WebSocketManager{clients: make(map[string]*websocket.Conn)}

func AddClient(userID string, conn *websocket.Conn) {
	logger.Info("Adicionando usuário '%s' na conexão!", userID)
	wsManager.mutex.Lock()
	defer wsManager.mutex.Unlock()
	wsManager.clients[userID] = conn
}

func RemoveClient(userID string) {
	logger.Info("Removedo usuário '%s' da conexão!", userID)
	wsManager.mutex.Lock()
	defer wsManager.mutex.Unlock()
	if conn, ok := wsManager.clients[userID]; ok {
		conn.Close()
		delete(wsManager.clients, userID)
	}
}

func SendMessageToUser(userID string, message string, typeMessage string) error {
	logger.Info("Enviando mensagem via WebSocket...")
	wsManager.mutex.Lock()
	defer wsManager.mutex.Unlock()
	data := map[string]interface{}{
		"message": message,
		"type":    typeMessage,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("Erro ao converter para []byte: %v", err)
		return fmt.Errorf("error internal server")
	}
	if conn, ok := wsManager.clients[userID]; ok {
		return conn.WriteMessage(websocket.TextMessage, dataBytes)
	}

	return fmt.Errorf("user %s offline", userID)
}
