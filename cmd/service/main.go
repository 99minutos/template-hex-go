package main

import (
	"context"

	services "service/internal/implementation/example"
	"service/internal/infrastructure/adapters/repository/mongo"
	"service/internal/infrastructure/driven/cmux"
	"service/internal/infrastructure/driven/core"
	mongodriven "service/internal/infrastructure/driven/mongodb"
	redisdriven "service/internal/infrastructure/driven/redis"
	"service/internal/infrastructure/driver/rest"
)

func main() {
	ctx := context.Background()
	cfg := core.GetEnviroments()

	// Initialize database
	mongoSocket := mongodriven.NewMongoConnection(ctx, cfg.MongoUrl, cfg.MongoDatabase, cfg.AppName)
	redisSocket := redisdriven.NewRedisConnection(ctx, cfg.RedisUrl, cfg.RedisBasePath)

	defer mongoSocket.DisconnectMongoDB(ctx)
	defer redisSocket.DisconnectRedisDB(ctx)

	// Initialize repositories
	exampleRep := mongo.NewExampleRepository(mongoSocket.GetDatabase())

	// Initialize services
	exampleSrv := services.NewExampleService(exampleRep)

	// Initialize the rest server
	restServer := rest.NewRestHandler(exampleSrv)
	restServer.InitializeRoutes()

	multiplexor := cmux.NewCmux(cfg.Port)
	go restServer.Start(multiplexor.GetHttpListener())
	multiplexor.Start()
}
