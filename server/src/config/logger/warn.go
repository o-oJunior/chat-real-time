package logger

import (
	"fmt"
	"time"
)

func Warn(message string, args ...any) {
	const formatDate string = "02/01/2006 | 15:04:05"
	date := time.Now().Format(formatDate)
	fmt.Printf("[ %s ] - \033[1;33m%s\033[0m: %s\n", date, "Warn", fmt.Sprintf(message, args...))
}
