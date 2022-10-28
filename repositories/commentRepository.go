package repositories

import (
	"MyGram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.Comment) (models.Comment, error)
	FindAll() ([]models.Comment, error)
	FindByID(ID int) (models.Comment, error)
	Save(comment models.Comment) (models.Comment, error)
	Delete(comment models.Comment) (models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment models.Comment) (models.Comment, error) {
	return comment, r.db.Create(&comment).Error
}

func (r *commentRepository) FindAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindByID(ID int) (models.Comment, error) {
	var comment models.Comment
	err := r.db.Where("id = ?", ID).First(&comment).Error
	return comment, err
}

func (r *commentRepository) Save(comment models.Comment) (models.Comment, error) {
	return comment, r.db.Save(&comment).Error
}

func (r *commentRepository) Delete(comment models.Comment) (models.Comment, error) {
	return comment, r.db.Delete(&comment).Error
}
