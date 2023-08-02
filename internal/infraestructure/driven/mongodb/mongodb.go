package mongodbrepo

import (
	"context"
	"example-service/internal/config"
	"example-service/internal/infraestructure/driven/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectMongoDB(ctx context.Context, cfg *config.AppConfig) {
	log := logger.Logger
	log.Info("MongoDB is starting...")

	clientOptions := options.Client()
	clientOptions.SetMonitor(otelmongo.NewMonitor())
	clientOptions.ApplyURI(cfg.MongoUrl)
	clientOptions.SetAppName(cfg.AppName)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databaseOptions := options.Database()

	mongoClient = client
	mongoDatabase = mongoClient.Database(cfg.MongoDatabase, databaseOptions)
	log.Info("MongoDB is connected")
}

func GetDatabase() *mongo.Database {
	return mongoDatabase
}

func DisconnectMongoDB(ctx context.Context) {
	log := logger.Logger
	if mongoClient == nil {
		log.Fatal("MongoDB client is nil")
		return
	}

	err := mongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("MongoDB is disconnected")
}
