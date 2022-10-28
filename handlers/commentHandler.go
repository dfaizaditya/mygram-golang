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

type CommentHandlers struct {
	commentService services.CommentService
}

func NewCommentHandlers(commentService services.CommentService) *CommentHandlers {
	return &CommentHandlers{
		commentService: commentService,
	}
}

func (h *CommentHandlers) CreateComment(c *gin.Context) {
	createCommentRequest := models.Comment{}
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	comment, err := h.commentService.Create(int(userId), createCommentRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToCreateCommentResponse(comment),
	})
}

func (h *CommentHandlers) GetAllComment(c *gin.Context) {
	comments, err := h.commentService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   comments,
	})
}

func (h *CommentHandlers) UpdateComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	commentRequest := models.Comment{}
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	comment, err := h.commentService.Update(commentID, int(userId), commentRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   models.ParseToUpdateCommentResponse(comment),
	})
}

func (h *CommentHandlers) DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	_, err = h.commentService.Delete(commentID, int(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.ResponseMessage{
		Status:  "success",
		Message: "Your comment has been sucessfully deleted",
	})
}
