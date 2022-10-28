package services

import (
	"errors"

	"MyGram/models"
	"MyGram/repositories"
)

type CommentService interface {
	Create(userID int, commentRequest models.Comment) (models.Comment, error)
	GetAll() ([]models.GetCommentResponse, error)
	Update(commentID int, userID int, commentRequest models.Comment) (models.Comment, error)
	Delete(commentID int, userID int) (models.Comment, error)
}

type commentService struct {
	commentRepository repositories.CommentRepository
	userRepository    repositories.UserRepository
	photoRepository   repositories.PhotoRepository
}

func NewCommentService(commentRepository repositories.CommentRepository, userRepository repositories.UserRepository, photoRepository repositories.PhotoRepository) *commentService {
	return &commentService{
		commentRepository: commentRepository,
		userRepository:    userRepository,
		photoRepository:   photoRepository,
	}
}

func (s *commentService) Create(userID int, commentRequest models.Comment) (models.Comment, error) {
	comment := models.Comment{
		UserID:  uint(userID),
		PhotoID: commentRequest.PhotoID,
		Message: commentRequest.Message,
	}

	return s.commentRepository.Create(comment)
}

func (s *commentService) GetAll() ([]models.GetCommentResponse, error) {
	var Response []models.GetCommentResponse
	comments, err := s.commentRepository.FindAll()
	if err != nil {
		return Response, err
	}

	for _, comment := range comments {
		user, err := s.userRepository.FindByID(int(comment.UserID))

		if err != nil {
			return Response, err
		}

		photo, err2 := s.photoRepository.FindByID(int(comment.PhotoID))

		if err2 != nil {
			return Response, err
		}

		c := models.ParseToGetCommentResponse(comment, user, photo)
		Response = append(Response, c)
	}

	return Response, nil
}

func (s *commentService) Update(commentID int, userID int, commentRequest models.Comment) (models.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != uint(userID) {
		return comment, errors.New("unauthorized")
	}

	comment.Message = commentRequest.Message
	return s.commentRepository.Save(comment)
}

func (s *commentService) Delete(commentID int, userID int) (models.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != uint(userID) {
		return comment, errors.New("unauthorized")
	}

	return s.commentRepository.Delete(comment)
}
