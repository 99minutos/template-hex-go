package ports

type LoggerPort interface {
	Debugw(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(template string, args ...interface{})
	Errorw(template string, args ...interface{})
	Fatalw(template string, args ...interface{})
}
