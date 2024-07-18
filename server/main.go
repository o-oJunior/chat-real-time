package main

import (
	"fmt"
	"io"
	"os"
	"server/src/config"
	"server/src/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := config.NewLogger("server")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Erro ao carregar as variaveis de ambiente: %v", err)
		panic(err)
	}
	port := fmt.Sprintf(":%s", os.Getenv("ENV_PORT"))
	logger.Info("Sucesso ao iniciar o servidor na porta %s", port)
	gin.DefaultWriter = io.Discard
	rt := gin.Default()
	rt.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir todos os domínios
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	v1 := rt.Group("/api/v1")
	router.UserRouters(v1)
	rt.Run(port)
}
