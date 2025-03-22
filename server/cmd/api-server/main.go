package main

import (
	"os"
	"server/internal/api/v1/server"
	"server/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	logger := config.NewLogger("main")
	if err := godotenv.Load("../../../.env"); err != nil {
		logger.Error("Erro ao carregar as variaveis de ambiente: %v", err)
		os.Exit(1)
	}
	server := server.NewServer()
	server.Run()
}
