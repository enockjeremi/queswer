package main

import (
	"fmt"
	"log"

	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/delivery/http"
	"github.com/enockjeremi/queswer/domain"
	"github.com/enockjeremi/queswer/repository"
	"github.com/enockjeremi/queswer/usecase"
	"github.com/enockjeremi/queswer/utils"
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
	config.DB, err = gorm.Open(postgres.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	if err := config.DB.AutoMigrate(
		&domain.Question{},
		&domain.Answer{},
		&domain.User{},
		&domain.Profile{},
	); err != nil {
		panic("failed to auto-migrate schema")
	}

	r := gin.Default()
	questionRepo := repository.NewQuestionRepository(config.DB)
	questionUsecase := usecase.NewQuestionUsecase(questionRepo)
	http.NewQuestionHandler(r, questionUsecase)

	answerRepo := repository.NewAnswerRepository(config.DB)
	answerUsecase := usecase.NewAnswerUsecase(answerRepo)
	http.NewAnswerHandler(r, answerUsecase)

	authRepo := repository.NewAuthRepository(config.DB)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	http.NewAuthHandler(r, authUsecase)

	utils.NewErrorFormatter().LengFormatter("en")
	r.Run(":1341")
}
