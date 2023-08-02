package redis

import (
	"context"
	"example-service/internal/infraestructure/driven/logger"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"net/url"
	"os"
	"strings"
)

var redisClient *redis.Client

func ConnectRedisDB(ctx context.Context) {
	log := logger.Logger
	log.Info("Redis is starting...")

	redisUrlParsed, err := url.Parse(os.Getenv("REDIS_URL"))
	password, _ := redisUrlParsed.User.Password()

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisUrlParsed.Host,
		Password: password,
	})

	if rClient == nil {
		log.Error(fmt.Sprintf("unable to connect to redis instance: %v", err))
	}

	if err != nil {
		log.Error(fmt.Sprintf("error creating redis connection pool: %v", err))
	}

	if err := redisotel.InstrumentTracing(rClient); err != nil {
		log.Error(fmt.Sprintf("error creating telemetry for redis: %v", err))
	}

	redisClient = rClient
	log.Info("Redis is connected")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
func DisconnectRedisDB(ctx context.Context) {
	log := logger.Logger
	if redisClient == nil {
		log.Fatal("redis client is nil")
		return
	}
	err := redisClient.Close()
	if err != nil {
		log.Info("disconnected from redis failed")
	} else {

		log.Info("disconnected from redis")
	}
}

// join array of strings with a separator : and return a string
func CachePath(s []string) string {
	return strings.Join(s, ":")
}
