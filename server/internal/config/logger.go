package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
	prefix  string
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	return &Logger{
		debug:   log.New(writer, "", 0),
		info:    log.New(writer, "", 0),
		warning: log.New(writer, "", 0),
		err:     log.New(writer, "", 0),
		writer:  writer,
		prefix:  prefix,
	}
}

func (logger *Logger) setMessage(
	color string,
	logLevel string,
	format string,
	v ...interface{},
) string {
	const formatDate string = "02/01/2006 15:04:05"
	date := time.Now().Format(formatDate)
	stringColor := fmt.Sprintf("\033[1;%sm%s\033[0m", color, logLevel)
	rest := fmt.Sprintf(format, v...)
	message := fmt.Sprintf("%s - %s: [%s] %s", date, stringColor, logger.prefix, rest)
	return message
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	message := logger.setMessage("34", "Debug", format, v...)
	logger.debug.Printf(message)
}

func (logger *Logger) Info(format string, v ...interface{}) {
	message := logger.setMessage("32", "Info", format, v...)
	logger.info.Printf(message)
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	message := logger.setMessage("33", "Warn", format, v...)
	logger.warning.Printf(message)
}

func (logger *Logger) Error(format string, v ...interface{}) {
	message := logger.setMessage("31", "Error", format, v...)
	logger.err.Printf(message)
}
