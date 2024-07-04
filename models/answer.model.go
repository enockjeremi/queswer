package models

import "encoding/json"

type Answer struct {
	Model
	Content    string `json:"content" binding:"required"`
	QuestionID uint   `json:"questionId" binding:"required"`
}

func (a *Answer) TableName() string {
	return "answer"
}

func (a *Answer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID      uint   `json:"id"`
		Content string `json:"content"`
	}{
		ID:      a.ID,
		Content: a.Content,
	})
}
