package models

import "encoding/json"

type Answer struct {
	Model
	Content    string `json:"content" binding:"required"`
	QuestionID uint   `json:"questionId" binding:"required"`
	UserID     uint   `json:"-"`
	User       User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a *Answer) TableName() string {
	return "answer"
}

func (a *Answer) MarshalJSON() ([]byte, error) {
	type CreatedByJSON struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	return json.Marshal(struct {
		ID        uint          `json:"id"`
		Content   string        `json:"content"`
		CreatedBy CreatedByJSON `json:"createdBy"`
	}{
		ID:      a.ID,
		Content: a.Content,
		CreatedBy: CreatedByJSON{
			ID:       a.User.ID,
			Username: a.User.Username,
		},
	})
}
