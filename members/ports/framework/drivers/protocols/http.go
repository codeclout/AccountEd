package protocols

type httpProtocol interface {
  Run(func(message ...interface{}))
}
