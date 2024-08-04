package domain

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
