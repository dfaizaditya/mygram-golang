package services

import (
	"errors"

	"MyGram/models"
	"MyGram/repositories"
)

type SocmedService interface {
	Add(userID int, socmed models.SocialMedia) (models.SocialMedia, error)
	GetAll() ([]models.GetSocmedResponse, error)
	UpdateSocmed(ID int, userID int, updateRequest models.SocialMedia) (models.SocialMedia, error)
	DeleteSocmed(ID int, userID int) (models.SocialMedia, error)
}

type socmedService struct {
	socmedRepository repositories.SocmedRepository
	userRepository   repositories.UserRepository
}

func NewSocmedService(socmedRepository repositories.SocmedRepository, userRepository repositories.UserRepository) *socmedService {
	return &socmedService{
		socmedRepository: socmedRepository,
		userRepository:   userRepository,
	}
}

func (s *socmedService) Add(userID int, socmedRequest models.SocialMedia) (models.SocialMedia, error) {
	socmed := models.SocialMedia{
		Name:           socmedRequest.Name,
		SocialMediaURL: socmedRequest.SocialMediaURL,
		UserID:         uint(userID),
	}

	return s.socmedRepository.Create(socmed)
}

func (s *socmedService) GetAll() ([]models.GetSocmedResponse, error) {
	response := []models.GetSocmedResponse{}
	socmeds, err := s.socmedRepository.FindAll()
	if err != nil {
		return response, err
	}

	for _, socmed := range socmeds {
		user, err := s.userRepository.FindByID(int(socmed.UserID))
		if err != nil {
			return response, err
		}

		response = append(response, models.ParseToGetSocmedResponse(socmed, user))
	}

	return response, nil
}

func (s *socmedService) UpdateSocmed(ID int, userID int, updateRequest models.SocialMedia) (models.SocialMedia, error) {
	socmed, err := s.socmedRepository.FindByID(ID)
	if err != nil {
		return socmed, err
	}

	if uint(userID) != socmed.UserID {
		return socmed, errors.New("Unauthorized")
	}

	socmed.Name = updateRequest.Name
	socmed.SocialMediaURL = updateRequest.SocialMediaURL

	return s.socmedRepository.Save(socmed)
}

func (s *socmedService) DeleteSocmed(ID int, userID int) (models.SocialMedia, error) {
	socmed, err := s.socmedRepository.FindByID(ID)
	if err != nil {
		return socmed, err
	}

	if uint(userID) != socmed.UserID {
		return socmed, errors.New("Unauthorized")
	}

	return s.socmedRepository.Delete(socmed)
}
