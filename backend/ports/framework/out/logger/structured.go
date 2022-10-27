package ports

type LoggerPort interface {
	HttpMiddlewareLogger(msg ...interface{})
	Initialize()
	Log(level string, msg string)
	Sync()
}
