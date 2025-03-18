package dependency

import (
	"server/internal/api/repository"
	"server/internal/api/service"
	"server/internal/api/v1/handler"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeContact(database *mongo.Database) handler.ContactHandler {
	contactRepository := repository.NewContactRepository(database)
	contactService := service.NewContactService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)
	return contactHandler
}
