package injector

import (
	"server/internal/api/repository"
	"server/internal/api/service"
	"server/internal/api/v1/handler"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeInvite(database *mongo.Database) handler.InviteHandler {
	inviteRepository := repository.NewInviteRepository(database)
	contactRepository := repository.NewContactRepository(database)
	notificationRepository := repository.NewNotificationRepository(database)
	inviteService := service.NewInviteService(inviteRepository, contactRepository, notificationRepository)
	return handler.NewInviteHandler(inviteService)
}
