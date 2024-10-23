package repository

import (
	"context"
	"fmt"
	"server/internal/api/entity"
	"server/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	InsertUser(*entity.User) error
	FindUsername(*entity.User) (*entity.User, error)
	GetUsersAndTotalExceptID(primitive.ObjectID, string, *options.FindOptions) (*[]entity.User, int64, error)
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
	user.CreatedAtMilliseconds = time.Now().UnixMilli()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		logger.Error("Erro ao inserir o usuário: %v", err)
		return err
	}
	logger.Info("Usuário inserido com sucesso!")
	return nil
}

func (repository *userRepository) FindUsername(user *entity.User) (*entity.User, error) {
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

func (repository *userRepository) GetUsersAndTotalExceptID(id primitive.ObjectID, username string, options *options.FindOptions) (*[]entity.User, int64, error) {
	logger.Info("Buscando usuários...")
	collection := repository.database.Collection("users")
	filter := bson.M{
		"_id": bson.M{"$ne": id},
		"username": bson.M{
			"$regex": primitive.Regex{Pattern: username, Options: "i"},
		},
	}
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		logger.Error("Erro ao buscar os usuários: %v", err)
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	var users []entity.User
	for cursor.Next(context.Background()) {
		var data entity.User
		if err := cursor.Decode(&data); err != nil {
			logger.Error("Erro ao decodificar usuário: %v", err)
			continue
		}
		user := entity.User{
			ID:          data.ID,
			Username:    data.Username,
			Description: data.Description,
			CreatedAt:   time.UnixMilli(data.CreatedAtMilliseconds).UTC().Format(time.RFC3339),
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		logger.Error("Erro ao iterar sobre o cursor: %v", err)
		return nil, 0, err
	}
	totalItems, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		logger.Error("Erro ao buscar quantidade total de usuários: %v", err)
		return nil, 0, err
	}
	logger.Info("Retornando %d usuários...", len(users))
	return &users, totalItems, nil
}
