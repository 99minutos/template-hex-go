package ports

import (
	"context"
	"service/internal/domain/entities"
)

// Crea tus repositorios con las funciones que necesites
// Define las entradas y las salidas de tus funciones
type IExampleRepository interface {
	CreateExample(ctx context.Context) (*entities.Example, error)
	GetExample(ctx context.Context, exampleId string) (*entities.Example, error)
}
