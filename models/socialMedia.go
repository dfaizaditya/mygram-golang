package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
	User           *User
}

type GetSocmedResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
}

type CreateSocialMediaResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

func ParseToCreateSocialMediaResponse(socmed SocialMedia) CreateSocialMediaResponse {
	return CreateSocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		CreatedAt:      socmed.CreatedAt,
	}
}

func ParseToUpdateSocialMediaResponse(socmed SocialMedia) UpdateSocialMediaResponse {
	return UpdateSocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		UpdatedAt:      socmed.UpdatedAt,
	}
}

func ParseToGetSocmedResponse(socmed SocialMedia, user User) GetSocmedResponse {
	return GetSocmedResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		CreatedAt:      socmed.CreatedAt,
		UpdatedAt:      socmed.UpdatedAt,
		User: struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
		},
	}
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
