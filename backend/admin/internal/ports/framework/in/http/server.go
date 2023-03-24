package http

type HttpFrameworkInPort interface {
  Run(logger func(m ...interface{}))
}
