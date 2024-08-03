package main

import (
	"server/internal/api/v1/server"
	"server/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	logger := config.NewLogger("main")
	if err := godotenv.Load("../../../.env"); err != nil {
		logger.Error("Erro ao carregar as variaveis de ambiente: %v", err)
		panic(err)
	}
	server.InitApiV1()
}
