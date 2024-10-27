package config

import (
	"io"
	"log"
	"os"
)

var logger *Logger

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.error.Printf(format, v...)
}

func newLogger() *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", log.LstdFlags),
		info:   log.New(writer, "INFO: ", log.LstdFlags),
		warn:   log.New(writer, "WARN: ", log.LstdFlags),
		error:  log.New(writer, "ERROR: ", log.LstdFlags),
		writer: writer,
	}
}
