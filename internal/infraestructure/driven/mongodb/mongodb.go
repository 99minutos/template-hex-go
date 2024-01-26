package mongodbrepo

import (
	"context"
	"example-service/internal/domain/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectMongoDB(ctx context.Context, acx *core.AppContext) {
	acx.Infow("MongoDB is starting...")

	clientOptions := options.Client()
	clientOptions.SetMonitor(otelmongo.NewMonitor())
	clientOptions.ApplyURI(acx.Envs.MongoUrl)
	clientOptions.SetAppName(acx.Envs.AppName)
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
	mongoDatabase = mongoClient.Database(acx.Envs.MongoDatabase, databaseOptions)
	acx.Infow("MongoDB is connected")
}

func GetDatabase() *mongo.Database {
	return mongoDatabase
}

func DisconnectMongoDB(ctx context.Context, acx *core.AppContext) {
	if mongoClient == nil {
		acx.Fatalw("mongodb client is nil")
		return
	}

	err := mongoClient.Disconnect(ctx)
	if err != nil {
		acx.Fatalw("mongodb disconnect error", "error", err)
	}

	acx.Infow("mongodb is disconnected")
}
