package models

import "encoding/json"

type Question struct {
	Model
	Title       string   `json:"title" binding:"required,min=6"`
	Description string   `json:"description" binding:"required"`
	Completed   bool     `json:"completed"`
	Answer      []Answer `json:"answers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint     `json:"-"`
	User        User     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (q *Question) TableName() string {
	return "question"
}

func (q *Question) MarshalJSON() ([]byte, error) {
	type CreatedByJSON struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	if q.Answer == nil {
		return json.Marshal(struct {
			ID          uint          `json:"id"`
			Title       string        `json:"title"`
			Description string        `json:"description"`
			Completed   bool          `json:"completed"`
			CreatedBy   CreatedByJSON `json:"createdBy"`
			Answer      []Answer      `json:"-"`
		}{
			ID:          q.ID,
			Title:       q.Title,
			Description: q.Description,
			Completed:   q.Completed,
			Answer:      q.Answer,
			CreatedBy: CreatedByJSON{
				ID:       q.User.ID,
				Username: q.User.Username,
			},
		})
	}

	return json.Marshal(struct {
		ID          uint          `json:"id"`
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Completed   bool          `json:"completed"`
		CreatedBy   CreatedByJSON `json:"createdBy"`
		Answer      []Answer      `json:"answers"`
	}{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		Completed:   q.Completed,
		Answer:      q.Answer,
		CreatedBy: CreatedByJSON{
			ID:       q.User.ID,
			Username: q.User.Username,
		},
	})
}
