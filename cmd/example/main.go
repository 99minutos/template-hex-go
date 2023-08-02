package main

import (
	"context"
	"example-service/internal/adapters/repository/mongo"
	services "example-service/internal/adapters/services/example"
	"example-service/internal/config"
	"example-service/internal/infraestructure/driven/cmux"
	"example-service/internal/infraestructure/driven/envs"
	"example-service/internal/infraestructure/driven/logger"
	mongodriven "example-service/internal/infraestructure/driven/mongodb"
	redisdriven "example-service/internal/infraestructure/driven/redis"
	"example-service/internal/infraestructure/driven/tracer"
	"example-service/internal/infraestructure/driver/rest"
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
	service := services.NewExampleService(ctx, exRep)

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
