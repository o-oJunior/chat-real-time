package mongodb

import (
	"context"
	"os"
	"server/src/config/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stringConnection := os.Getenv("MONGO_STRING_CONNECTION")
	MONGO_NAME_DATABASE := os.Getenv("MONGO_NAME_DATABASE")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(stringConnection))
	if err != nil {
		logger.Error("[DATABASE (Connect)] Erro ao conectar ao banco de dados", err, true)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error("[DATABASE (Ping)] Erro ao conectar ao banco de dados", err, true)
	}
	logger.Info("[DATABASE] Conectado ao bando de dados...")
	return client.Database(MONGO_NAME_DATABASE)
}

func Disconnect(database *mongo.Database) {
	if err := database.Client().Disconnect(context.TODO()); err != nil {
		logger.Error("[DATABASE] Erro ao desconectar do banco de dados", err, true)
	}
	logger.Info("[DATABASE] Desconectado do banco de dados")
}
