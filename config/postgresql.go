package config

import (
	"os"

	"github.com/Lzrb0x/study-go-api.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSql() (*gorm.DB, error) {
	logger := GetLogger("postgresql")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to auto migrate schemas: %v", err)
		return nil, err
	}

	return db, nil
}
