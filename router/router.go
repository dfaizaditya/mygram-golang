package router

import (
	"MyGram/handlers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(handler *handlers.UserHandler,
	photoHandler *handlers.PhotoHandlers,
	commentHandler *handlers.CommentHandlers,
	socmedHandler *handlers.SocialMediaHandlers,
) *gin.Engine {

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handler.UserRegisterHandler)
		userRouter.POST("/login", handler.UserLoginHandler)
		userRouter.PUT("/", middlewares.Authentication(), handler.UserUpdateHandler)
		userRouter.DELETE("/", middlewares.Authentication(), handler.DeleteUserHandler)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoHandler.UploadPhoto)
		photoRouter.GET("/", photoHandler.GetAllPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), photoHandler.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.GET("/", commentHandler.GetAllComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), commentHandler.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), commentHandler.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", socmedHandler.CreateSocmed)
		socialMediaRouter.GET("/", socmedHandler.GetAllSocmeds)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), socmedHandler.UpdateSocmed)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), socmedHandler.DeleteSocmed)
	}

	return r
}
