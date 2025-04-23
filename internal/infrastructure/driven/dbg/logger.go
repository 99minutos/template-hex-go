package dbg

import (
	"context"
	"fmt"
	"sync"

	"service/internal/infrastructure/driven/core"
	"service/internal/infrastructure/driven/tracer"

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

func GetLogger() zap.SugaredLogger {
	onceLogger.Do(func() {
		cfg := core.GetEnvironments()
		logger = CreateLogger(cfg.AppName)
	})

	return logger
}

func GetLoggerWithContext(ctx context.Context) zap.SugaredLogger {
	cfg := core.GetEnvironments()
	logger = CreateLogger(cfg.AppName)

	_, span := tracer.GetTracer().Start(ctx, "GetLogger")
	traceId := span.SpanContext().TraceID().String()
	logger.With("trace_id", traceId)
	logger.With("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", cfg.ProjectId, traceId))

	return logger
}
