package repositories

import (
	"MyGram/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(id int) (models.User, error)
	Save(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id int) (models.User, error) {
	user := models.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Delete(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
