package models

type Question struct {
	Model
	Title       string `json:"title" binding:"required,min=6"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}

func (q *Question) TableName() string {
	return "question"
}
