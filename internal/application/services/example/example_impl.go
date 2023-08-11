package example

import (
	"context"
	ports2 "example-service/internal/application/ports"
	"go.opentelemetry.io/otel/trace"
)

type ExampleService struct {
	tracer    trace.Tracer
	globalCtx context.Context
	exRep     ports2.IExampleRepository
}

// implementacion del servicio, notese como el puerto es devuelto en NewExampleService
// esto le indica al codigo que para poder funcionar necesita tener implementado todo aquello que la
// interface IExampleService tenga declarado en este ejemplo 2 funciones
// CreateExample y GetExample, pero a su vez pueden ver que este archivo no las contiene, si no que estan en otro
// archivo al mismo nivel que este "example_handlers_impl.go"
// esto es valido ya que ambos archivos estan con un "package example" y se concidera que pertenecen al mismo lugar
// Tambien obserca como obtienes un puerto de repositorio, esto se debe hacer para todos tu repositorios requridos para
// que este servicio funcione.
func NewExampleService(
	ctx context.Context,
	exRep ports2.IExampleRepository,
	tracer trace.Tracer,
) ports2.IExampleService {
	return &ExampleService{
		globalCtx: ctx,
		exRep:     exRep,
		tracer:    tracer,
	}
}
