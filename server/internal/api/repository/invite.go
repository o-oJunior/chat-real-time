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
	FindInvitesByUsers(string, []string) ([]entity.Invite, error)
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
	logger.Info("Inserindo convite no banco de dados...")
	collection := repository.database.Collection("invites")
	_, err := collection.InsertOne(context.Background(), invite)
	if err != nil {
		logger.Error("Erro ao inserir o convite: %v", err)
		return err
	}
	logger.Info("Convite inserido com sucesso!")
	return nil
}

func (repository *inviteRepository) FindInvitesByUsers(userIdLogged string, userIds []string) ([]entity.Invite, error) {
	logger.Info("Buscando convites no banco de dados para múltiplos usuários...")
	collection := repository.database.Collection("invites")
	filter := bson.M{
		"$or": []bson.M{
			{"userIdInvited": userIdLogged, "userIdInviter": bson.M{"$in": userIds}},
			{"userIdInviter": userIdLogged, "userIdInvited": bson.M{"$in": userIds}},
		},
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Erro ao buscar convites: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var invites []entity.Invite
	if err := cursor.All(context.Background(), &invites); err != nil {
		logger.Error("Erro ao decodificar convites: %v", err)
		return nil, err
	}

	logger.Info("Convites obtidos com sucesso!")
	return invites, nil
}

func (repository *inviteRepository) UpdateStatusInvite(idInvite string, statusInvite string) error {
	logger.Info("Atualizando o status do convite no banco de dados...")
	collection := repository.database.Collection("invites")
	id, err := primitive.ObjectIDFromHex(idInvite)
	if err != nil {
		logger.Error("Erro ao converter ID: %v", err)
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"inviteStatus": statusInvite}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Erro ao atualizar o convite: %v", err)
		return err
	}
	logger.Info("Convite atualizado com sucesso!")
	return nil
}

func (repository *inviteRepository) DeleteInviteById(idInvite string) error {
	logger.Info("Excluindo o convite do banco de dados...")
	collection := repository.database.Collection("invites")
	id, err := primitive.ObjectIDFromHex(idInvite)
	if err != nil {
		logger.Error("Erro ao converter ID: %v", err)
		return err
	}
	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Erro ao excluir o convite: %v", err)
		return err
	}
	logger.Info("Convite excluido com sucesso!")
	return nil
}
