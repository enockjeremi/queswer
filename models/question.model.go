package models

import "encoding/json"

type Question struct {
	Model
	Title       string   `json:"title" binding:"required,min=6"`
	Description string   `json:"description" binding:"required"`
	Completed   bool     `json:"completed"`
	Answer      []Answer `json:"answers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint     `json:"userid"`
	User        User     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (q *Question) TableName() string {
	return "question"
}

func (q *Question) MarshalJSON() ([]byte, error) {
	type UserJSON struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	return json.Marshal(struct {
		ID          uint     `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Completed   bool     `json:"completed"`
		User        UserJSON `json:"user"`
		Answer      []Answer `json:"answers"`
		UserID      uint     `json:"-"`
	}{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		Completed:   q.Completed,
		Answer:      q.Answer,
		UserID:      q.UserID,
		User: UserJSON{
			ID:       q.User.ID,
			Username: q.User.Username,
		},
	})
}
