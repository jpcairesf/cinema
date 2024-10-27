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

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
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
