package core

import (
	"context"
	"example-service/internal/domain"
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	appConfigs *domain.AppConfig
	onceEnvs   sync.Once
)

func GetEnviroments() *domain.AppConfig {
	onceEnvs.Do(func() {
		appConfigs = initConfigs()
	})

	return appConfigs
}

func initConfigs() *domain.AppConfig {
	_ = godotenv.Load()
	var appConfig domain.AppConfig
	err := envconfig.Process(context.Background(), &appConfig)
	if err != nil {
		panic(fmt.Sprintf("error on load configs: %s", err))
	}
	return &appConfig
}
