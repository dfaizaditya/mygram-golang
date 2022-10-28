package repositories

import (
	"MyGram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(user models.User, photo models.Photo) (models.Photo, error)
	Find() ([]models.Photo, error)
	FindByID(ID int) (models.Photo, error)
	Save(photo models.Photo) (models.Photo, error)
	Delete(photo models.Photo) (models.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (r *photoRepository) Create(user models.User, photo models.Photo) (models.Photo, error) {
	err := r.db.Model(&user).Association("Photos").Append(&photo)
	return photo, err
}

func (r *photoRepository) Find() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Find(&photos).Error
	return photos, err
}

func (r *photoRepository) FindByID(ID int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Where("id = ?", ID).Find(&photo).Error
	return photo, err
}

func (r *photoRepository) Save(photo models.Photo) (models.Photo, error) {
	return photo, r.db.Save(&photo).Error
}

func (r *photoRepository) Delete(photo models.Photo) (models.Photo, error) {
	return photo, r.db.Delete(&photo).Error
}
