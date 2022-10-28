package handlers

import (
	"net/http"
	"strconv"

	"MyGram/helpers"
	"MyGram/models"
	"MyGram/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandlers struct {
	socmedService services.SocmedService
}

func NewSocialMediaHandlers(socmedService services.SocmedService) *SocialMediaHandlers {
	return &SocialMediaHandlers{
		socmedService: socmedService,
	}
}

func (h *SocialMediaHandlers) CreateSocmed(c *gin.Context) {
	socmedRequest := models.SocialMedia{}
	if err := c.ShouldBindJSON(&socmedRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	socialMedia, err := h.socmedService.Add(int(userId), socmedRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToCreateSocialMediaResponse(socialMedia),
	})
}

func (h *SocialMediaHandlers) GetAllSocmeds(c *gin.Context) {
	socmeds, err := h.socmedService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: gin.H{
			"social_medias": socmeds,
		},
	})
}

func (h *SocialMediaHandlers) UpdateSocmed(c *gin.Context) {
	socmedID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	socmedUpdateRequest := models.SocialMedia{}
	if err := c.ShouldBindJSON(&socmedUpdateRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	socmed, err := h.socmedService.UpdateSocmed(socmedID, int(userId), socmedUpdateRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToUpdateSocialMediaResponse(socmed),
	})
}

func (h *SocialMediaHandlers) DeleteSocmed(c *gin.Context) {
	socmedID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	_, err = h.socmedService.DeleteSocmed(socmedID, int(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseMessage{
		Message: "Your social media has been successfully deleted",
	})
}
