package aws

type Credentials interface {
	LoadCreds() ([]byte, error)
}
