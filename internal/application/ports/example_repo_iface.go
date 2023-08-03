package ports

import (
	"context"
	"example-service/internal/domain"
)

// Crea tus repositorios con las funciones que necesites
// Define las entradas y las salidas de tus funciones
type IExampleRepository interface {
	CreateExample(ctx context.Context) (*domain.Example, error)
	GetExample(ctx context.Context, exampleId string) (*domain.Example, error)
}
