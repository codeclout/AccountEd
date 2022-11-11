package ports

type HttpFrameworkInPort interface {
	Run(logger func(m ...interface{}))
}
