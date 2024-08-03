package mongodb

import (
	"context"
	"os"
	"server/internal/config"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logger *config.Logger = config.NewLogger("mongodb")

func Connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stringConnection := os.Getenv("MONGO_STRING_CONNECTION")
	MONGO_NAME_DATABASE := os.Getenv("MONGO_NAME_DATABASE")
	clientOptions := options.Client().
		ApplyURI(stringConnection).
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(15 * time.Minute).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(30 * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("(Connect) Erro ao conectar ao banco de dados: %v", err)
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error("(Ping) Erro ao conectar ao banco de dados: %v", err)
		panic(err)
	}
	logger.Info("Conectado ao bando de dados...")
	return client.Database(MONGO_NAME_DATABASE)
}

func Disconnect(database *mongo.Database) {
	if err := database.Client().Disconnect(context.TODO()); err != nil {
		logger.Error("Erro ao desconectar do banco de dados %v", err)
		panic(err)
	}
	logger.Info("Desconectado do banco de dados")
}
