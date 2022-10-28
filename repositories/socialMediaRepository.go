package repositories

import (
	"MyGram/models"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	Create(socmed models.SocialMedia) (models.SocialMedia, error)
	FindAll() ([]models.SocialMedia, error)
	Save(socmed models.SocialMedia) (models.SocialMedia, error)
	FindByID(ID int) (models.SocialMedia, error)
	Delete(socmed models.SocialMedia) (models.SocialMedia, error)
}

type socmedRepository struct {
	db *gorm.DB
}

func NewSocmedRepository(db *gorm.DB) *socmedRepository {
	return &socmedRepository{
		db: db,
	}
}

func (r *socmedRepository) Create(socmed models.SocialMedia) (models.SocialMedia, error) {
	return socmed, r.db.Create(&socmed).Error
}

func (r *socmedRepository) FindAll() ([]models.SocialMedia, error) {
	socialMedias := []models.SocialMedia{}
	err := r.db.Find(&socialMedias).Error

	return socialMedias, err
}

func (r *socmedRepository) Save(socmed models.SocialMedia) (models.SocialMedia, error) {
	return socmed, r.db.Save(&socmed).Error
}

func (r *socmedRepository) FindByID(ID int) (models.SocialMedia, error) {
	socmed := models.SocialMedia{}
	err := r.db.Where("id = ?", ID).First(&socmed).Error
	return socmed, err
}

func (r *socmedRepository) Delete(socmed models.SocialMedia) (models.SocialMedia, error) {
	return socmed, r.db.Delete(&socmed).Error
}
