package middleware

import (
	"os"
	"server/internal/config"
)

var logger *config.Logger = config.NewLogger("middleware")

func NewMiddlewareToken() Token {
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")
	return &token{[]byte(PRIVATE_KEY)}
}
