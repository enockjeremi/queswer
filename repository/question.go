package repository

import (
	"github.com/enockjeremi/queswer/domain"
	"gorm.io/gorm"
)

type questionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domain.QuestionRepository {
	return &questionRepository{DB: db}
}

func (r *questionRepository) GetAll() ([]domain.Question, error) {
	var question []domain.Question
	err := r.DB.Model(&question).Preload("Answers.User").Preload("User").Find(&question).Error
	return question, err
}

func (r *questionRepository) Create(question *domain.Question) (err error) {
	if err = r.DB.Model(&question).Create(&question).Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) GetOne(question *domain.Question, id string) (err error) {
	if err = r.DB.Model(&question).Where("id = ?", id).Preload("Answers.User").Preload("User").First(&question).Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) Update(question *domain.Question) (err error) {
	r.DB.Model(&question).Save(question)
	return nil
}

func (r *questionRepository) Delete(question *domain.Question, id string) (err error) {
	r.DB.Model(&question).Where("id = ?", id).Delete(question)
	return nil
}
