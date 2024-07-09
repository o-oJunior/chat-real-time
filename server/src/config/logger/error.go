package logger

import (
	"fmt"
	"time"
)

func Error(message string, err error, isCritical bool) {
	const formatDate string = "02/01/2006 | 15:04:05"
	date := time.Now().Format(formatDate)
	if err != nil && isCritical {
		fmt.Printf("[ %s ] -  \033[1;31m%s\033[0m: %s: %s\n", date, "Error", message, err)
		panic(err)
	} else if err != nil {
		fmt.Printf("[ %s ] -  \033[1;31m%s\033[0m: %s: %s\n", date, "Error", message, err)
	}
}
