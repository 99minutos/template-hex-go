package core

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     zap.SugaredLogger
	onceLogger sync.Once
)

func CreateLogger(loggerName string) zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.TimeKey = "timestamp"
	log, _ := config.Build()
	return *log.Sugar().Named(loggerName)
}

func GetDefaultLogger() zap.SugaredLogger {
	onceLogger.Do(func() {
		cfg := GetEnviroments()
		logger = CreateLogger(cfg.AppName)
	})

	return logger
}
