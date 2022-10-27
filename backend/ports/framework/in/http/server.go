package ports

type HTTPPort interface {
	Run(logger func(m ...interface{}))
}
