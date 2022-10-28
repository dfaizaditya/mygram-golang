package services

import (
	"errors"
	"strconv"

	"MyGram/models"
	"MyGram/repositories"
)

type PhotoService interface {
	Create(userID int, createRequest models.Photo) (models.Photo, error)
	GetAll() ([]models.GetPhotoResponse, error)
	Update(paramID string, userID int, updateRequest models.Photo) (models.Photo, error)
	Delete(paramID string, userID int) (models.Photo, error)
}

type photoService struct {
	photoRepository repositories.PhotoRepository
	userRepository  repositories.UserRepository
}

func NewPhotoService(photoRepository repositories.PhotoRepository, userRepository repositories.UserRepository) *photoService {
	return &photoService{
		photoRepository: photoRepository,
		userRepository:  userRepository,
	}
}

func (s *photoService) Create(userID int, createRequest models.Photo) (models.Photo, error) {
	newPhoto := models.Photo{
		Title:    createRequest.Title,
		Caption:  createRequest.Caption,
		PhotoURL: createRequest.PhotoURL,
	}

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return newPhoto, err
	}

	return s.photoRepository.Create(user, newPhoto)
}

func (s *photoService) GetAll() ([]models.GetPhotoResponse, error) {
	response := []models.GetPhotoResponse{}
	photos, err := s.photoRepository.Find()
	if err != nil {
		return response, err
	}

	for _, photo := range photos {
		user, err := s.userRepository.FindByID(int(photo.UserID))
		if err != nil {
			return response, err
		}
		response = append(response, models.ParseToGetPhotoResponse(photo, user))
	}
	return response, nil
}

func (s *photoService) Update(paramID string, userID int, updateRequest models.Photo) (models.Photo, error) {
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		return models.Photo{}, err
	}

	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return models.Photo{}, err
	}

	if photo.UserID != uint(userID) {
		return models.Photo{}, errors.New("unauthorized")
	}

	photo.Title = updateRequest.Title
	photo.Caption = updateRequest.Caption
	photo.PhotoURL = updateRequest.PhotoURL

	return s.photoRepository.Save(photo)
}

func (s *photoService) Delete(paramID string, userID int) (models.Photo, error) {
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		return models.Photo{}, err
	}

	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return models.Photo{}, err
	}

	if photo.UserID != uint(userID) {
		return models.Photo{}, errors.New("unauthorized")
	}

	return s.photoRepository.Delete(photo)
}
