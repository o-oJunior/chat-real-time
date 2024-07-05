package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/src/logger"
	"server/src/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("[ENV] Erro ao carregar as variaveis de ambiente!", err, true)
	}
	port := fmt.Sprintf(":%s", os.Getenv("ENV_PORT"))
	logger.Info("[SERVER] Sucesso ao iniciar o servidor na porta %s", port)
	router := router.Generate()
	log.Fatal(http.ListenAndServe(port, router))
}
