package example

import (
	"context"
	"example-service/internal/domain/entities"
)

// cada funcion de tu servicio tiene la capacidad de acceder a los repositorios que tenga asociado la estructura
func (s *ExampleService) CreateExample(ctx context.Context) (*entities.Example, error) {
	ctx, span := s.tracer.Start(ctx, "Request/CreateExample")
	defer span.End()

	return s.exRep.CreateExample(ctx) // -- este es un ejemplo de como el repositorio puede ser accedido desde el servicio
}

// tambien observe como es importante agregar de esta forma el Tracer a cada funcion de tu servicio
// para que de esta forma se pueda tener un seguimiento de las peticiones que se realizan en Trace de google.
func (s *ExampleService) GetExample(ctx context.Context, exampleId string) (*entities.Example, error) {
	ctx, span := s.tracer.Start(ctx, "Request/GetExample")
	defer span.End()
	return s.exRep.GetExample(ctx, exampleId) // -- este es un ejemplo de como el repositorio puede ser accedido desde el servicio
}
