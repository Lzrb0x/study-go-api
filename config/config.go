package config

import (
	"fmt"

	"gorm.io/gorm"
)


var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	
	db, err = InitPostgreSql()
	if err != nil {
		return fmt.Errorf("failed to initialize PostgreSQL: %v", err)
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}


func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}