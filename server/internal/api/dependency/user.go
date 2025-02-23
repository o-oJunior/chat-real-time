package dependency

import (
	"server/internal/api/repository"
	"server/internal/api/service"
	"server/internal/api/v1/handler"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeUser(database *mongo.Database) handler.UserHandler {
	userRepository := repository.NewUserRepository(database)
	inviteRepository := repository.NewInviteRepository(database)
	userService := service.NewUserService(userRepository, inviteRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}
