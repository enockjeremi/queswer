package main

import (
	"fmt"
	"log"

	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	r.Run(":1341")
}
