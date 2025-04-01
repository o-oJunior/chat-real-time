package server

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"server/internal/api/v1/router"
	"server/internal/api/v1/websocket"
	"server/internal/config"
	"server/internal/config/mongodb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router   *gin.Engine
	database *mongo.Database
	logger   *config.Logger
}

func NewServer() *Server {
	logger := config.NewLogger("server")
	database := mongodb.Connect()
	gin.DefaultWriter = io.Discard
	rt := gin.Default()
	rt.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           0,
	}))

	websocket.Router(rt)

	v1 := rt.Group("/api/v1")
	router.UserRouters(v1, database)
	router.InviteRouters(v1, database)
	router.ContactRouters(v1, database)

	return &Server{
		router:   rt,
		database: database,
		logger:   logger,
	}
}

func (server *Server) Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("ENV_PORT_SERVER"))
	server.logger.Info("Sucesso ao iniciar o servidor na porta %s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		if err := server.router.Run(port); err != nil {
			errChan <- fmt.Errorf("erro crítico no servidor: %v", err)
		}
	}()

	select {
	case <-quit:
		server.logger.Warn("Servidor perdeu a conexão! Iniciando procedimento para encerrar corretamente.")
	case err := <-errChan:
		server.logger.Error("Falha detectada: %v", err)
	}

	server.Shutdown()
	return nil
}

func (server *Server) Shutdown() {
	server.logger.Info("Desligando o servidor...")
	mongodb.Disconnect(server.database)
	server.logger.Info("Banco de dados desconectado! Servidor encerrado.")
}
