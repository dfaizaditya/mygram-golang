package services

import (
	"MyGram/models"
	"MyGram/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userRegisterRequest models.User) (models.User, error)
	Login(userLoginRequest models.User) (models.User, error)
	Update(id int, userUpdateRequest models.User) (models.User, error)
	Delete(id int) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) Register(userRegisterRequest models.User) (models.User, error) {
	newUser := models.User{
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: userRegisterRequest.Password,
		Age:      userRegisterRequest.Age,
	}

	return s.repository.Create(newUser)
}

func (s *userService) Login(userLoginRequest models.User) (models.User, error) {
	userFound, err := s.repository.FindByEmail(userLoginRequest.Email)
	if err != nil {
		return userFound, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))

	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

func (s *userService) Update(id int, userUpdateRequest models.User) (models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Email = userUpdateRequest.Email
	user.Username = userUpdateRequest.Username

	return s.repository.Save(user)
}

func (s *userService) Delete(id int) (models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	return s.repository.Delete(user)
}
