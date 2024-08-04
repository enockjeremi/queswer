package repository

import (
	"github.com/enockjeremi/queswer/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepository{DB: db}
}

func (r *authRepository) RegisterUserRepo(user *domain.User) (err error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) VerifyCredentialsRepo(user *domain.User) (err error) {
	err = r.DB.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		err = r.DB.Where("email = ?", user.Email).First(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *authRepository) GetUserRepo(user *domain.User, id interface{}) (err error) {
	if err = r.DB.Where("id = ?", id).Preload("Profile").First(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) VerifyUsernameRepo(user *domain.User, userInput string) (err error) {
	err = r.DB.Where("username = ?", userInput).Find(&user).First(&user).Error
	if err != nil {
		return err
	}
	return nil

}

func (r *authRepository) CreateProfileRepo(profile *domain.Profile) (err error) {
	if err := r.DB.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) UpdateProfileRepo(profile *domain.Profile) (err error) {
	r.DB.Save(profile)
	return nil
}

func (r *authRepository) ChangePasswordRepo(user *domain.User, newPassword string) (err error) {
	r.DB.Model(&user).Update("Password", newPassword)
	return nil
}

func (r *authRepository) GetProfileRepo(profile *domain.Profile, id interface{}) (err error) {
	if err = r.DB.Where("id = ?", id).First(profile).Error; err != nil {
		return err
	}
	return nil
}
