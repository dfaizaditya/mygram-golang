package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint
	User    *User
	PhotoID uint `json:"photo_id,string" form:"photo_id"`
	Photo   *Photo
	Message string `json:"message" form:"message" valid:"required~Comment is required"`
}

type GetCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	Photo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	}
}

type CreateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ParseToCreateCommentResponse(comment Comment) CreateCommentResponse {
	return CreateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
}

func ParseToUpdateCommentResponse(comment Comment) UpdateCommentResponse {
	return UpdateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}
}

func ParseToGetCommentResponse(comment Comment, user User, photo Photo) GetCommentResponse {
	return GetCommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User: struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		Photo: struct {
			ID       uint   `json:"id"`
			Title    string `json:"title"`
			Caption  string `json:"caption"`
			PhotoURL string `json:"photo_url"`
			UserID   uint   `json:"user_id"`
		}{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.UserID,
		},
	}
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
