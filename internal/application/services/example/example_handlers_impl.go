package example

import (
	"context"
	"example-service/internal/domain"
	"example-service/internal/infraestructure/adapters/driven/tracer"
	"fmt"
)

// cada funcion de tu servicio tiene la capacidad de acceder a los repositorios que tenga asociado la estructura
func (s *ExampleService) CreateExample(ctx context.Context) error {
	ctx, span := tracer.Tracer.Start(ctx, "Request/CreateExample")
	defer span.End()

	s.exRep.CreateExample(ctx) // -- este es un ejemplo de como el repositorio puede ser accedido desde el servicio
	return fmt.Errorf("not implemented")
}

// tambien observe como es importante agregar de esta forma el Tracer a cada funcion de tu servicio
// para que de esta forma se pueda tener un seguimiento de las peticiones que se realizan en Trace de google.
func (s *ExampleService) GetExample(ctx context.Context, exampleId string) (*domain.Example, error) {
	ctx, span := tracer.Tracer.Start(ctx, "Request/GetExample")
	defer span.End()

	s.exRep.GetExample(ctx, exampleId) // -- este es un ejemplo de como el repositorio puede ser accedido desde el servicio
	return nil, fmt.Errorf("not implemented")
}
