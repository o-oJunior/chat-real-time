package repository

import (
	"context"
	"server/internal/api/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepository interface {
	InsertContact(*entity.Contact) error
	GetContactsByUser(primitive.ObjectID) ([]entity.Contact, error)
	FindContactByUsers(primitive.ObjectID, []primitive.ObjectID) ([]entity.Contact, error)
	UpdateStatusContact(primitive.ObjectID, string, int64) error
	DeleteContactById(primitive.ObjectID) error
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

func (repository *contactRepository) GetContactsByUser(userIdLogged primitive.ObjectID) ([]entity.Contact, error) {
	collection := repository.database.Collection("contacts")
	filter := bson.M{
		"$or": []bson.M{
			{"userIdTarget": userIdLogged},
			{"userIdActor": userIdLogged},
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

func (repository *contactRepository) FindContactByUsers(userIdLogged primitive.ObjectID, userIds []primitive.ObjectID) ([]entity.Contact, error) {
	collection := repository.database.Collection("contacts")
	filter := bson.M{
		"$or": []bson.M{
			{"userIdTarget": userIdLogged, "userIdActor": bson.M{"$in": userIds}},
			{"userIdActor": userIdLogged, "userIdTarget": bson.M{"$in": userIds}},
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

func (repository *contactRepository) UpdateStatusContact(id primitive.ObjectID, statusContact string, timestamp int64) error {
	collection := repository.database.Collection("contacts")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": statusContact, "updatedAt": timestamp}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repository *contactRepository) DeleteContactById(id primitive.ObjectID) error {
	collection := repository.database.Collection("contacts")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
