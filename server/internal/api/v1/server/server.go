package server

import (
	"fmt"
	"io"
	"os"
	"server/internal/api/dependency"
	"server/internal/api/v1/router"
	"server/internal/config"
	"server/internal/config/mongodb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApiV1() {
	logger := config.NewLogger("server")
	port := fmt.Sprintf(":%s", os.Getenv("ENV_PORT_SERVER"))
	logger.Info("Sucesso ao iniciar o servidor na porta %s", port)
	database := mongodb.Connect()
	defer mongodb.Disconnect(database)
	gin.DefaultWriter = io.Discard
	rt := gin.Default()
	rt.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	userHandler := dependency.InitializeUser(database)
	v1 := rt.Group("/api/v1")
	userGroup := v1.Group("/user")
	router.UserRouters(userGroup, userHandler)
	rt.Run(port)
}
