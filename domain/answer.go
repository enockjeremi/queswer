package domain

import (
	"encoding/json"

	"github.com/lib/pq"
)

type Answer struct {
	Model
	Answer     string        `json:"answer" binding:"required"`
	QuestionID uint          `json:"questionId" binding:"required"`
	UserID     uint          `json:"-"`
	User       User          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes      pq.Int64Array `gorm:"type:integer[]"`
}

func (a *Answer) TableName() string {
	return "answer"
}

type AnswerRepository interface {
	GetAll() ([]Answer, error)
	Create(answer *Answer) error
	GetOne(answer *Answer, id string) error
	Update(answer *Answer) error
	Delete(answer *Answer, id string) error

	GetOneQuestion(Question *Question, id string) error
}

type AnswerUsecase interface {
	GetAllAnswer() ([]Answer, error)
	CreateAnswer(answer *Answer) error
	GetAnswer(answer *Answer, id string) error
	UpdateAnswer(answer *Answer) error
	RemoveAnswer(answer *Answer, id string) error

	GetQuestion(question *Question, id string) error
}

func (a *Answer) MarshalJSON() ([]byte, error) {
	type CreatedByJSON struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	return json.Marshal(struct {
		ID        uint          `json:"id"`
		Answer    string        `json:"answer"`
		Likes     uint          `json:"likes"`
		CreatedBy CreatedByJSON `json:"createdBy"`
	}{
		ID:     a.ID,
		Answer: a.Answer,
		Likes:  uint(len(a.Likes)),
		CreatedBy: CreatedByJSON{
			ID:       a.User.ID,
			Username: a.User.Username,
		},
	})
}
