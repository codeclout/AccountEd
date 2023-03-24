package hcl

type RuntimeConfigPort interface {
	GetConfig(p []byte) []byte
}
