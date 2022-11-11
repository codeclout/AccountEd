package ports

type LogFrameworkOutPort interface {
	HttpMiddlewareLogger(msg ...interface{})
	Initialize()
	Log(level string, msg string)
	Sync()
}
