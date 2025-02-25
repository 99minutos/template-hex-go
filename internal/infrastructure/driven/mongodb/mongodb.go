package mongodbrepo

import (
	"context"

	"service/internal/infrastructure/driven/dbg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type MongoRepository struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
}

func NewMongoConnection(ctx context.Context, mongoUrl string, database string, appName string) *MongoRepository {
	debugger := dbg.GetLogger()
	debugger.Infow("MongoDB is starting...")

	clientOptions := options.Client()
	clientOptions.SetMonitor(otelmongo.NewMonitor())
	clientOptions.ApplyURI(mongoUrl)
	clientOptions.SetAppName(appName)
	clientOptions.SetCompressors([]string{"zstd", "zlib"})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		debugger.Fatalw("Error creating client for MongoDB", "error", err)
	}

	databaseOptions := options.Database()

	debugger.Infow("MongoDB is connected")
	return &MongoRepository{
		mongoClient:   client,
		mongoDatabase: client.Database(database, databaseOptions),
	}
}

func (m *MongoRepository) GetDatabase() *mongo.Database {
	return m.mongoDatabase
}

func (m *MongoRepository) DisconnectMongoDB(ctx context.Context) {
	debugger := dbg.GetLogger()
	if m.mongoClient == nil {
		debugger.Fatalw("MongoDB client is nil")
		return
	}

	err := m.mongoClient.Disconnect(ctx)
	if err != nil {
		debugger.Fatalw("Error disconnecting from MongoDB", "error", err)
	}

	debugger.Infow("MongoDB is disconnected")
}
