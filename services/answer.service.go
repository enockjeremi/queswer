package services

import (
	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
)

func CreateAnswer(answer *models.Answer) (err error) {
	if err := config.DB.Create(&answer).Error; err != nil {
		return err
	}
	return nil
}
