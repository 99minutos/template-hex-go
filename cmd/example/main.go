package main

import (
	"context"
	"example-service/internal/application/repository/mongo"
	services "example-service/internal/application/services/example"
	"example-service/internal/config"
	"example-service/internal/infraestructure/adapters/driven/cmux"
	"example-service/internal/infraestructure/adapters/driven/envs"
	"example-service/internal/infraestructure/adapters/driven/logger"
	mongodriven "example-service/internal/infraestructure/adapters/driven/mongodb"
	redisdriven "example-service/internal/infraestructure/adapters/driven/redis"
	"example-service/internal/infraestructure/adapters/driven/tracer"
	"example-service/internal/infraestructure/adapters/driver/rest"
)

func main() {
	var cfg config.AppConfig

	ctx := context.Background()
	ctx = envs.WithEnvs(ctx, &cfg)
	tracerInstance, err := tracer.Setup(ctx, &cfg)
	if err != nil {
		logger.Logger.Info("Failed to initialize tracer.")
	}

	// Initialize database
	mongodriven.ConnectMongoDB(ctx, &cfg)
	defer mongodriven.DisconnectMongoDB(ctx)

	redisdriven.ConnectRedisDB(ctx)
	defer redisdriven.DisconnectRedisDB(ctx)

	// Initialize repositories
	exRep := mongo.NewExampleRepository("examples", mongodriven.GetDatabase(), tracerInstance)

	// Initialize services
	service := services.NewExampleService(ctx, exRep, tracerInstance)

	// Initialize the rest server
	restServer := rest.NewRestHandler(service)
	restServer.InitializeRoutes(cfg)

	// grpcServer := grpc.NewSotGrpcServer(service)

	multiplexor := cmux.NewCmux(cfg.Port)
	go restServer.Start(multiplexor.GetHttpListener())
	// si tu servicio necesita ofrecer GRPC descomenta esta linea
	// y busca realizar un proceso similar al que se realizo en src/internal/infraestructure/driver/grpc/server.go
	// de lo contrario borra el folder src/internal/infraestructure/driver/grpc y estas lineas
	// go grpcServer.Start(multiplexor.GetGrpcListener())
	multiplexor.Start()
}
