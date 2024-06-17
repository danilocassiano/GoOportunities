package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	//Initialize Log

	logger = NewLogger(p)
	return logger
}