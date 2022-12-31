package aws

type ParameterPort interface {
	GetParam(name *string) (*[]byte, error)
	GetSecret(id *string) (*string, error)
	GetRoleConnectionString(*string) (*string, error)
}
