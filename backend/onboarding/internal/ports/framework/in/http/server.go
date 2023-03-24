package http

type ServerFrameworkPort interface {
	Run(logger func(m ...interface{}))
}
