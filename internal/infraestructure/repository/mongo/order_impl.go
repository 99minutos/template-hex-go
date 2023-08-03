package mongo

import (
	"context"
	"example-service/internal/application/ports"
	"example-service/internal/domain"
	"example-service/internal/infraestructure/driven/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type ExampleRepository struct {
	tracer    trace.Tracer
	tableName string
	database  *mongo.Database
}

func NewExampleRepository(tableName string, database *mongo.Database, tracer trace.Tracer) ports.IExampleRepository {
	return &ExampleRepository{
		database:  database,
		tracer:    tracer,
		tableName: tableName,
	}
}

// CreateExample
func (r *ExampleRepository) CreateExample(ctx context.Context) (*domain.Example, error) {
	ctx, span := r.tracer.Start(ctx, "Repository/Example/CreateExample")
	defer span.End()

	span.SetAttributes(
		attribute.String("cosas.random", "valor"),
		attribute.String("cosas.random", "valor"),
	)
	example := &domain.Example{
		FirstName: "John",
		LastName:  "Doe",
		SubExample: domain.SubExample{
			SubExampleId:   123,
			SubExampleName: "subExampleName",
		},
	}
	collection := r.database.Collection(r.tableName)
	result, err := collection.InsertOne(ctx, example)
	if err != nil {
		return nil, err
	}

	example.Id = result.InsertedID.(string)
	return example, nil
}

// GetExample
func (r *ExampleRepository) GetExample(ctx context.Context, exampleId string) (*domain.Example, error) {
	ctx, span := r.tracer.Start(ctx, "Repository/Example/GetExample")
	defer span.End()

	collection := r.database.Collection(r.tableName)
	query := bson.M{
		"_id": exampleId,
	}
	var example *domain.Example

	err := collection.FindOne(ctx, query).Decode(example)
	if err != nil {
		logger.Logger.Error("error getting example", zap.Error(err))
		return nil, err
	}

	return example, nil
}
