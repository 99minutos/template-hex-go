package mongo

import (
	"context"
	"example-service/internal/domain/entities"
	"example-service/internal/domain/ports"
	"example-service/internal/infrastructure/driven/core"
	"example-service/internal/infrastructure/driven/tracer"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/attribute"
)

type ExampleRepository struct {
	tableName         string
	database          *mongo.Database
	exampleCollection *mongo.Collection
}

func NewExampleRepository(database *mongo.Database) ports.IExampleRepository {
	tableName := "example"
	collection := database.Collection(tableName)
	return &ExampleRepository{
		database:          database,
		tableName:         tableName,
		exampleCollection: collection,
	}
}

// CreateExample
func (r *ExampleRepository) CreateExample(ctx context.Context) (*entities.Example, error) {
	ctx, span := tracer.GetTracer().Start(ctx, "Repository/Example/CreateExample")
	defer span.End()

	span.SetAttributes(
		attribute.String("cosas.random1", "valor"),
		attribute.String("cosas.random2", "valor"),
	)
	example := &entities.Example{
		FirstName: "John",
		LastName:  "Doe",
		SubExample: entities.SubExample{
			SubExampleId:   123,
			SubExampleName: "subExampleName",
		},
	}
	collection := r.database.Collection(r.tableName)
	result, err := collection.InsertOne(ctx, example)
	if err != nil {
		return nil, err
	}

	example.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return example, nil
}

// GetExample
func (r *ExampleRepository) GetExample(ctx context.Context, exampleId string) (*entities.Example, error) {
	ctx, span := tracer.GetTracer().Start(ctx, "Repository/Example/GetExample")
	defer span.End()

	log := core.GetDefaultLogger()

	collection := r.database.Collection(r.tableName)
	objectId, err := primitive.ObjectIDFromHex(exampleId)
	if err != nil {
		log.Errorw("error parsing objectId", "error", err)

	}
	query := bson.M{
		"_id": objectId,
	}
	var example *entities.Example
	err = collection.FindOne(ctx, query).Decode(&example)
	if err != nil {
		log.Errorw("error getting example", "error", err)

	}
	return example, err
}
