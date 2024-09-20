package mongodbrepo

import (
	"context"
	"service/internal/infrastructure/driven/core"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectMongoDB(ctx context.Context, mongoUrl string, database string, appName string) {
	log := core.GetDefaultLogger()

	clientOptions := options.Client()
	clientOptions.SetMonitor(otelmongo.NewMonitor())
	clientOptions.ApplyURI(mongoUrl)
	clientOptions.SetAppName(appName)
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
	mongoDatabase = mongoClient.Database(database, databaseOptions)
	log.Infow("MongoDB is connected")
}

func GetDatabase() *mongo.Database {
	return mongoDatabase
}

func DisconnectMongoDB(ctx context.Context) {
	log := core.GetDefaultLogger()
	if mongoClient == nil {
		log.Fatalw("mongodb client is nil")
		return
	}

	err := mongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatalw("mongodb disconnect error", "error", err)
	}

	log.Infow("mongodb is disconnected")
}
