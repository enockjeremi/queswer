package services

import (
	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
)

func FindAllQuestion(question *[]models.Question) (err error) {
	if err = config.DB.Model(&question).Preload("Answer").Find(&question).Error; err != nil {
		return err
	}
	return nil
}

func CreateQuestion(question *models.Question) (err error) {
	if err = config.DB.Create(&question).Error; err != nil {
		return err
	}
	return nil
}

func GetOneQuestion(question *models.Question, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Preload("Answer").First(question).Error; err != nil {
		return err
	}
	return nil
}

func UpdateQuestion(question *models.Question, id string) (err error) {
	config.DB.Save(question)
	return nil
}

func DeleteQuestion(question *models.Question, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(question)
	return nil
}
