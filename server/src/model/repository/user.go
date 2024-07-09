package repository

import (
	"context"
	"server/src/config/logger"
	"server/src/model/dto"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(dto.UserDTO) error
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

func (ur *userRepository) CreateUser(user dto.UserDTO) error {
	logger.Info("[REPOSITORY] Inserindo o usuário no banco de dados...")
	collection := ur.database.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		logger.Error("[REPOSITORY] Erro ao inserir o usuário!", err, false)
		return err
	}
	logger.Info("[REPOSITORY] Usuário inserido com sucesso!")
	return nil
}
