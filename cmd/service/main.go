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
	"sync"
)

func main() {
	ctx := context.Background()
	cfg := core.GetEnviroments()

	// Initialize database
	wg := &sync.WaitGroup{}
	callbacks := []func(wait *sync.WaitGroup){
		func(wait *sync.WaitGroup) {
			defer wg.Done()
			mongodriven.ConnectMongoDB(ctx, cfg.MongoUrl, cfg.MongoDatabase, cfg.AppName)
		},
		func(wait *sync.WaitGroup) {
			defer wg.Done()
			redisdriven.ConnectRedisDB(ctx, cfg.RedisUrl)
		},
		// Add more connections here
	}

	wg.Add(len(callbacks))
	for _, cb := range callbacks {
		go cb(wg)
	}
	wg.Wait()

	defer mongodriven.DisconnectMongoDB(ctx)
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
