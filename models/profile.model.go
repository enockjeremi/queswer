package models

import "encoding/json"

type Profile struct {
	Model
	FirstName string `json:"firstName,omitempty" binding:"omitempty,gt=0"`
	LastName  string `json:"lastName,omitempty" binding:"omitempty,gt=0"`
	Role      string `json:"role" gorm:"default:customer"`
	UserID    uint   `json:"userId"`
}

func (c *Profile) TableName() string {
	return "profile"
}

func (c Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{
		FirstName: c.FirstName,
		LastName:  c.LastName,
	})
}
