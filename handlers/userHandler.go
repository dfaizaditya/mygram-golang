package handlers

import (
	"net/http"

	"MyGram/helpers"
	"MyGram/models"
	"MyGram/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserRegisterHandler(c *gin.Context) {
	userRequest := models.User{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	newUser, err := h.service.Register(userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: models.UserRegisterResponse{
			ID:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
			Age:      newUser.Age,
		},
	})
}

func (h *UserHandler) UserLoginHandler(c *gin.Context) {
	loginRequest := models.User{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.Login(loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "unauthenticated",
			Message: err.Error(),
		})
		return
	}

	signedToken := helpers.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "unauthenticated",
			Message: err.Error(),
		})
		return
	}

	c.SetCookie("token", signedToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: models.LoginResponse{
			Token: signedToken,
		},
	})
}

func (h *UserHandler) UserUpdateHandler(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	updatedUser, err := h.service.Update(int(userId), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: models.UserUpdateResponse{
			ID:        updatedUser.ID,
			Username:  updatedUser.Username,
			Email:     updatedUser.Email,
			Age:       updatedUser.Age,
			UpdatedAt: *updatedUser.UpdatedAt,
		},
	})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	_, err := h.service.Delete(int(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.SetCookie("token", "", 0, "", "", false, true)

	c.JSON(http.StatusOK, helpers.ResponseMessage{
		Status:  "success",
		Message: "Your account has been successfully deleted",
	})
}
