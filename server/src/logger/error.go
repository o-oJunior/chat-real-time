package logger

import (
	"fmt"
	"log"
	"time"
)

func Error(message string, err error, isCritical bool) {
	const formatDate string = "02/01/2006 | 15:04:05"
	date := time.Now().Format(formatDate)
	fmt.Printf("[ %s ] -  \033[1;31m%s\033[0m: %s\n", date, "Error", message)
	if err != nil && isCritical {
		log.Fatal(err)
	}
}
