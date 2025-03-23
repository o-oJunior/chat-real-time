package repository

import (
	"context"
	"server/internal/api/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository interface {
	InsertNotification(*entity.Notification) error
}

type notificationRepository struct {
	database *mongo.Database
}

func NewNotificationRepository(database *mongo.Database) NotificationRepository {
	return &notificationRepository{database}
}

func (repository *notificationRepository) InsertNotification(notification *entity.Notification) error {
	collection := repository.database.Collection("notifications")
	_, err := collection.InsertOne(context.Background(), notification)
	if err != nil {
		return err
	}
	return nil
}
