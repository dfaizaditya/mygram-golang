package models

import (
	"MyGram/helpers"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email       string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Password    string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length 6 characters"`
	Age         uint   `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}

type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllPhotosUserResponse struct {
	Username string
	Email    string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if len(u.Password) < 6 {
		err = errors.New("password must have at least 6 characters length")
		return err
	}

	if u.Age < 8 {
		err = errors.New("minimum age to register is 8")
		return err
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
