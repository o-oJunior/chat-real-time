package main

import (
	"fmt"
	"io"
	"os"
	"server/src/config/logger"
	"server/src/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("[ENV] Erro ao carregar as variaveis de ambiente!", err, true)
	}
	port := fmt.Sprintf(":%s", os.Getenv("ENV_PORT"))
	logger.Info("[SERVER] Sucesso ao iniciar o servidor na porta %s", port)
	gin.DefaultWriter = io.Discard
	rt := gin.Default()
	userGroup := rt.Group("/api/user")
	router.UserRouters(userGroup)
	rt.Run(port)
}
