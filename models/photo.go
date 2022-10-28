package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint
	User     *User
}

type GetPhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}

type CreatePhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ParseToCreatePhotoResponse(photo Photo) CreatePhotoResponse {
	return CreatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
}

func ParseToUpdatePhotoResponse(photo Photo) UpdatePhotoResponse {
	return UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
}

func ParseToGetPhotoResponse(p Photo, u User) GetPhotoResponse {
	return GetPhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		User: struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			Email:    u.Email,
			Username: u.Username,
		},
	}
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
