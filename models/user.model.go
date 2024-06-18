package models

type User struct {
	Model
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) TableName() string {
	return "user"
}
