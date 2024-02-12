package config

import (
	"os"

	"github.com/Ribeiro/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath string = "./db/main.db"

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger()

	//Check if DB file already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Creating DB file: %v", dbPath)
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	//Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to open database: %v", err)
		return nil, err
	}

	//Migration Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Automigration error: %v", err)
		return nil, err
	}

	//Return DB
	return db, nil
}
