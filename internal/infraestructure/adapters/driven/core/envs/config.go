package envs

import (
	"context"
	"example-service/internal/domain/core"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func WithEnvs(ctx context.Context, cfg *core.AppConfig) context.Context {
	_ = godotenv.Load()
	envconfig.Process(ctx, cfg)
	ctx = context.WithValue(ctx, "envs", cfg)
	return ctx
}
