package example

import (
	"context"
	"example-service/internal/domain/entities"
	"example-service/internal/domain/ports"
	"example-service/internal/infrastructure/driven/tracer"
)

type ExampleService struct {
	exRep ports.IExampleRepository
}

func NewExampleService(
	exRep ports.IExampleRepository,
) *ExampleService {
	return &ExampleService{
		exRep: exRep,
	}
}

func (s *ExampleService) CreateExample(ctx context.Context) (*entities.Example, error) {
	ctx, span := tracer.GetTracer().Start(ctx, "Request/CreateExample")
	defer span.End()

	return s.exRep.CreateExample(ctx)
}

func (s *ExampleService) GetExample(ctx context.Context, exampleId string) (*entities.Example, error) {
	ctx, span := tracer.GetTracer().Start(ctx, "Request/GetExample")
	defer span.End()
	return s.exRep.GetExample(ctx, exampleId)
}
