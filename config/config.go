package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Initialize() error {
	var err error

	//Initialize DB
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Error initializing SQLite database: %v", err)
	}
	return nil
}

func GetLogger() *Logger {
	//Initialize Logger
	logger = NewLogger()
	return logger
}

func GetSQLite() *gorm.DB {
	return db
}
