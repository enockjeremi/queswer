package models

import (
	"encoding/json"
)

type User struct {
	Model
	Username  string  `json:"username" binding:"required,gt=6"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,gt=6"`
	ProfileID uint    `json:"-"`
	Profile   Profile `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) TableName() string {
	return "user"
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        uint    `json:"id"`
		Username  string  `json:"username"`
		Email     string  `json:"email"`
		ProfileID uint    `json:"-"`
		Profile   Profile `json:"profile"`
	}{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		ProfileID: u.ProfileID,
		Profile:   u.Profile,
	})
}
