package repository

import (
	"context"
	"fmt"
	"server/internal/api/entity"
	"server/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(*entity.User) error
	FindUsername(*entity.User) (*entity.User, error)
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

var logger *config.Logger = config.NewLogger("repository")

func (repository *userRepository) InsertUser(user *entity.User) error {
	logger.Info("Inserindo o usuário no banco de dados...")
	collection := repository.database.Collection("users")
	user.CreateAt = time.Now().UnixMilli()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		logger.Error("Erro ao inserir o usuário: %v", err)
		return err
	}
	logger.Info("Usuário inserido com sucesso!")
	return nil
}

func (repository userRepository) FindUsername(user *entity.User) (*entity.User, error) {
	logger.Info("Buscando o usuário pelo username...")
	collection := repository.database.Collection("users")
	filterUserRegex := bson.M{"$regex": fmt.Sprintf("^%s$", user.Username), "$options": "i"}
	filter := bson.D{{Key: "username", Value: filterUserRegex}}
	var result entity.User
	if err := collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return &result, fmt.Errorf("usuário não foi encontrado")
	}
	logger.Info("Usuário foi encontrado, retornando...")
	return &result, nil
}
