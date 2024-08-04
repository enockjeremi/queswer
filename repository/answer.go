package repository

import (
	"github.com/enockjeremi/queswer/domain"
	"gorm.io/gorm"
)

type answerRepository struct {
	DB *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) domain.AnswerRepository {
	return &answerRepository{DB: db}
}

func (r *answerRepository) GetAll() ([]domain.Answer, error) {
	var answer []domain.Answer
	err := r.DB.Model(&answer).Preload("User").Find(&answer).Error
	return answer, err
}

func (r *answerRepository) Create(answer *domain.Answer) (err error) {
	if err = r.DB.Model(&answer).Create(&answer).Error; err != nil {
		return err
	}
	return nil
}

func (r *answerRepository) GetOne(answer *domain.Answer, id string) (err error) {
	if err = r.DB.Model(&answer).Where("id = ?", id).Preload("Answer.User").Preload("User").First(&answer).Error; err != nil {
		return err
	}
	return nil
}

func (r *answerRepository) Update(answer *domain.Answer) (err error) {
	r.DB.Model(&answer).Save(answer)
	return nil
}

func (r *answerRepository) Delete(answer *domain.Answer, id string) (err error) {
	r.DB.Model(&answer).Where("id = ?", id).Delete(answer)
	return nil
}

//GET QUESTION

func (r *answerRepository) GetOneQuestion(question *domain.Question, id string) (err error) {
	if err = r.DB.Model(&question).Where("id = ?", id).Preload("Answer.User").Preload("User").First(&question).Error; err != nil {
		return err
	}
	return nil
}
