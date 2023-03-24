package config

import (
	"fmt"

	"github.com/LyoDekken/go-api/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	// sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	pgInfo := fmt.Sprintf("%s://%s:%s@%s:%s/%s", 
		config.DB, config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)

	db, err := gorm.Open(postgres.Open(pgInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
