package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", logger.Flags()),
		info:   log.New(writer, "INFO: ", logger.Flags()),
		warn:   log.New(writer, "WARN: ", logger.Flags()),
		error:  log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}
