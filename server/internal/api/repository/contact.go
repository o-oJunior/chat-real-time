package repository

import (
	"context"
	"server/internal/api/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepository interface {
	InsertContact(*entity.Contact) error
	GetContactsByUser(string) ([]entity.Contact, error)
}

type contactRepository struct {
	database *mongo.Database
}

func NewContactRepository(database *mongo.Database) ContactRepository {
	return &contactRepository{database}
}

func (repository contactRepository) InsertContact(contact *entity.Contact) error {
	collection := repository.database.Collection("contacts")
	_, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		return err
	}
	return nil
}

func (repository *contactRepository) GetContactsByUser(userIdLogged string) ([]entity.Contact, error) {
	collection := repository.database.Collection("contacts")
	filter := bson.M{
		"$or": []bson.M{
			{"userIdInvited": userIdLogged},
			{"userIdInviter": userIdLogged},
		},
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var contacts []entity.Contact
	if err := cursor.All(context.Background(), &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}
