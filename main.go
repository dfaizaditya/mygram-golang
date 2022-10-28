package main

import (
	"MyGram/database"
	"MyGram/handlers"
	"MyGram/repositories"
	"MyGram/router"
	"MyGram/services"
)

func main() {
	database.StartDB()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	photoRepository := repositories.NewPhotoRepository(database.GetDB())
	photoService := services.NewPhotoService(photoRepository, userRepository)
	photoHandler := handlers.NewPhotoHandlers(photoService)

	commentRepository := repositories.NewCommentRepository(database.GetDB())
	commentService := services.NewCommentService(commentRepository, userRepository, photoRepository)
	commentHandler := handlers.NewCommentHandlers(commentService)

	socmedRepository := repositories.NewSocmedRepository(database.GetDB())
	socmedService := services.NewSocmedService(socmedRepository, userRepository)
	socmedHandler := handlers.NewSocialMediaHandlers(socmedService)

	r := router.StartApp(userHandler, photoHandler, commentHandler, socmedHandler)
	r.Run(":8080")
}
