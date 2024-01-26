package core

import (
	ports "example-service/internal/domain/ports/core"
)

type AppContext struct {
	Envs   *AppConfig
	Logger ports.LoggerPort
}

func NewAppContext(
	envs *AppConfig,
	logger ports.LoggerPort,
) *AppContext {
	return &AppContext{
		Envs:   envs,
		Logger: logger,
	}
}

func (c *AppContext) Debugw(template string, args ...interface{}) {
	if c.Envs.DebugMode {
		c.Logger.Debugw(template, args...)
	}
}
func (c *AppContext) Infow(msg string, keysAndValues ...interface{}) {
	c.Logger.Infow(msg, keysAndValues...)
}

func (c *AppContext) Warnw(template string, args ...interface{}) {
	c.Logger.Warnw(template, args...)
}

func (c *AppContext) Errorw(template string, args ...interface{}) {
	c.Logger.Errorw(template, args...)
}
func (c *AppContext) Fatalw(template string, args ...interface{}) {
	c.Logger.Fatalw(template, args...)
}
