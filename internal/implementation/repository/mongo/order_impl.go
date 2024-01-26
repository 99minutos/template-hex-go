package mongo

import (
	"context"
	"example-service/internal/domain/core"
	"example-service/internal/domain/entities"
	"example-service/internal/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ExampleRepository struct {
	acx       *core.AppContext
	tracer    trace.Tracer
	tableName string
	database  *mongo.Database
}

func NewExampleRepository(acx *core.AppContext, tableName string, database *mongo.Database, tracer trace.Tracer) ports.IExampleRepository {
	return &ExampleRepository{
		acx:       acx,
		database:  database,
		tracer:    tracer,
		tableName: tableName,
	}
}

// CreateExample
func (r *ExampleRepository) CreateExample(ctx context.Context) (*entities.Example, error) {
	ctx, span := r.tracer.Start(ctx, "Repository/Example/CreateExample")
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
	ctx, span := r.tracer.Start(ctx, "Repository/Example/GetExample")
	defer span.End()

	collection := r.database.Collection(r.tableName)
	objectId, err := primitive.ObjectIDFromHex(exampleId)
	if err != nil {
		r.acx.Errorw("error parsing objectId", "error", err)

	}
	query := bson.M{
		"_id": objectId,
	}
	var example *entities.Example
	err = collection.FindOne(ctx, query).Decode(&example)
	if err != nil {
		r.acx.Errorw("error getting example", "error", err)

	}
	return example, err
}
