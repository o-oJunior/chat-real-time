package repository

import (
	"context"
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/v1/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	InsertUser(*entity.User) error
	GetUserByID(primitive.ObjectID) (*entity.User, error)
	GetUsersWithFilter(bson.M, *middleware.Pagination) (*[]entity.User, int, error)
	FindUsername(string) (*entity.User, error)
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

func (repository *userRepository) InsertUser(user *entity.User) error {
	collection := repository.database.Collection("users")
	user.CreatedAtMilliseconds = time.Now().UnixMilli()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (repository *userRepository) GetUserByID(id primitive.ObjectID) (*entity.User, error) {
	collection := repository.database.Collection("users")
	filter := bson.M{"_id": id}
	var user *entity.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) FindUsername(username string) (*entity.User, error) {
	collection := repository.database.Collection("users")
	filterUserRegex := bson.M{"$regex": fmt.Sprintf("^%s$", username), "$options": "i"}
	filter := bson.D{{Key: "username", Value: filterUserRegex}}
	var result entity.User
	if err := collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return &result, fmt.Errorf("usuário não foi encontrado")
	}
	return &result, nil
}

func (repository *userRepository) GetUsersWithFilter(filter bson.M, pagination *middleware.Pagination) (*[]entity.User, int, error) {
	collection := repository.database.Collection("users")
	options := repository.buildPaginationOptions(pagination)
	users, err := repository.findUsers(collection, filter, options)
	if err != nil {
		return nil, 0, err
	}
	totalItems, err := repository.countItemsDocuments(collection, filter)
	if err != nil {
		return nil, 0, err
	}
	return &users, totalItems, nil
}

func (repository *userRepository) buildPaginationOptions(pagination *middleware.Pagination) *options.FindOptions {
	return options.Find().
		SetLimit(int64(pagination.Limit)).
		SetSkip(int64(pagination.Offset)).
		SetSort(bson.D{{Key: "username", Value: 1}}).
		SetCollation(&options.Collation{
			Locale:   "en",
			Strength: 1,
		})
}

func (repository *userRepository) findUsers(collection *mongo.Collection, filter bson.M, options *options.FindOptions) ([]entity.User, error) {
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	return repository.decodeUsers(cursor)
}

func (repository *userRepository) decodeUsers(cursor *mongo.Cursor) ([]entity.User, error) {
	var users []entity.User
	for cursor.Next(context.Background()) {
		var data entity.User
		if err := cursor.Decode(&data); err != nil {
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
		return nil, err
	}
	return users, nil
}

func (repository *userRepository) countItemsDocuments(collection *mongo.Collection, filter bson.M) (int, error) {
	totalItems, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(totalItems), nil
}
