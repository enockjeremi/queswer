package models

import "encoding/json"

type Question struct {
	Model
	Title       string   `json:"title" binding:"required,min=6"`
	Description string   `json:"description" binding:"required"`
	Completed   bool     `json:"completed"`
	Answer      []Answer `json:"answers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (q *Question) TableName() string {
	return "question"
}

func (q *Question) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          uint     `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Completed   bool     `json:"completed"`
		Answer      []Answer `json:"answers"`
	}{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		Completed:   q.Completed,
		Answer:      q.Answer,
	})
}
