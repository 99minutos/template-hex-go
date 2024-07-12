package main

import (
	"context"
	"example-service/internal/implementation/repository/mongo"
	services "example-service/internal/implementation/services/example"
	"example-service/internal/infraestructure/driven/cmux"
	"example-service/internal/infraestructure/driven/core"
	mongodriven "example-service/internal/infraestructure/driven/mongodb"
	redisdriven "example-service/internal/infraestructure/driven/redis"
	"example-service/internal/infraestructure/driver/rest"
)

func main() {
	ctx := context.Background()
	cfg := core.GetEnviroments()

	// Initialize database
	mongodriven.ConnectMongoDB(ctx, cfg.MongoUrl, cfg.MongoDatabase, cfg.AppName)
	defer mongodriven.DisconnectMongoDB(ctx)

	redisdriven.ConnectRedisDB(ctx, cfg.RedisUrl)
	defer redisdriven.DisconnectRedisDB(ctx)

	// Initialize repositories
	exampleRep := mongo.NewExampleRepository(mongodriven.GetDatabase())

	// Initialize services
	exampleSrv := services.NewExampleService(exampleRep)

	// Initialize the rest server
	restServer := rest.NewRestHandler(exampleSrv)
	restServer.InitializeRoutes()

	multiplexor := cmux.NewCmux(cfg.Port)
	go restServer.Start(multiplexor.GetHttpListener())
	multiplexor.Start()
}
