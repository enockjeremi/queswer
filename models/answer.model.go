package models

type Answer struct {
	Model
	Content    string `json:"content" binding:"required"`
	QuestionID uint   `json:"question"`
}

func (a *Answer) TableName() string {
	return "answer"
}
