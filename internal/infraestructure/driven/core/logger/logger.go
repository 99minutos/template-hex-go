package logger

import (
	ports "example-service/internal/domain/ports/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CoreLogger struct {
	logger *zap.SugaredLogger
}

var _ ports.LoggerPort = (*CoreLogger)(nil)

func initLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.TimeKey = "timestamp"
	log, _ := config.Build()
	return log.Sugar()
}

func NewLogger() ports.LoggerPort {
	return &CoreLogger{
		logger: initLogger(),
	}
}

func (s *CoreLogger) Debugw(template string, args ...interface{}) {
	s.logger.Debugw(template, args...)
}
func (s *CoreLogger) Infow(msg string, keysAndValues ...interface{}) {
	s.logger.Infow(msg, keysAndValues...)
}

func (s *CoreLogger) Warnw(template string, args ...interface{}) {
	s.logger.Warnw(template, args...)
}

func (s *CoreLogger) Errorw(template string, args ...interface{}) {
	s.logger.Errorw(template, args...)
}
func (s *CoreLogger) Fatalw(template string, args ...interface{}) {
	s.logger.Fatalw(template, args...)
}
