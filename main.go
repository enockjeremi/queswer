package main

import (
	"fmt"
	"log"

	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file")
	}
	config.DB, err = gorm.Open(postgres.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	config.DB.AutoMigrate(
		&models.Question{},
		&models.Answer{},
		&models.User{},
	)

	r := routes.SetupRoute()

	r.Run(":1341")
}
