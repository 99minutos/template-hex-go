package redis

import (
	"context"
	"example-service/internal/domain/core"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"net/url"
	"os"
	"strings"
)

var redisClient *redis.Client

func ConnectRedisDB(ctx context.Context, acx *core.AppContext) {
	acx.Infow("Redis is starting...")

	redisUrlParsed, err := url.Parse(os.Getenv("REDIS_URL"))
	password, _ := redisUrlParsed.User.Password()

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisUrlParsed.Host,
		Password: password,
	})

	if rClient == nil {
		acx.Errorw("unable to connect to redis instance", "error", err)
	}

	if err != nil {
		acx.Errorw("error creating redis connection pool", "error", err)
	}

	if err := redisotel.InstrumentTracing(rClient); err != nil {
		acx.Errorw("error creating telemetry for redis", "error", err)
	}

	redisClient = rClient
	acx.Infow("Redis is connected")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
func DisconnectRedisDB(ctx context.Context, acx *core.AppContext) {
	if redisClient == nil {
		acx.Fatalw("redis client is nil")
		return
	}
	err := redisClient.Close()
	if err != nil {
		acx.Infow("disconnected from redis failed", "error", err)
	} else {

		acx.Infow("disconnected from redis")
	}
}

// join array of strings with a separator : and return a string
func CachePath(s []string) string {
	return strings.Join(s, ":")
}
