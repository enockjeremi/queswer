package domain

import (
	"encoding/json"

	"github.com/lib/pq"
)

type Question struct {
	Model
	Title       string        `json:"title" binding:"required,min=6"`
	Description string        `json:"description" binding:"required"`
	Completed   bool          `json:"completed"`
	Answers     []Answer      `json:"answers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint          `json:"-"`
	User        User          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes       pq.Int64Array `gorm:"type:integer[]"`
}

func (q *Question) TableName() string {
	return "question"
}

type QuestionRepository interface {
	GetAll() ([]Question, error)
	Create(question *Question) error
	GetOne(question *Question, id string) error
	Update(question *Question) error
	Delete(question *Question, id string) error
}

type QuestionUsecase interface {
	GetAllQuestion() ([]Question, error)
	CreateQuestion(question *Question) error
	GetQuestion(question *Question, id string) error
	UpdateQuestion(question *Question) error
	RemoveQuestion(question *Question, id string) error
}

func (q *Question) MarshalJSON() ([]byte, error) {
	type CreatedByJSON struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	return json.Marshal(struct {
		ID          uint          `json:"id"`
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Completed   bool          `json:"completed"`
		CreatedBy   CreatedByJSON `json:"createdBy"`
		Likes       uint          `json:"likes"`
		Answers     []Answer      `json:"answers"`
	}{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		Completed:   q.Completed,
		Answers:     q.Answers,
		Likes:       uint(len(q.Likes)),
		CreatedBy: CreatedByJSON{
			ID:       q.User.ID,
			Username: q.User.Username,
		},
	})
}
