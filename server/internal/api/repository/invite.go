package repository

import (
	"context"
	"server/internal/api/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InviteRepository interface {
	InsertInvite(*entity.Invite) error
	FindInvitesByUsers(string, []string, string) ([]entity.Invite, error)
	UpdateStatusInvite(string, string) error
	DeleteInviteById(string) error
}

type inviteRepository struct {
	database *mongo.Database
}

func NewInviteRepository(database *mongo.Database) InviteRepository {
	return &inviteRepository{database}
}

func (repository *inviteRepository) InsertInvite(invite *entity.Invite) error {
	collection := repository.database.Collection("invites")
	_, err := collection.InsertOne(context.Background(), invite)
	if err != nil {
		return err
	}
	return nil
}

func (repository *inviteRepository) FindInvitesByUsers(userIdLogged string, userIds []string, searchField string) ([]entity.Invite, error) {
	collection := repository.database.Collection("invites")
	var filter bson.M
	if searchField == "userIdInvited" || searchField == "userIdInviter" {
		filter = bson.M{searchField: userIdLogged}
	} else {
		filter = bson.M{
			"$or": []bson.M{
				{"userIdInvited": userIdLogged, "userIdInviter": bson.M{"$in": userIds}},
				{"userIdInviter": userIdLogged, "userIdInvited": bson.M{"$in": userIds}},
			},
		}
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var invites []entity.Invite
	if err := cursor.All(context.Background(), &invites); err != nil {
		return nil, err
	}

	return invites, nil
}

func (repository *inviteRepository) UpdateStatusInvite(idInvite string, statusInvite string) error {
	collection := repository.database.Collection("invites")
	id, err := primitive.ObjectIDFromHex(idInvite)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"inviteStatus": statusInvite}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repository *inviteRepository) DeleteInviteById(idInvite string) error {
	collection := repository.database.Collection("invites")
	id, err := primitive.ObjectIDFromHex(idInvite)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
