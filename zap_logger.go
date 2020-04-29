package zapLogger

// Logger 定制化
type Logger interface {
	With(args ...interface{}) *logger
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
}
