package main

import (
	"context"
	"example-service/internal/domain/core"
	"example-service/internal/implementation/repository/mongo"
	services "example-service/internal/implementation/services/example"
	"example-service/internal/infraestructure/driven/cmux"
	"example-service/internal/infraestructure/driven/core/envs"
	"example-service/internal/infraestructure/driven/core/logger"
	mongodriven "example-service/internal/infraestructure/driven/mongodb"
	redisdriven "example-service/internal/infraestructure/driven/redis"
	"example-service/internal/infraestructure/driven/tracer"
	"example-service/internal/infraestructure/driver/rest"
)

func main() {
	ctx, cfg := loadEnvsWithContext(context.Background())

	logs := logger.NewLogger()
	acx := core.NewAppContext(cfg, logs)

	tracerInstance, err := tracer.Setup(ctx, acx)

	if err != nil {
		acx.Infow("Failed to initialize tracer.")
	}

	// Initialize database
	mongodriven.ConnectMongoDB(ctx, acx)
	defer mongodriven.DisconnectMongoDB(ctx, acx)

	redisdriven.ConnectRedisDB(ctx, acx)
	defer redisdriven.DisconnectRedisDB(ctx, acx)

	// Initialize repositories
	exRep := mongo.NewExampleRepository(acx, "examples", mongodriven.GetDatabase(), tracerInstance)

	// Initialize services
	service := services.NewExampleService(ctx, acx, exRep, tracerInstance)

	// Initialize the rest server
	restServer := rest.NewRestHandler(acx, service)
	restServer.InitializeRoutes(cfg)

	// grpcServer := grpc.NewSotGrpcServer(service)

	multiplexor := cmux.NewCmux(acx)
	go restServer.Start(multiplexor.GetHttpListener())
	// si tu servicio necesita ofrecer GRPC descomenta esta linea
	// y busca realizar un proceso similar al que se realizo en src/internal/infraestructure/driver/grpc/server.go
	// de lo contrario borra el folder src/internal/infraestructure/driver/grpc y estas lineas
	// go grpcServer.Start(multiplexor.GetGrpcListener())
	multiplexor.Start()
}

func loadEnvsWithContext(ctx context.Context) (context.Context, *core.AppConfig) {
	var cfg core.AppConfig
	ctx = envs.WithEnvs(ctx, &cfg)
	return ctx, &cfg
}
