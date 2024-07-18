package repository

import (
	"context"
	"server/src/config"
	"server/src/model/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(*dto.User) error
	FindUsername(*dto.User) (dto.User, error)
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

var logger *config.Logger = config.NewLogger("repository")

func (repository *userRepository) CreateUser(user *dto.User) error {
	logger.Info("Inserindo o usuário no banco de dados...")
	collection := repository.database.Collection("users")
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		logger.Error("Erro ao inserir o usuário: %v", err)
		return err
	}
	logger.Info("Usuário inserido com sucesso!")
	return nil
}

func (repository userRepository) FindUsername(user *dto.User) (dto.User, error) {
	logger.Info("Buscando o usuário pelo username")
	collection := repository.database.Collection("users")
	filterUserRegex := bson.M{"$regex": user.Username, "$options": "i"}
	filter := bson.D{{Key: "username", Value: filterUserRegex}}
	var result dto.User
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
