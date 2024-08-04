package domain

import "encoding/json"

type User struct {
	Model
	Username  string  `json:"username" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,gt=6"`
	ProfileID uint    `json:"-"`
	Profile   Profile `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	Model
	FirstName string `json:"firstName,omitempty" binding:"omitempty,gt=0"`
	LastName  string `json:"lastName,omitempty" binding:"omitempty,gt=0"`
	Role      string `json:"role" gorm:"default:customer"`
	UserID    uint   `json:"userId"`
}

func (u *User) TableName() string {
	return "user"
}

func (c *Profile) TableName() string {
	return "profile"
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

func (c Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{
		FirstName: c.FirstName,
		LastName:  c.LastName,
	})
}

type AuthRepository interface {
	RegisterUserRepo(user *User) error
	VerifyCredentialsRepo(user *User) error
	GetUserRepo(user *User, id interface{}) error
	VerifyUsernameRepo(user *User, userInput string) error

	CreateProfileRepo(profile *Profile) error
	UpdateProfileRepo(profile *Profile) error
	GetProfileRepo(profile *Profile, id interface{}) error
	ChangePasswordRepo(user *User, newPassword string) error
}

type AuthUsecase interface {
	RegisterUser(user *User) error
	VerifyCredentials(user *User) error
	GetUser(user *User, id interface{}) error
	VerifyUsername(user *User, userInput string) error

	CreateProfile(profile *Profile) error
	UpdateProfile(profile *Profile) error
	GetProfile(profile *Profile, id interface{}) error
	ChangePassword(user *User, newPassword string) error
}
