package handlers

import (
	"net/http"

	"MyGram/helpers"
	"MyGram/models"
	"MyGram/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoHandlers struct {
	photoService services.PhotoService
}

func NewPhotoHandlers(photoService services.PhotoService) *PhotoHandlers {
	return &PhotoHandlers{
		photoService: photoService,
	}
}

func (h *PhotoHandlers) UploadPhoto(c *gin.Context) {
	createPhotoRequest := models.Photo{}
	if err := c.ShouldBindJSON(&createPhotoRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	photo, err := h.photoService.Create(int(userId), createPhotoRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToCreatePhotoResponse(photo),
	})
}

func (h *PhotoHandlers) GetAllPhotos(c *gin.Context) {
	photos, err := h.photoService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   photos,
	})
}

func (h *PhotoHandlers) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	photoRequest := models.Photo{}
	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	result, err := h.photoService.Update(photoID, int(userId), photoRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToUpdatePhotoResponse(result),
	})
}

func (h *PhotoHandlers) DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	_, err := h.photoService.Delete(photoID, int(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseMessage{
		Status:  "success",
		Message: "Your photo has been successfully deleted",
	})
}
