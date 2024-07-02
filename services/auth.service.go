package services

import (
	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
)

func CreateUser(user *models.User) (err error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func VerifyCredentials(user *models.User) (err error) {
	err = config.DB.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		err = config.DB.Where("email = ?", user.Email).First(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUser(user *models.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func VerifyUsername(user *models.User, userInput string) (err error) {
	err = config.DB.Where("username = ?", userInput).Find(&user).First(&user).Error
	if err != nil {
		return err
	}
	return nil

}
