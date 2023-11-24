package ports

import (
	"context"
	"example-service/internal/domain"
)

// Crea tus servicios con las funciones que necesites
// un servicio hace referencia a una entidad de negocio que dentro de ella contendra la
// capacidad de invocar a los repositorios
type IExampleService interface {
	CreateExample(ctx context.Context) (*domain.Example, error)
	GetExample(ctx context.Context, exampleId string) (*domain.Example, error)
}
