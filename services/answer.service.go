package services

import (
	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
)

func FindAllAnswer(answer *[]models.Answer) (err error) {
	if err := config.DB.Find(&answer).Error; err != nil {
		return err
	}
	return nil
}

func CreateAnswer(answer *models.Answer) (err error) {
	if err := config.DB.Create(&answer).Error; err != nil {
		return err
	}
	return nil
}

func FindOneAnswer(answer *models.Answer, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(answer).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAnswer(answer *models.Answer, id string) (err error) {
	config.DB.Save(answer)
	return nil
}

func DeleteAnswer(answer *models.Answer, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(answer)
	return
}
