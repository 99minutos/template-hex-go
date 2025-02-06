package redis

import (
	"context"
	"net/url"
	"strings"

	"service/internal/infrastructure/driven/logs"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisConnection(ctx context.Context, redisUrl string) *RedisRepository {
	debugger := logs.GetLogger()
	debugger.Infow("Redis is starting...")

	redisUrlParsed, err := url.Parse(redisUrl)
	password, _ := redisUrlParsed.User.Password()

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisUrlParsed.Host,
		Password: password,
	})

	if rClient == nil {
		debugger.Errorw("unable to connect to redis instance", "err", err)
	}

	if err != nil {
		debugger.Errorw("error creating redis connection pool", "err", err)
	}

	if err := redisotel.InstrumentTracing(rClient); err != nil {
		debugger.Errorw("error creating telemetry for redis", "err", err)
	}

	debugger.Infow("Redis is connected")
	return &RedisRepository{
		redisClient: rClient,
	}
}

func (r *RedisRepository) GetRedisClient() *redis.Client {
	return r.redisClient
}

func (r *RedisRepository) DisconnectRedisDB(ctx context.Context) {
	debugger := logs.GetLogger()
	if r.redisClient == nil {
		debugger.Fatalw("redis client is nil")
		return
	}
	err := r.redisClient.Close()
	if err != nil {
		debugger.Infow("disconnected from redis failed")
	} else {
		debugger.Infow("disconnected from redis")
	}
}

// join array of strings with a separator : and return a string
func CachePath(s []string) string {
	return strings.Join(s, ":")
}
