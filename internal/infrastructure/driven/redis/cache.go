package redis

import (
	"context"
	"example-service/internal/infrastructure/driven/core"
	"net/url"
	"strings"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func ConnectRedisDB(ctx context.Context, redisUrl string) {
	log := core.GetDefaultLogger()
	log.Infow("Redis is starting...")

	redisUrlParsed, err := url.Parse(redisUrl)
	password, _ := redisUrlParsed.User.Password()

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisUrlParsed.Host,
		Password: password,
	})

	if rClient == nil {
		log.Errorw("unable to connect to redis instance", "error", err)
	}

	if err != nil {
		log.Errorw("error creating redis connection pool", "error", err)
	}

	if err := redisotel.InstrumentTracing(rClient); err != nil {
		log.Errorw("error creating telemetry for redis", "error", err)
	}

	redisClient = rClient
	log.Infow("Redis is connected")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
func DisconnectRedisDB(ctx context.Context) {
	log := core.GetDefaultLogger()
	if redisClient == nil {
		log.Fatalw("redis client is nil")
		return
	}
	err := redisClient.Close()
	if err != nil {
		log.Infow("disconnected from redis failed", "error", err)
	} else {

		log.Infow("disconnected from redis")
	}
}

// join array of strings with a separator : and return a string
func CachePath(s []string) string {
	return strings.Join(s, ":")
}
