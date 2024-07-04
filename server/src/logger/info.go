package logger

import (
	"fmt"
	"time"
)

func Info(message string, args ...any) {
	const formatDate string = "02/01/2006 | 15:04:05"
	date := time.Now().Format(formatDate)
	fmt.Printf("[ %s ] - \033[1;32m%s\033[0m: %s\n", date, "Info", fmt.Sprintf(message, args...))
}
