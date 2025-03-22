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
	inviteService := service.NewInviteService(inviteRepository, contactRepository)
	return handler.NewInviteHandler(inviteService)
}
