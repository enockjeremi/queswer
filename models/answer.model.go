package models

type Answer struct {
	Model
	Content string `json:"content" binding:"required"`
}

func (a *Answer) TableName() string {
	return "answer"
}
