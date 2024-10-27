package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Initialize() error {
	return nil
}

func GetLogger() *Logger {
	logger = newLogger()
	return logger
}
