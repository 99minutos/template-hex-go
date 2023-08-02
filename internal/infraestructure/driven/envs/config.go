package envs

import (
	"context"
	cfgs "example-service/internal/config"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func WithEnvs(ctx context.Context, cfg *cfgs.AppConfig) context.Context {
	_ = godotenv.Load()
	envconfig.Process(ctx, cfg)
	ctx = context.WithValue(ctx, "envs", cfg)
	return ctx
}
