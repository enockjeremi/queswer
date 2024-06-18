package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.HOST,
		dbConfig.USER,
		dbConfig.PASSWORD,
		dbConfig.DBNAME,
		dbConfig.PORT,
	)
}
