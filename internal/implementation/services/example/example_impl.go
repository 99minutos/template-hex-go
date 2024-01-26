package example

import (
	"context"
	"example-service/internal/domain/core"
	"example-service/internal/domain/ports"
	"go.opentelemetry.io/otel/trace"
)

type ExampleService struct {
	acx       *core.AppContext
	tracer    trace.Tracer
	globalCtx context.Context
	exRep     ports.IExampleRepository
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
	acx *core.AppContext,
	exRep ports.IExampleRepository,
	tracer trace.Tracer,
) ports.IExampleService {
	return &ExampleService{
		acx:       acx,
		globalCtx: ctx,
		exRep:     exRep,
		tracer:    tracer,
	}
}
